package helper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIdFromContext(c *gin.Context) (int, error) {
	userIdStr := c.GetString("userId")
	userId, err := strconv.Atoi(userIdStr)

	return userId, err
}

func SetToCookie(Data int, cookieName string, c *gin.Context) {

	maxAge := int(time.Now().Add(time.Minute * 6).Unix())
	c.SetCookie(cookieName, fmt.Sprint(Data), maxAge, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
}

func GetFromCookie(cookieName string, c *gin.Context) (int, error) {
	cookieData, err := c.Cookie(cookieName)
	if err != nil {
		fmt.Println("FAILED TO GET ID FROM COOKIE", err)
		return 0, err
	}

	data, err := strconv.Atoi(cookieData)

	return data, err

}

func DeleteCookie(cookieName string, c *gin.Context) {

	c.SetCookie(cookieName, "", -1, "", "", false, true)
}

func GenerateUniqueID() string {
	id := uuid.New().String()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s-%d", id, timestamp)
}

func MakeSku(name string) string {
	name = strings.ReplaceAll(name, " ", "-")
	return name
}