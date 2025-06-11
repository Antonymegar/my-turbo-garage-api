package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ProductType ...
type ProductType string

// ProductTypes ...
var (
	BuyProduct     ProductType = "buy-n-order"
	ServiceProduct ProductType = "service"
)

// Field ...
type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Options  string `json:"options"`
	Required bool   `json:"required"`
}

// Form ...
type Form struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

// ProductTag ...
type ProductTag struct {
	ID          string `json:"id"`
	VendorID    string `json:"vendorID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Product ...
type Product struct {
	ID             string          `json:"id"`
	TaskID         string          `json:"taskID"`
	ProductGroupID string          `json:"productGroupID"`
	SubtaskID      string          `json:"subtaskID"`
	CategoryID     string          `json:"categoryID"`
	VendorID       string          `json:"vendorID"`
	Form           json.RawMessage `json:"form" gorm:"type:jsonb"`
	Type           ProductType     `json:"type"`
	Images         json.RawMessage `json:"images" gorm:"type:jsonb"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Price          float64         `json:"price"`
	PackagingFee   float64         `json:"packagingFee"`
	TotalRatings   int64           `json:"totalRatings"`
	TotalReviews   int64           `json:"totalReviews"`
	Custom         json.RawMessage `json:"custom" gorm:"type:jsonb"`
	Filters        json.RawMessage `json:"filters" gorm:"type:jsonb"`
	Extras         json.RawMessage `json:"extras" gorm:"type:jsonb"`
	Active         bool            `json:"active"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt  `json:"deletedAt"`
	CreatedByID    string          `json:"createdByID"`
	UpdatedByID    string          `json:"updatedByID"`
	Garage         *Garage         `json:"garage,omitempty" gorm:"-"`   // virtual field
	Category       *Category       `json:"category,omitempty" gorm:"-"` // virtual field
}

// AfterFind ...
func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	if p.Images == nil {
		p.Images = json.RawMessage(`[]`)
	}
	if p.Custom == nil {
		p.Custom = json.RawMessage(`{}`)
	}
	if p.Filters == nil {
		p.Filters = json.RawMessage(`{}`)
	}
	if p.Extras == nil {
		p.Extras = json.RawMessage(`{}`)
	}

	if err := tx.Where("id = ?", p.VendorID).First(&p.Garage).Error; err != nil {
		fmt.Println("product vendor not found. error: ", err.Error())
	}

	if err := tx.Where("id = ?", p.CategoryID).First(&p.Category).Error; err != nil {
		fmt.Println("product category not found. error: ", err.Error())
	}
	return
}

// GetImage ...
func (p *Product) GetImage() string {
	var images []string
	if err := json.Unmarshal(p.Images, &images); err != nil {
		return ""
	}

	if len(images) == 0 {
		return ""
	}

	return images[0]
}

// ProductRating ...

// ProductGroup ...
type ProductGroup struct {
	ID        string    `json:"id"`
	VendorID  string    `json:"vendorID"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
