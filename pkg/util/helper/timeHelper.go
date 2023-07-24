package helper

import (
	"fmt"
	"time"
)

func GetDateAndTimeFromUnix(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	date := t.Format("2006-01-02")
	timeFormatted := t.Format("03:04 PM")

	return fmt.Sprintf("%s - %s", date, timeFormatted)
}

func SetTime(days int) time.Time {
	return time.Now().AddDate(0, 0, days)

}

func IsCouponValid(expiration time.Time) bool {
	// Get the current date and time
	now := time.Now()

	// Compare the current date with the expiration date
	return now.Before(expiration) || now.Equal(expiration)
}

func IsValidReturn(orderPlacedTime time.Time) bool {
	returnPeriodDuration := 7 * 24 * time.Hour // 7 days

	// Get the current time
	currentTime := time.Now()

	// Calculate the time difference between the current time and the order placed time
	timeDifference := currentTime.Sub(orderPlacedTime)

	// Check if the time difference is greater than or equal to the return period duration
	return timeDifference <= returnPeriodDuration
}
