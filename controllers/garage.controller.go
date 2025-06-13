package controllers

import (
	"encoding/json"
	"fmt"
	"myturbogarage/config"
	"myturbogarage/inputs"
	"myturbogarage/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateGarage(ctx *gin.Context) error {
	var req inputs.CreateGarageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err
	}
	ctxUser, err := getCtxUser(ctx)
	if err != nil {
		return err
	}
	meta, err := json.Marshal(req.Meta)
	if err != nil {
		return fmt.Errorf("failed to marshal meta: %s", err.Error())
	}

	garage := models.Garage{
		ID:            uuid.NewString(),
		OwnerID:       ctxUser.ID,
		Name:          req.Name,
		Email:         req.Email,
		PhoneNumber:   req.PhoneNumber,
		TradeName:     req.TradeName,
		Description:   req.Description,
		KRAPin:        req.KRAPin,
		PostalAddress: req.PostalAddress,
		PostalCode:    req.PostalCode,
		County:        req.County,
		Town:          req.Town,
		StreetAddress: req.StreetAddress,
		Building:      req.Building,
		Meta:          meta,
	}
	if err := config.Create(&garage); err != nil {
		return err
	}
	ctx.JSON(200, garage)
	return nil

}

func UpdateGarage(ctx *gin.Context) error {
	var req models.Garage
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err
	}
	garage, err := config.FindOne(&models.Garage{ID: ctx.Param("id")})
	if err != nil {
		return err
	}

	meta, err := json.Marshal(req.Meta)
	if err != nil {
		return fmt.Errorf("failed to marshal meta:%s", err.Error())
	}
	garage.Name = req.Name
	garage.Logo = req.Logo
	garage.Banner = req.Banner
	garage.Email = req.Email
	garage.PhoneNumber = req.PhoneNumber
	garage.TradeName = req.TradeName
	garage.BusinessRegistrationCertificate = req.BusinessRegistrationCertificate
	garage.Description = req.Description
	garage.KRAPin = req.KRAPin
	garage.KRACertificate = req.KRACertificate
	garage.CR12Certificate = req.CR12Certificate
	garage.PostalAddress = req.PostalAddress
	garage.PostalCode = req.PostalCode
	garage.County = req.County
	garage.Town = req.Town
	garage.StreetAddress = req.StreetAddress
	garage.Building = req.Building
	garage.IsVerified = req.IsVerified
	garage.Meta = meta

	if err := config.Update(&garage); err != nil {
		return err
	}
	return nil
}
func DeleteGarage(ctx *gin.Context) error {
	garage, err := config.FindOne(&models.Garage{ID: ctx.Param("id")})
	if err != nil {
		return err
	}
	if err := config.Delete(&garage); err != nil {
		return err
	}
	ctx.JSON(200, gin.H{"message": "Garage Deleted successfully"})
	return nil
}

func GetGarages(ctx *gin.Context) error {
	garages, err := config.FindAll(&models.Garage{})
	if err != nil {
		return err
	}
	ctx.JSON(200, garages)
	return nil
}

func GetMyGarages(ctx *gin.Context) error {
	ctxUser, err := getCtxUser(ctx)
	if err != nil {
		return err
	}
	var garages []*models.Garage
	if err:= config.Conn().Where("owner_id = ?",ctxUser.ID).Find(&garages).Error; err!=nil{
		return err
	}
	ctx.JSON(200, garages)
	return nil
}
func GetGarage(ctx *gin.Context) error {
	garage, err :=config.FindOne(&models.Garage{ID:ctx.Param("id")})
	if err!= nil{
		return err
	}
	ctx.JSON(200, garage)
	return nil
}
