package inputs

import (
	"mime/multipart"
)

// CustomField ...
type CustomField map[string]any

// CreateProductRequest ...
type CreateProductRequest struct {
	Type         string                  `form:"type"`
	CategoryID   string                  `form:"categoryID"`
	GarageID     string                  `form:"garageID"`
	Form         CustomField             `form:"form"`
	Images       []*multipart.FileHeader `form:"images"`
	Name         string                  `form:"name"`
	Description  string                  `form:"description"`
	Price        float64                 `form:"price"`
	PackagingFee float64                 `form:"packagingFee"`
	Custom       CustomField             `form:"custom"`
	Filter       CustomField             `form:"filter"`
	Extras       CustomField             `form:"extras"`
	CreatedByID  string                  `form:"createdById"`
}

// CreateProductGroupInput ...
type CreateProductGroupInput struct {
	GarageID string `json:"garageID"`
	Name     string `json:"name"`
}

// CreateProductTagRequest ...
type CreateProductTagRequest struct {
	GarageID    string `json:"garageID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
