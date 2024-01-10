package request

import "github.com/anazibinurasheed/project-device-mart/pkg/domain"

type Category struct {
	CategoryName string `json:"category_name" binding:"required,min=2"`
}

type Product struct {
	CategoryID         int          `json:"-"`
	ProductName        string       `json:"product_name" binding:"required"`
	ProductDescription string       `json:"product_description" binding:"required"`
	Price              int          `json:"price" binding:"required"`
	Images             domain.JSONB `json:"-" `
	SKU                string       `json:"-"`
	Brand              string       `json:"-"`
	IsBlocked          bool         `json:"-"`
}

type UpdateProduct struct {
	CategoryID         int    `json:"category_id" binding:"required"`
	ProductName        string `json:"product_name" binding:"required"`
	ProductDescription string `json:"product_description" binding:"required"`
	Price              int    `json:"price" binding:"required"`
}

type Rating struct {
	UserID      int    `json:"-"`
	ProductID   int    `json:"-"`
	Rating      int    `json:"rating" binding:"required"`
	Description string `json:"description" binding:"required"`
}
