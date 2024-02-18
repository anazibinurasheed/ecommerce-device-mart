package response

type PaymentDetails struct {
	Username        string `json:"username"`
	RazorPayOrderID string `json:"razorpay_order_id"`
	Amount          int    `json:"amount"`
}