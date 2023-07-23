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
