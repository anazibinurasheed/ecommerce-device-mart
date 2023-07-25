package response

import "time"

type CheckOut struct {
	Address        []Address       `json:"delivey_address"`
	Cart           []Cart          `json:"items"`
	Discount       float32         `json:"discount"`
	Total          float32         `json:"total"`
	PaymentOptions []PaymentMethod `json:"payment_options"`
}

type PaymentMethod struct {
	ID         int    `json:"id"`
	MethodName string `json:"method"`
}

type OrderLine struct {
	ID              uint      `json:""`
	UserID          uint      `json:""`
	AddressesID     uint      `json:""`
	ProductID       uint      `json:""`
	PaymentMethodID int       `json:""`
	OrderStatusId   int       `json:""`
	Qty             int       `json:""`
	Price           float32   `json:""`
	CouponID        uint      `json:""`
	CreatedAt       time.Time `json:""`
	UpdatedAt       time.Time `json:""`
}
type OrderStatus struct {
	ID     uint   `json:"id"`
	Status string `json:"order_status"`
}

type Orders struct {
	OrderID         int     `json:"order_id"`
	ProductID       int     `json:"product_id"`
	ProductImage    string  `json:"product_image"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float32 `json:"product_price"`
	OrderStatusID   int     `json:"order_status_id"`
	OrderStatus     string  `json:"order_status"`
	PaymentMethodID int     `json:"payment_method_id"`
	PaymentMethod   string  `json:"payment_method"`
	AddressesID     int     `json:"delivery_address_id"`
	DeliveryAddress string  `json:"delivery_address"`
}
type OrderManagement struct {
	OrderStatuses []OrderStatus `json:"order_statuses"`
	Orders        []Orders      `json:"orders"`
}
