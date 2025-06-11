package inputs

import (
	"mime/multipart"
	"time"
)

// UserBuyRequest ...
type UserBuyRequest struct {
	ProductID string `json:"productID" binding:"required"`
	UserID    string `json:"userID"`
	GarageID  string `json:"garageID"`
	Quantity  int    `json:"quantity"`
}

// UserRegisterRequest ...
type UserRegisterRequest struct {
	UserID   string         `json:"userID" binding:"required"`
	GarageID string         `json:"garageID" binding:"required"`
	Data     map[string]any `json:"data" binding:"required"`
}

// UserTaskCheckout ...
type UserTaskCheckout struct {
	UserTaskID            string     `json:"userTaskID" binding:"required"`
	CheckoutType          string     `json:"checkoutType" binding:"required"`
	CheckoutExpectedTime  *time.Time `json:"checkoutExpectedTime"`
	CheckoutDeliveryType  string     `json:"checkoutDeliveryType" binding:"required"`
	CheckoutLocation      string     `json:"checkoutLocation" binding:"required"`
	CheckoutPaymentMethod string     `json:"checkoutPaymentMethod" binding:"required"`
	CheckoutAccountNumber string     `json:"checkoutAccountNumber" binding:"required"`
	CheckoutComments      string     `json:"checkoutComments"`
}

// CreateGarageRequest ...
type CreateGarageRequest struct {
	Name                            string                `form:"name"`
	Logo                            *multipart.FileHeader `form:"logo"`
	Email                           string                `form:"email"`
	PhoneNumber                     string                `form:"phoneNumber"`
	TradeName                       string                `form:"tradeName"`
	BusinessRegistrationCertificate *multipart.FileHeader `form:"businessRegistrationCertificate"`
	Description                     string                `form:"description"`
	KRAPin                          string                `form:"KRAPin"`
	KRACertificate                  *multipart.FileHeader `form:"KRACertificate"`
	CR12Certificate                 *multipart.FileHeader `form:"CR12Certificate"`
	PostalAddress                   string                `form:"postalAddress"`
	PostalCode                      string                `form:"postalCode"`
	County                          string                `form:"county"`
	Town                            string                `form:"town"`
	StreetAddress                   string                `form:"streetAddress"`
	Building                        string                `form:"building"`
	Meta                            CustomField           `form:"meta"`
}
