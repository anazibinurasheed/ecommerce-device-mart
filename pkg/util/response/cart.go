package response

import "github.com/anazibinurasheed/project-device-mart/pkg/domain"

type Cart struct {
	ID          uint         `json:"cart_id"`
	ProductID   uint         `json:"product_id"`
	ProductName string       `json:"product_name"`
	Images      domain.JSONB `json:"images"`
	Price       int          `json:"price"`
	Brand       string       `json:"brand"`
	Qty         int          `json:"qty"`
}

type CartItems struct {
	Cart     []Cart  `json:"items"`
	Discount float32 `json:"discount"`
	Total    float32 `json:"total"`
}
