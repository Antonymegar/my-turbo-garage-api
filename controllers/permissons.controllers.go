package controllers

import (
	"myturbogarage/inputs"
	"myturbogarage/models"
	"net/http"

	"myturbogarage/config"

	"github.com/gin-gonic/gin"
)

func AddPermissionToRole(ctx *gin.Context) error {

	var req inputs.AddPermissionToRoleInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err
	}
	role, err := config.FindOne(&models.Role{ID: req.RoleID})
	if err != nil {
		return err
	}
	role.Permissions = req.Permissions
	if err := config.Update(&role); err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Permission added to role successfully"})
	return nil
}
