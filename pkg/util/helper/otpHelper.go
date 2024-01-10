package helper

import (
	"fmt"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TWILIO_ACCOUNT_SID string
var TWILIO_AUTH_TOKEN string
var VERIFY_SERVICE_SID string

var client *twilio.RestClient

func SendOtp(phone string) error {
	return nil // set predefined in development mode
	TWILIO_ACCOUNT_SID = config.GetConfig().TwilioAccountSid
	TWILIO_AUTH_TOKEN = config.GetConfig().TwilioAuthToken
	VERIFY_SERVICE_SID = config.GetConfig().TwilioServiceSid
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
	// fmt.Printf("1:%s,2:%s,3:%s:", TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN, VERIFY_SERVICE_SID)
	// fmt.Println("")
	// fmt.Printf("1:%s,2:%s,3:%s:", config.GetConfig().TwiliAccountSid, config.GetConfig().TwilioAuthToken, config.GetConfig().TwilioServiceSid)

	phone = "+91" + phone
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	_, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)

	if err != nil {
		return fmt.Errorf("sending otp failed %s", err)
	}

	return nil
}

func CheckOtp(phone string, code string) (string, error) {
	TestOtp := "0000" // Because of development mode
	if code == TestOtp {
		return "approved", nil
	} else {
		return "incorrect", fmt.Errorf("Failed to verify , incorrect otp provided")
	}
	TWILIO_ACCOUNT_SID = config.GetConfig().TwilioAccountSid
	TWILIO_AUTH_TOKEN = config.GetConfig().TwilioAuthToken
	VERIFY_SERVICE_SID = config.GetConfig().TwilioServiceSid
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})

	phone = "+91" + phone
	fmt.Println("REACHED 1")
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(code)
	fmt.Println("REACHED 2")

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)
	fmt.Println("REACHED 3")

	if err != nil {
		return "", fmt.Errorf("Failed to verify otp : %s", err)

	} else if *resp.Status == "approved" {
		return *resp.Status, nil
	} else {
		return "incorrect", nil
	}

}
