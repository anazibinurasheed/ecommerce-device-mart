package request

type Category struct {
	CategoryName string `json:"categoryname" binding:"required,min=2"`
}

type Product struct {
	CategoryID          int    `json:"category_id"`
	ProductName         string `json:"product_name" binding:"required"`
	Brand               string `json:"-"`
	Product_Description string `json:"product_description" binding:"required"`
	SKU                 string `json:"-"`
	Price               int    `json:"price" binding:"required"`
	ProductImage        string `json:"product_image" `
	IsBlocked           bool   `json:"-"`
}

type Rating struct {
	UserID      int    `json:"-"`
	ProductID   int    `json:"-"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}
