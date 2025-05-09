package routes

import (
	"myturbogarage/controllers"

	"github.com/gin-gonic/gin"
)

func initAuth(router *gin.Engine) {
	r := router.Group("/auth")
	{
		r.POST("/register", check(controllers.Register))
	}
}
