package routes

import (
	"myturbogarage/controllers"

	"github.com/gin-gonic/gin"
)

func initAuth(router *gin.Engine) {
	r := router.Group("/auth")
	{
		r.POST("/register", check(controllers.Register))
		r.POST("/login", check(controllers.Login))
		r.POST("/forgot-password", check(controllers.ForgotPassword))
		r.POST("/reset-password", check(controllers.ResetPassword))

		r.POST("/verify-email", check(controllers.VerifyEmail))
		r.POST("/resend-verification-email", check(controllers.ResendVerificationEmail))

		r.POST("/verify-otp", check(controllers.VerifyOTP))
		r.POST("/resend-otp", check(controllers.RequestOTP))

		// who am i
		r.GET("/whoami", authCheck(controllers.WhoAmI))

		//upload avatar
		//r.POST("/upload-avatar", authCheck(controllers.UploadAvatar))
	}
}
