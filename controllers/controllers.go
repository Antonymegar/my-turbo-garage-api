package controllers

import (
	"fmt"

	"myturbogarage/helpers"
	logger "myturbogarage/loggers"

	"github.com/gin-gonic/gin"
)

var log = logger.NewLogger()

func getCtxUser(ctx *gin.Context) (*helpers.AuthClaims, error) {
	ctxUser, ok := ctx.Get("claims")
	if !ok {
		return nil, fmt.Errorf("user claims not found in context")
	}

	claims, ok := ctxUser.(*helpers.AuthClaims)
	if !ok {
		return nil, fmt.Errorf("user in context not parsed")
	}

	return claims, nil
}
