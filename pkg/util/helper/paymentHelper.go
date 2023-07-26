package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/razorpay/razorpay-go"
)

func MakeRazorPayPaymentId(amount int) (string, error) {
	razorPayKey := config.GetConfig().RazorPayKeyId
	razorPaySecret := config.GetConfig().RazorPayKeySecret

	client := razorpay.NewClient(razorPayKey, razorPaySecret)

	data := map[string]interface{}{
		"amount":   amount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		return "", fmt.Errorf("Problem in getting repository information %s", err)
	}
	//save order id from the body
	RazorpayOrderId := body["id"].(string)

	return RazorpayOrderId, nil
}

func VerifyRazorPayPayment(signature, orderId, paymentId string) error {
	// Get razor pay api config
	razorPayKey := config.GetConfig().RazorPayKeyId
	razorPaySecret := config.GetConfig().RazorPayKeySecret

	// Verify signature
	data := orderId + "|" + paymentId
	h := hmac.New(sha256.New, []byte(razorPaySecret))
	_, err := h.Write([]byte(data))
	if err != nil {
		return err
	}
	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return err
	}
	// verify payment
	razorpayClient := razorpay.NewClient(razorPayKey, razorPaySecret)

	// fetch payment and verify
	payment, err := razorpayClient.Payment.Fetch(paymentId, nil, nil)
	if err != nil {
		return err
	}
	// check payment status
	if payment["status"] != "captured" {
		return fmt.Errorf("failed to verify payment \n razor pay payment with payment_id %v", paymentId)
	}
	return nil
}

func ValidateWebhookSignature(webhookBody, webhookSignature string) bool {
	mac := hmac.New(sha256.New, []byte(config.GetConfig().RazorPayKeySecret))
	mac.Write([]byte(webhookBody))
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(expectedSignature), []byte(webhookSignature))
}
