package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	unauthorizedStatus = http.StatusUnauthorized
	statusCode         = "statusCode"
	message            = "message"
)

// AuthMiddleware provides authentication and authorization functionality.
type AuthMiddleware struct {
	userUseCase services.UserUseCase
}

// NewAuthMiddleware creates a new instance of the authentication middleware.
func NewAuthMiddleware(useCase services.UserUseCase) *AuthMiddleware {
	return &AuthMiddleware{userUseCase: useCase}
}

// unauthorized sets an appropriate response for unauthorized access.
func (a *AuthMiddleware) unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		statusCode: unauthorizedStatus,
		message:    "Unauthorized User",
	})
	c.Abort()
}

// tokenExpired sets an appropriate response for expired tokens.
func (a *AuthMiddleware) tokenExpired(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		statusCode: unauthorizedStatus,
		message:    "Token expired",
	})
	c.Abort()
}

// Todo
func (a *AuthMiddleware) checkIsBlockedUser(userID int) (ok bool) {
	userData, err := a.userUseCase.FindUserById(userID)
	if err != nil {
		return false
	}
	return !userData.IsBlocked
}

// tokenAuth checks the user's token for authentication.
func (a *AuthMiddleware) tokenAuth(c *gin.Context, role string) bool {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		a.unauthorized(c)
		return false
	}

	val := strings.SplitN(tokenString, " ", 2)
	if len(val) < 2 {
		a.unauthorized(c)
		return false
	}

	tokenString = val[1]

	token, err := a.parseToken(tokenString)

	if err != nil {
		a.unauthorized(c)
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		if float64(time.Now().Unix()) > claims["expires_at"].(float64) {
			a.tokenExpired(c)
			return false
		}

		if claims["role"] != role {
			a.unauthorized(c)
			return false
		}

		c.Set("userID", fmt.Sprint(claims["userID"]))
		return true
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		statusCode: unauthorizedStatus,
		message:    "Invalid claims",
	})
	c.Abort()
	return false
}

func (a *AuthMiddleware) parseToken(tokenString string) (token *jwt.Token, err error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetConfig().JwtSecret), nil

	})
}

// UserAuthRequired enforces user authentication.
func (a *AuthMiddleware) UserAuthRequired(c *gin.Context) {
	if !a.tokenAuth(c, "user") {
		return
	}
}

// AdminAuthRequired enforces admin authentication.
func (a *AuthMiddleware) AdminAuthRequired(c *gin.Context) {
	if !a.tokenAuth(c, "admin") {
		return
	}
}
