package outputs

import (
	"myturbogarage/helpers"
	"myturbogarage/models"
	"time"
)

// LoginResponse ...
type LoginResponse struct {
	ID               string    `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	IsMobileVerified bool      `json:"isMobileVerified"`
	IsEmailVerified  bool      `json:"IsEmailVerified"`
	IsActive         bool      `json:"isActive"`
	IsAdmin          bool      `json:"isAdmin"`
	ImageUrl         string    `json:"imageUrl"`
	LastLogin        time.Time `json:"lastLogin"`
	AccessToken      string    `json:"accessToken"`
	RefreshToken     string    `json:"refreshToken"`
}

// NewLoginResponse ...
func NewLoginResponse(user *models.User) *LoginResponse {
	token, err := helpers.GenerateAuthTokens(&helpers.AuthClaims{ID: user.ID, Email: user.Email})
	if err != nil {
		return nil
	}

	return &LoginResponse{
		ID:               user.ID,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
		Phone:            user.Phone,
		IsMobileVerified: user.IsMobileVerified,
		IsEmailVerified:  user.IsEmailVerified,
		IsActive:         user.IsActive,
		IsAdmin:          user.IsAdmin,
		ImageUrl:         user.ImageUrl,
		LastLogin:        user.LastLogin,
		AccessToken:      token.AccessToken,
		RefreshToken:     token.RefreshToken,
	}
}
