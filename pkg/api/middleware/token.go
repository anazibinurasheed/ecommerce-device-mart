package middleware

import (
	"fmt"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/gin-gonic/gin"
)

// TokenManager provides token-related functionality.
type TokenManager struct{}

// NewTokenManager creates a new TokenManager instance.
func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

// GenerateAdminToken generates a token for Admin.
func (t *TokenManager) GenerateAdminToken() (token string, err error) {
	role := "admin"
	userID := 0
	token, err = helper.GenerateToken(userID, role)

	return
}

// GenerateUserToken generates a token for the given user ID.
func (t *TokenManager) GenerateUserToken(userID int) (token string, err error) {
	role := "user"
	token, err = helper.GenerateToken(userID, role)

	return
}

// SetTokenHeader sets the token in the Authorization header of the request.
func (t *TokenManager) SetTokenHeader(c *gin.Context, token string) {
	key := "Authorization"
	c.Request.Header.Set(key, token)
	fmt.Println(c.Request.Header.Get(key))
}

func (t *TokenManager) RemoveToken(c *gin.Context) {
	key := "Authorization"

	c.Request.Header.Set(key, "")

}
