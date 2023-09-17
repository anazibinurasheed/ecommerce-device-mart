package request

type VerifyPayment struct {
	Signature         string `json:"razorpay_signature" binding:"required"`
	RazorpayOrderID   string `json:"razorpay_order_id" binding:"required"`
	RazorPayPaymentID string `json:"razorpay_payment_id" binding:"required"`
}
