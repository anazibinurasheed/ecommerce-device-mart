package handler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type CommonHandler struct {
	commonUseCase services.CommonUseCase
}

// for wire
func NewCommonHandler(useCase services.CommonUseCase) *CommonHandler {
	return &CommonHandler{
		commonUseCase: useCase,
	}
}

var (
	//phoneDataMap
	// The phoneDataMap is a map used to store users' phone numbers retrieved from an API.
	// The stored phone number will be used for OTP verification and to fill up sign up credentials without asking the user to enter
	// the phone number again.
	// It is stored in the map with a unique key.
	// The unique key will be passed to the frontend.
	//From the frontend, the key will then be passed to the next API that requires the phone number.
	// Once the user completes all the authentication steps, the phone number will be deleted from the phoneDataMap, and the phone number,
	//along with other user sign up credentials, will be inserted into the database.
	//Theme : To decrease the amount of database operations
	phoneDataMap = make(map[string]string)
	//Here we are using normal map instead of sync.Map so we should ensure  not to come  race condition .
	//phoneDataMutex is for preventing from race condition.
	phoneDataMutex = new(sync.Mutex)
)

// SendOTP godoc
//
//	@Summary		Send sign up OTP to Phone
//	@Description	Sends an OTP to the provided phone number.
//	@Tags			common
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Phone	true	"Phone number"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Router			/send-otp [post]
func (ch *CommonHandler) SendOTP(c *gin.Context) {
	var body request.Phone
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	Phone, err := ch.commonUseCase.ValidateSignUpRequest(body)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	phoneDataMutex.Lock()
	uid := helper.GenerateUniqueID()
	phoneDataMap[uid] = fmt.Sprint(Phone)
	phoneDataMutex.Unlock()

	response := response.ResponseMessage(202, "Success, otp sended", uid, nil)
	c.JSON(http.StatusAccepted, response)
}

// VerifyOTP godoc
//
//	@Summary		Verify sign up  OTP
//	@Description	Validates the provided OTP for a phone number.
//	@Tags			common
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Otp	true	"OTP"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		401		{object}	response.Response
//	@Router			/verify-otp [post]
func (ch *CommonHandler) VerifyOTP(c *gin.Context) {
	var body request.Otp

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	phoneDataMutex.Lock()
	number, ok := phoneDataMap[body.UUID]
	phoneDataMutex.Unlock()
	if !ok {
		response := response.ResponseMessage(500, "Failed", nil, fmt.Errorf("failed to fetch phone number from phoneDataMap").Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	status, err := helper.CheckOtp(number, body.Otp)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if status == "incorrect" {
		response := response.ResponseMessage(400, "Incorrect otp", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tokenString, _, _ := helper.GenerateJwtToken(0)
	maxAge := int(time.Now().Add(time.Minute * 30).Unix())
	c.SetCookie("PhoneAuthorization", tokenString, maxAge, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)

	response := response.ResponseMessage(202, "Success, verified phone number", body.UUID, nil)
	c.JSON(http.StatusAccepted, response)
}

//	@Summary		User Logout
//	@Description	Logs out user and remove cookie from browser.
//	@Tags			common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response{}
//	@Router			/logout [post]
func (uh *CommonHandler) Logout(c *gin.Context) {
	helper.DeleteCookie("AdminAuthorization", c)
	helper.DeleteCookie("SudoAdminAuthorization", c)
	helper.DeleteCookie("UserAuthorization", c)

	response := response.ResponseMessage(200, "Logged out, success", nil, nil)
	c.JSON(http.StatusAccepted, response)
}

// func (uh *CommonHandler) RefreshToken(c *gin.Context) {
// 	refreshToken, err := c.Cookie("RefreshToken")

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 			"StatusCode": 401,
// 			"msg":        "Unauthorized User",
// 		})
// 		return
// 	}

// 	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 		if !ok {
// 			return nil, fmt.Errorf("Unexpected signing method:%v", token.Header["alg"])
// 		}
// 		return []byte(config.GetConfig().JwtSecret), nil
// 	})

// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	if !ok && !token.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"Statuscode": 401,
// 			"Msg":        "Invalid claims",
// 		})
// 	}

// 	if claims["exp"].(float64) > float64(time.Now().Add(time.Minute*60).Unix()) {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": "Allocated refresh time expired",
// 		})
// 		return
// 	}

// 	userID, _ := helper.GetUserIDFromContext(c)
// 	tokenString, _, err := helper.GenerateJwtToken(userID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"msg": "Failed to generate access token",
// 		})
// 		return
// 	}
// 	MaxAge := time.Now().Add((time.Hour * 24 * 30)).Unix()
// 	c.SetSameSite(http.SameSiteLaxMode)
// 	c.SetCookie("UserAuthorization", tokenString, int(MaxAge), "", "", false, true)
// }
