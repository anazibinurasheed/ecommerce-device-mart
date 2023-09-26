package helper

import (
	"fmt"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     userId,
		"expires_at": time.Now().Add((time.Hour * 24 * 30)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {
		return "", fmt.Errorf("Failed to sign JWT access token :%s", err)
	}

	return tokenString, err
}
