package models

import (
	"myturbogarage/helpers"
	"time"
	"gorm.io/gorm"
)

// User ...
type User struct {
	ID               string    `json:"id"`
	UserName         string    `json:"userName"`
	Password         string    `json:"password"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	IsMobileVerified bool      `json:"isMobileVerified"`
	IsEmailVerified  bool      `json:"IsEmailVerified"`
	OTP              string    `json:"otp"`
	OTPGenerated     time.Time `json:"otpExpiry"`
	IsActive         bool      `json:"isActive"`
	CreatedAt        time.Time `json:"createdAt"`
	LastLogin        time.Time `json:"lastLogin"`
	CreatedByID      *string   `json:"createdByID"`
	IsAdmin          bool      `json:"isAdmin"`
	ImageUrl         string    `json:"imageUrl"`
}

// SetPassword
func (u *User) SetPassword(password string) {
	hash := helpers.HashPassword(password)
	u.Password = string(hash)
}

// IsPasswordValid
func (u *User) IsPasswordValid(password string) bool {
	return helpers.ComparePassword([]byte(u.Password), password)
}

// IsOTPValid ...
func (u *User) IsOTPValid(otp string) bool {
	return helpers.ComparePassword([]byte(u.OTP), otp) && time.Since(u.OTPGenerated).Minutes() < 5
}

func (u *User) SetOTP(otp string) {
	hash := helpers.HashPassword(otp)
	u.OTP = string(hash)
	u.OTPGenerated = time.Now()
}

// Role ...
type Role struct {
	ID          string `json:"id"`
	GarageID    string `json:"garageID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions string `json:"permissions"`
}

// Staff models contains users with access rights
type Staff struct {
	ID       string `json:"id"`
	UserID   string `json:"userID"`
	GarageID string `json:"garageID"`
	RoleID   string `json:"roleID"`
	Role     *Role  `json:"role"`
}

func (s *Staff) AfterFind(tx *gorm.DB) (err error) {
	if err := tx.Where("id = ?", s.RoleID).First(&s.Role).Error; err != nil {
		return err
	}
	return nil
}
