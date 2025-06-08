package inputs

// RegisterRequest ...
type RegisterRequest struct {
	UserName  string `json:"userName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`

}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ForgotPasswordRequest ...
type ForgotPasswordRequest struct {
	Email  string `json:"email" binding:"required"`
	Domain string `json:"domain" binding:"required"`
}

// ResetPasswordRequest ...
type ResetPasswordRequest struct {
	Token           string `json:"token" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Domain          string `json:"domain" binding:"required"`
}

// VerifyEmailRequest ...
type VerifyEmailRequest struct {
	Token  string `json:"token" binding:"required"`
	Domain string `json:"domain" binding:"required"`
}

// ResendVerificationEmailRequest ...
type ResendVerificationEmailRequest struct {
	Email  string `json:"email" binding:"required"`
	Domain string `json:"domain" binding:"required"`
}

// ResendOTPRequest ...
type ResendOTPRequest struct {
	Email string `json:"email" binding:"required"`
}

// VerifyOTPRequest ...
type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}