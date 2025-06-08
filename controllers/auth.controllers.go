package controllers

import (
	"fmt"
	"myturbogarage/config"
	"myturbogarage/errors"
	"myturbogarage/helpers"
	"myturbogarage/inputs"
	"myturbogarage/models"
	"myturbogarage/outputs"
	"myturbogarage/services/mail/sendgrid"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(ctx *gin.Context) error {
	var req inputs.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.BadRequest(err.Error())
	}
	count, err := config.Count(&models.User{Email: req.Email})
	if err != nil {
		return errors.NotFound(err.Error())
	}

	if count > 0 {
		return errors.Conflict("user with email already exists")
	}
	user := &models.User{
		ID:        uuid.New().String(),
		UserName:  req.UserName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
	user.SetPassword(req.Password)
	if err := config.Create(user); err != nil {
		return errors.Conflict(err.Error())
	}
	token, err := helpers.GenerateAuthTokens(&helpers.AuthClaims{
		ID:    user.ID,
		Email: user.Email,
	}, time.Hour*24)
	fmt.Println("token", token)
	if err != nil {
		return errors.Forbidden(err.Error())
	}
	ctx.JSON(200, gin.H{"message": "user created successfully "})
	return nil
}

func Login(ctx *gin.Context) error {
	var req inputs.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	user, err := config.FindOne(&models.User{Email: req.Email})
	if err != nil {
		return errors.NotFound(err.Error())
	}

	if !user.IsPasswordValid(req.Password) {
		return errors.Forbidden("invalid password provided")
	}

	if err := config.Update(&user); err != nil {
		return errors.Conflict(err.Error())
	}

	res := outputs.NewLoginResponse(user)
	ctx.JSON(200, res)
	return nil
}

func ForgotPassword(ctx *gin.Context) error {
	var req inputs.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	user, err := config.FindOne(&models.User{Email: req.Email})
	if err != nil {
		return errors.NotFound("user account not found")
	}

	token, err := helpers.GenerateAuthTokens(&helpers.AuthClaims{
		ID:    user.ID,
		Email: user.Email,
	}, time.Hour*24)

	if err != nil {
		return errors.Unauthorized("invalid credentials")
	}

	// send  reset password email
	go func() {
		data := map[string]interface{}{
			"firstName": user.FirstName,
			"appName":   "EpicApp",
			"link":      fmt.Sprintf("%s/reset-password?token=%s", req.Domain, token.AccessToken),
		}

		if err := sendgrid.SendEmail("templates/reset-password.html", "Reset Password", user.Email, data); err != nil {
			log.Error("failed to send verification email: %s", err.Error())
		}
	}()

	ctx.JSON(200, gin.H{"message": "reset password email sent successfully"})
	return nil
}

// ResetPassword ...
func ResetPassword(ctx *gin.Context) error {
	var req inputs.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	if req.NewPassword != req.ConfirmPassword {
		return errors.Unauthorized("passwords do not match")
	}

	claims, err := helpers.ValidateToken(req.Token)
	if err != nil {
		return errors.Unauthorized("invalid credentials")
	}

	user, err := config.FindOne(&models.User{ID: claims.ID})
	if err != nil {
		return errors.NotFound("user account not found")
	}

	user.SetPassword(req.NewPassword)
	if err := config.Update(&user); err != nil {
		return errors.Conflict("updating user profile failed")
	}

	res := outputs.NewLoginResponse(user)

	// send password reset confirmation email
	go func() {
		data := map[string]interface{}{}

		if err := sendgrid.SendEmail("templates/reset-password-done.html", "Password Updated", user.Email, data); err != nil {
			log.Error("failed to send password reset confirmation email: %s", err.Error())
		}
	}()

	ctx.JSON(200, res)
	return nil
}

// VerifyEmail ...
func VerifyEmail(ctx *gin.Context) error {
	var req inputs.VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	claims, err := helpers.ValidateToken(req.Token)
	if err != nil {
		return errors.Unauthorized("invalid credentials")
	}

	user, err := config.FindOne(&models.User{ID: claims.ID})
	if err != nil {
		return errors.NotFound("user account not found")
	}

	user.IsEmailVerified = true
	user.IsActive = true
	if err := config.Update(&user); err != nil {
		return errors.Conflict("updating user profile failed")
	}

	res := outputs.NewLoginResponse(user)

	// send email verification confirmation email done
	go func() {
		data := map[string]interface{}{
			"link": fmt.Sprintf("%s/login", req.Domain),
		}

		if err := sendgrid.SendEmail("templates/email-activation-result.html", "Account Verified", user.Email, data); err != nil {
			log.Error("failed to send email verification confirmation email: %s", err.Error())
		}
	}()

	ctx.JSON(200, res)
	return nil
}

// ResendVerificationEmail ...
func ResendVerificationEmail(ctx *gin.Context) error {
	var req inputs.ResendVerificationEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	user, err := config.FindOne(&models.User{Email: req.Email})
	if err != nil {
		return errors.NotFound("user account not found")
	}

	token, err := helpers.GenerateAuthTokens(&helpers.AuthClaims{
		ID:    user.ID,
		Email: user.Email,
	}, time.Hour*24)

	if err != nil {
		return errors.Unauthorized("invalid credentials")
	}

	fmt.Println("token: ", token.AccessToken)
	// send  reset password email
	go func() {
		data := map[string]interface{}{
			"firstName": user.FirstName,
			"appName":   "EpicApp",
			"link":      fmt.Sprintf("%s/verify-email?token=%s", req.Domain, token.AccessToken),
		}

		if err := sendgrid.SendEmail("templates/verify-email.html", "Verify your email", user.Email, data); err != nil {
			log.Error("failed to send verification email: %s", err.Error())
		}
	}()

	ctx.JSON(200, gin.H{"message": "verification email sent successfully"})
	return nil
}
