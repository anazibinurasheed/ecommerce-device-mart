package request

type VerifyPayment struct {
	Signature         string `json:"razorpay_signature" binding:"required"`
	RazorpayOrderId   string `json:"razorpay_order_id" binding:"required"`
	RazorPayPaymentId string `json:"razorpay_payment_id" binding:"required"`
}
