package response

import "time"

// No external connections
type OrderLine struct {
	ID              uint      `json:"id"`
	UserID          uint      `json:"user_id"`
	AddressesID     uint      `json:"addresses_id"`
	ProductID       uint      `json:"product_id"`
	PaymentMethodID int       `json:"payment_method_id"`
	OrderStatusID   int       `json:"order_status_id"`
	Qty             int       `json:"qty"`
	Price           float32   `json:"price"`
	CouponID        uint      `json:"coupon_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Orders struct {
	OrderID         int     `json:"order_id"`
	ProductID       int     `json:"product_id"`
	ProductImage    string  `json:"product_image"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float32 `json:"product_price"`
	OrderStatusID   int     `json:"-"`
	OrderStatus     string  `json:"order_status"`
	PaymentMethodID int     `json:"-"`
	PaymentMethod   string  `json:"payment_method"`
	AddressesID     int     `json:"-"`
	DeliveryAddress string  `json:"delivery_address"`
}

type Invoice struct {
	Date            time.Time `json:"date"`
	OrderDate       string    `json:"order_date"`
	OrderID         int       `json:"order_id"`
	DeliveryAddress string    `json:"delivery_address"`
	ProductName     string    `json:"product_name"`
	PaymentMethod   string    `json:"payment_method"`
	ProductPrice    int       `json:"product_price"`
	Discount        float32   `json:"discount"`
	TotalAmount     float32   `json:"total_amount"`
}

type MonthlySalesReport struct {
	Date                  string  `json:"date"`
	ReportFromDate        string  `json:"report_from"`
	TopSellingBrand       string  `json:"top_selling_brand"`
	TopSellingProduct     string  `json:"top_selling_product"`
	TopSoldQuantity       int     `json:"total_quantity_sold"`
	TotalSalesCount       int     `json:"total_sales_count"`
	AverageOrderValue     float32 `json:"average_order_value"`
	SalesGrowthPercentage float32 `json:"sales_growth_percentage"`
	TotalCouponIncentive  float32 `json:"total_coupon_incentive"`
	TotalRevenue          float32 `json:"total_revenue"`
}

type Checkout struct {
	Address        []Address       `json:"delivery_address"`
	Cart           []Cart          `json:"items"`
	Discount       float32         `json:"discount"`
	Total          float32         `json:"total"`
	PaymentOptions []PaymentMethod `json:"payment_options"`
}

type OrderManagement struct {
	OrderStatuses []OrderStatus `json:"order_statuses"`
	Orders        []Orders      `json:"orders"`
}

type PaymentMethod struct {
	ID         int    `json:"id"`
	MethodName string `json:"method"`
}

type OrderStatus struct {
	ID     uint   `json:"id"`
	Status string `json:"order_status"`
}

type TopSelling struct {
	ProductID int
	Quantity  int
}
