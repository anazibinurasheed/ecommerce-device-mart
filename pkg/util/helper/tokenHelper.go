package helper

import (
	"fmt"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/golang-jwt/jwt"
)

const (
	userID    = "userID"
	expiresAt = "expires_at"
	role      = "role"
)

func GenerateToken(userId int, roleName string) (tokenString string, err error) {
	maxAge := time.Now().Add((time.Hour * 24 * 30)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userID:    userId,
		expiresAt: maxAge,
		role:      roleName,
	})

	tokenString, err = token.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign access token :%s", err)
	}

	return
}
