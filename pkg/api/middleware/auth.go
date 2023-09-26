package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthenticateUserJwt(c *gin.Context) {
	JwtAuth(c, "User")

}

func AuthenticateAdminJwt(c *gin.Context) {
	JwtAuth(c, "Admin")
}

func AuthenticateSudoAdminJwt(c *gin.Context) {
	JwtAuth(c, "SudoAdmin")
}

func Verified(c *gin.Context) {
	JwtAuth(c, "Phone")
}

func AuthChangePass(c *gin.Context) {
	JwtAuth(c, "PassChange")
}

// for admin routes
func AdminAuthJWT(c *gin.Context) {
	if _, err := c.Cookie("AdminAuthorization"); err == nil {
		AuthenticateAdminJwt(c)
	} else if _, err := c.Cookie("SudoAdminAuthorization"); err == nil {
		AuthenticateSudoAdminJwt(c)
	} else {
		JwtAuth(c, "") //for requests without any required token
	}

}

func JwtAuth(c *gin.Context, name string) {
	tokenString, err := c.Cookie(name + "Authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User",
		})
		return
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
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"StatusCode": 401,
				"msg":        "Jwt session expired",
			})
			return
		}

		c.Set("userId", fmt.Sprint(claims["sub"]))
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"Msg":        "Invalid claims",
		})
		return
	}
}
