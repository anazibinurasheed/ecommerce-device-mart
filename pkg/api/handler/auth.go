package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase services.AuthUseCase
}

func NewAuthHandler(useCase services.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: useCase,
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
//	@Description	Sends an OTP to the provided phone number. Should take the uuid and verify the otp using verify otp api then take the uuid and include it also in the sign up credentials. Else will not work
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Phone	true	"Phone number"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Router			/send-otp [post]
func (ch *AuthHandler) SendOTP(c *gin.Context) {
	var body request.Phone
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	Phone, err := ch.authUseCase.ValidateSignUpRequest(body)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	phoneDataMutex.Lock()
	uid := helper.GenerateUniqueID()
	phoneDataMap[uid] = fmt.Sprint(Phone)
	phoneDataMutex.Unlock()

	go helper.GoClean(phoneDataMap, uid, phoneDataMutex)

	go func() {
		time.Sleep(65 * time.Second)
		phoneDataMutex.Lock()
		fmt.Println(phoneDataMap)
		phoneDataMutex.Unlock()
	}()

	response := response.ResponseMessage(202, "Success, otp sended.The otp will be expire within 1 minute.", uid, nil)
	c.JSON(http.StatusAccepted, response)
}

// VerifyOTP godoc
//
//	@Summary		Verify sign up  OTP
//	@Description	Validates the provided OTP for a phone number. Provide the accurate uuid and otp = 0000(predefined).
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Otp	true	"OTP"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		401		{object}	response.Response
//	@Router			/verify-otp [post]
func (ch *AuthHandler) VerifyOTP(c *gin.Context) {
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
		response := response.ResponseMessage(500, "Failed", nil, fmt.Errorf("otp expired").Error())
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

// UserSignUp is the handler function for user sign-up.
//
//	@Summary		User Sign-Up after otp validation
//	@Description	Creates a new user account.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.SignUpData	true	"User Sign-Up Data"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Router			/sign-up [post]
func (u *AuthHandler) UserSignUp(c *gin.Context) {
	var body request.SignUpData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//phoneDataMutex and phoneDataMap declared on the top of common.go file .
	//use of these variable also mentioned near to the declaration.
	phoneDataMutex.Lock()
	Phone, ok := phoneDataMap[body.UUID]
	phoneDataMutex.Unlock()
	if !ok {
		response := response.ResponseMessage(500, "Failed.", nil, fmt.Errorf("failed to fetch phone number from phoneDataMap").Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	Number, err := strconv.Atoi(Phone)
	if err != nil {
		response := response.ResponseMessage(200, "Failed.", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	body.Phone = Number

	err = u.authUseCase.SignUp(body)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	phoneDataMutex.Lock()
	delete(phoneDataMap, body.UUID)
	phoneDataMutex.Unlock()

	response := response.ResponseMessage(200, "Success, account created", nil, nil)
	c.JSON(http.StatusOK, response)
}

// UserLogin godoc
//
//	@Summary		User login data, verify it and send otp
//	@Description	Logs in a user and sends an OTP for verification.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.LoginData	true	"User login data"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		401		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/login [post]
func (uh *AuthHandler) UserLogin(c *gin.Context) {
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	UserData, err := uh.authUseCase.ValidateUserLoginCredentials(body)
	if err != nil {
		response := response.ResponseMessage(401, "Failed", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	TokenString, RefreshTokenString, err := helper.GenerateJwtToken(UserData.ID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to generate jwt token", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var coockieName string

	if UserData.IsAdmin {
		coockieName = "AdminAuthorization"
	} else {
		coockieName = "UserAuthorization"
	}

	MaxAge := int(time.Now().Add(time.Hour * 24 * 30).Unix())
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(coockieName, TokenString, MaxAge, "", "", false, true)
	c.SetCookie("RefreshToken", RefreshTokenString, MaxAge, "", "", false, true)

	response := response.ResponseMessage(200, "Login success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// @Summary		User Logout
// @Description	Logs out user and remove cookie from browser.
// @Tags			auth
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Response{}
// @Router			/logout [post]
func (uh *AuthHandler) Logout(c *gin.Context) {

	helper.DeleteCookie("AdminAuthorization", c)
	helper.DeleteCookie("SudoAdminAuthorization", c)
	helper.DeleteCookie("UserAuthorization", c)

	response := response.ResponseMessage(200, "Logged out, success", nil, nil)
	c.JSON(http.StatusAccepted, response)
}
