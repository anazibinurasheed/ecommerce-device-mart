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
	ID     uint   `json:""`
	Status string `json:""`
}

// type Orders struct {
// 	Order_Id          int     `json:"order_id"`
// 	Product_Id        int     `json:"product_id"`
// 	Product_Image     string  `json:"product_image"`
// 	Product_Name      string  `json:"product_name"`
// 	Product_Price     float32 `json:"product_price"`
// 	Order_Status_Id   int     `json:"order_status_id"`
// 	Order_Status      string  `json:"order_status"`
// 	Payment_Method_Id int     `json:"payment_method_id"`
// 	Payment_Method    string  `json:"payment_method"`
// 	Addresses_Id      int
// 	DeliveryAddress   string `json:"delivery_address"`
// }

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

// order_id | product_id | product_image |  product_name   | product_price | order_status_id | order_status | payment_method_id |  payment_method  | delivery_address

type Invoice struct {
}
