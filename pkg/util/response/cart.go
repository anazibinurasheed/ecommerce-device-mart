package response

type Cart struct {
	ID          uint   `json:"cart_id"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Brand       string `json:"brand"`
	Qty         int    `json:"qty"`
}

type CartItems struct {
	Cart     []Cart  `json:"items"`
	Discount float32 `json:"discount"`
	Total    float32 `json:"total"`
}
