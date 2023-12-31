package response

import "github.com/anazibinurasheed/project-device-mart/pkg/domain"

type Category struct {
	ID            int          `json:"id"`
	Category_Name string       `json:"category_name"`
	Images        domain.JSONB `json:"images"`
	IsBlocked     bool         `json:"is_blocked"`
}

// type Product struct {
// 	ID                  uint   `json:"id"`
// 	CategoryID          int    `json:"category_id"`
// 	ProductName         string `json:"product_name"`
// 	Price               int    `json:"price"`
// 	SKU                 string `json:"sku,omitempty"`
// 	Brand               string `json:"brand"`
// 	Product_Description string `json:"product_description,omitempty"`
// 	Product_Image       string `json:"product_image"`
// 	IsBlocked           bool   `json:"is_blocked"`
// }

type Product struct {
	ID                  uint         `json:"id"`
	CategoryID          int          `json:"category_id"`
	ProductName         string       `json:"product_name"`
	Price               int          `json:"price"`
	SKU                 string       `json:"sku,omitempty"`
	Brand               string       `json:"brand"`
	Product_Description string       `json:"product_description,omitempty"`
	Images              domain.JSONB `json:"images,omitempty"`
	IsWishlisted        bool         `json:"is_wishlisted,omitempty"`
	IsBlocked           bool         `json:"is_blocked,omitempty"`
}

//	type ProductItem struct {
//		ID                  uint     `json:"id"`
//		CategoryID          int      `json:"category_id"`
//		Product_Name        string   `json:"product_name"`
//		Price               int      `json:"price"`
//		SKU                 string   `json:"sku"`
//		Brand               string   `json:"brand"`
//		Product_Description string   `json:"product_description"`
//		Product_Image       string   `json:"product_image"`
//		Is_Blocked          bool     `json:"is_blocked"`
//		RatingAndReviews    []Rating `json:"rating_and_reviews"`
//	}
type ProductItem struct {
	ID                  uint         `json:"id"`
	CategoryID          int          `json:"category_id"`
	Product_Name        string       `json:"product_name"`
	Price               int          `json:"price"`
	SKU                 string       `json:"sku"`
	Brand               string       `json:"brand"`
	Product_Description string       `json:"product_description"`
	Images              domain.JSONB `json:"images"`
	IsWishlisted        bool         `json:"is_wishlisted"`
	Is_Blocked          bool         `json:"is_blocked"`
	RatingAndReviews    []Rating     `json:"rating_and_reviews"`
}

type Rating struct {
	ID          int    `json:"rating_id"`
	User_name   string `json:"user_name"`
	Rating      int    `json:"rating"`
	Description string `json:"desription"`
}
