package helper

import (
	"math/rand"
	"time"
)

func GenerateReferralCode() string {

	codeLength := 15
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	code := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		code[i] = chars[rand.Intn(len(chars))]
	}

	return string(code)
}
