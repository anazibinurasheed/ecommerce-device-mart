package helper

import (
	"fmt"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	UserId   string
	Username string
	IsAdmin  bool
	jwt.StandardClaims
}

// func GenerateJwtToken(UserData response.UserData) (string, error) {

// 	claims := &SignedDetails{
// 		UserId:   fmt.Sprint(UserData.Id),
// 		Username: UserData.UserName,
// 		IsAdmin:  UserData.IsAdmin,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
// 		},
// 	}
// 	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetConfig().JwtSecret))

// 	if err != nil {
// 		log.Println("FAILED TO CREATE SIGNED TOKEN")
// 		return "", err
// 	}
// 	return tokenString, err
// }

func GenerateJwtToken(userId int) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add((time.Hour * 24 * 30)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {
		return "", "", fmt.Errorf("Failed to sign JWT access token :%s", err)
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add((time.Hour * 24 * 40)).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {

		return "", "", fmt.Errorf("Failed to sign JWT refresh token :%s", err)
	}
	return tokenString, refreshTokenString, err
}

// func TokenForSecure(c *gin.Context) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	secureString, err := token.SignedString([]byte(config.GetConfig().JwtSecret))
// 	if err != nil {
// 		log.Println("FAILED TO CREATE SIGNED TOKEN")
// 		return
// 	}
// 	maxAge := int(time.Now().Add(time.Minute * 6).Unix())
// 	c.SetCookie("LoginOtpToken", secureString, maxAge, "", "", false, true)
// 	c.SetSameSite(http.SameSiteLaxMode)

// }
