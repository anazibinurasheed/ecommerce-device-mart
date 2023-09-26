package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthRequired(c *gin.Context) {
	if !TokenAuth(c, "User") {
		return
	}

}

func AdminAuthRequired(c *gin.Context) {
	if !TokenAuth(c, "SudoAdmin") {
		return
	}
}

func Verified(c *gin.Context) {
	if !TokenAuth(c, "Phone") {
		return
	}
}

func AuthChangePass(c *gin.Context) {
	if !TokenAuth(c, "PassChange") {
		return
	}
}


func TokenAuth(c *gin.Context, name string) bool {
	tokenString, err := c.Cookie(name + "Authorization")

	if err != nil {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User",
		})
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method:%v", token.Header["alg"])
		}
		return []byte(config.GetConfig().JwtSecret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		if float64(time.Now().Unix()) > claims["expires_at"].(float64) {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"StatusCode": 401,
				"msg":        "Jwt session expired",
			})

			return false
		}

		c.Set("userId", fmt.Sprint(claims["userID"]))
		return true
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"Msg":        "Invalid claims",
		})
		return false
	}
}
