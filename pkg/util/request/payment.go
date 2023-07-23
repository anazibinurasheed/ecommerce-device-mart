package request

type VerifyPayment struct {
	Signature         string `json:"razorpay_signature"`
	RazorpayOrderId   string `json:"razorpay_order_id"`
	RazorPayPaymentId string `json:"razorpay_payment_id"`
}
