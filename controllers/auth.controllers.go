package controllers

import (
	"fmt"
	"myturbogarage/config"
	"myturbogarage/errors"
	"myturbogarage/helpers"
	"myturbogarage/inputs"
	"myturbogarage/models"
	"myturbogarage/outputs"
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
		ID:       uuid.NewString(),
		UserName: req.UserName,
		Email:    req.Email,
	}
	user.SetPassword(req.Password)
	if err := config.Create(user); err != nil {
		return errors.Conflict(err.Error())
	}
	token, err := helpers.GenerateAuthTokens(&helpers.AuthClaims{
		ID:    user.ID,
		Email: user.Email,
	}, time.Hour*24)
	fmt.Sprint(token)
	if err != nil {
		return errors.Forbidden(err.Error())
	}
	ctx.JSON(200, gin.H{"message": "user created successfully"})
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
