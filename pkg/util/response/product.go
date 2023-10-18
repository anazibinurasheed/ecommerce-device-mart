package response

type Category struct {
	ID            int    `json:"id"`
	Category_Name string `json:"category_name"`
}
type Product struct {
	ID                  uint   `json:"id"`
	CategoryID          int    `json:"category_id"`
	ProductName         string `json:"product_name"`
	Price               int    `json:"price"`
	SKU                 string `json:"sku,omitempty"`
	Brand               string `json:"brand"`
	Product_Description string `json:"product_description,omitempty"`
	Product_Image       string `json:"product_image"`
	IsBlocked           bool   `json:"is_blocked"`
}

// ProductSlice is a type to specify detailed API specification on swagger.
// Not used anywhere in the program except the swagger documentation. 
type ProductSlice = []Product

type ProductItem struct {
	ID                  uint     `json:"id"`
	CategoryID          int      `json:"category_id"`
	Product_Name        string   `json:"product_name"`
	Price               int      `json:"price"`
	SKU                 string   `json:"sku"`
	Brand               string   `json:"brand"`
	Product_Description string   `json:"product_description"`
	Product_Image       string   `json:"product_image"`
	Is_Blocked          bool     `json:"is_blocked"`
	RatingAndReviews    []Rating `json:"rating_and_reviews"`
}

type Rating struct {
	ID          int    `json:"rating_id"`
	User_name   string `json:"user_name"`
	Rating      int    `json:"rating"`
	Description string `json:"desription"`
}
