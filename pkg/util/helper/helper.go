package helper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// it's a util function for pagination
func Paginate(page, count int) (startIndex, endIndex int) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}

	startIndex = (page - 1) * count
	endIndex = startIndex + count
	return
}

func GetIDFromContext(c *gin.Context) (int, error) {
	userIDStr := c.GetString("userID")
	userID, err := strconv.Atoi(userIDStr)
	return userID, err
}

func SetToCookie(Data int, cookieName string, c *gin.Context) {

	maxAge := int(time.Now().Add(time.Minute * 6).Unix())
	c.SetCookie(cookieName, fmt.Sprint(Data), maxAge, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
}

func DeleteCookie(cookieName string, c *gin.Context) {

	c.SetCookie(cookieName, "", -1, "", "", false, true)
}

func GenerateUniqueID() string {
	id := uuid.New().String()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s-%d", id, timestamp)
}

func MakeSKU(name string) string {
	return strings.ReplaceAll(name, " ", "-")
}

func CalculateTotalRevenue(args ...response.OrderLine) float64 {

	return func() (totalRevenue float64) {

		for _, orders := range args {
			totalRevenue += float64(orders.Qty) * float64(orders.Price)
		}

		return
	}()
}

// Validates the struct using validate tags, if the validation fails it will return false else true
func ValidateStruct(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}

// ValidateData will validate the struct fields according to validate tags and if ok it will return true else it will return false. And it will also write appropriate response.(swag:status 417, util.response)
func ValidateData(c *gin.Context, obj any) bool {

	if err := ValidateStruct(obj); err != nil {
		response := response.ResponseMessage(http.StatusExpectationFailed, "failed, not enough credentials", nil, err.Error())
		c.AbortWithStatusJSON(http.StatusExpectationFailed, &response)
		return false
	}
	return true

}
