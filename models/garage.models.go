package models

import (
	"encoding/json"
	"time"
)

// type CustomField struct
type CustomField map[string]any

// Vendor ...
type Garage struct {
	ID                              string          `json:"id"`
	OwnerID                         string          `json:"ownerID"`
	Name                            string          `json:"name"`
	Logo                            string          `json:"logo"`
	Banner                          string          `json:"banner"`
	Email                           string          `json:"email"`
	PhoneNumber                     string          `json:"phoneNumber"`
	TradeName                       string          `json:"tradeName"`
	BusinessRegistrationCertificate string          `json:"businessRegistrationCertificate"`
	Description                     string          `json:"description"`
	KRAPin                          string          `json:"KRAPin"`
	KRACertificate                  string          `json:"KRACertificate"`
	CR12Certificate                 string          `json:"CR12Certificate"`
	PostalAddress                   string          `json:"postalAddress"`
	PostalCode                      string          `json:"postalCode"`
	County                          string          `json:"county"`
	Town                            string          `json:"town"`
	StreetAddress                   string          `json:"streetAddress"`
	Building                        string          `json:"building"`
	IsVerified                      bool            `json:"isVerified"`
	Meta                            json.RawMessage `json:"meta" gorm:"type:jsonb"`
	CreatedAt                       time.Time       `json:"createdAt"`
	UpdatedAt                       time.Time       `json:"updatedAt"`
}
