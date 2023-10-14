package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/api/auth"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase services.AuthUseCase
	token       auth.TokenManager
}

func NewAuthHandler(useCase services.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: useCase,
	}
}

var contact = helper.NewPhone()

// SULogin godoc.
//
//	@Summary		Admin Login
//	@Description	For admin login.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.SudoLoginData	true	"Sudo admin login credentials"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		401		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/admin/su-login [post]
func (ah *AuthHandler) SULogin(c *gin.Context) {
	var body request.SudoLoginData
	if err := c.BindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := ah.authUseCase.SudoAdminLogin(body)
	if err != nil {
		response := response.ResponseMessage(401, "Failed", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	TokenString, err := ah.token.GenerateAdminToken()
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	ah.token.SetTokenHeader(c, TokenString)

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

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

	if !helper.ValidateData(c, &body) {
		return
	}

	phone, err := ch.authUseCase.ValidateSignUpRequest(body)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uuid := helper.GenerateUniqueID()
	contact.Set(uuid, fmt.Sprint(phone))

	go contact.Clean(uuid)

	go func() {
		time.Sleep(65 * time.Second)
		contact.Print(uuid)
		fmt.Println(contact)
	}()

	response := response.ResponseMessage(202, "Success, otp sended.The otp will be expire within 1 minute.", uuid, nil)
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
func (ah *AuthHandler) VerifyOTP(c *gin.Context) {
	var body request.Otp

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	phone, ok, _ := contact.Get(body.UUID)
	if ok {
		contact.Verified(body.UUID, phone)
	}

	if !ok {
		response := response.ResponseMessage(500, "Failed", nil, fmt.Errorf("otp expired").Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	status, err := helper.CheckOtp(phone, body.Otp)
	if err != nil {

		contact.NotVerified(body.UUID, phone)

		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if status == "incorrect" {
		response := response.ResponseMessage(400, "Incorrect otp", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

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

	if !helper.ValidateData(c, &body) {
		return
	}

	phoneStr, ok, verified := contact.Get(body.Uuid)

	switch {
	case !ok:
		response := response.ResponseMessage(401, "Failed.", nil, fmt.Errorf("otp not verified").Error())
		c.JSON(http.StatusUnauthorized, response)
		return

	case !verified:
		response := response.ResponseMessage(401, "Failed.", nil, fmt.Errorf("invalid try, user not verified otp").Error())
		c.JSON(http.StatusUnauthorized, response)
		return

	}

	phone, err := strconv.Atoi(phoneStr)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return

	}

	body.Phone = phone

	err = u.authUseCase.SignUp(body)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	contact.Delete(body.Uuid)

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

	if !helper.ValidateData(c, &body) {
		return
	}

	UserData, err := uh.authUseCase.ValidateUserLoginCredentials(body)
	if err != nil {
		response := response.ResponseMessage(401, "Failed", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	TokenString, err := uh.token.GenerateUserToken(UserData.ID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to generate jwt token", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	uh.token.SetTokenHeader(c, TokenString)

	response := response.ResponseMessage(200, "Login success", gin.H{"token": TokenString}, nil)

	c.JSON(http.StatusOK, response)
}

// @Summary		User Logout
// @Description	Logs out user and remove cookie from browser.
// @Tags			auth
// @Security		JWT
// @Accept			json
// @Produce		json
// @Success		202	{object}	response.Response{}
// @Router			/logout [post]
func (ah *AuthHandler) Logout(c *gin.Context) {
	ah.token.RemoveToken(c)

	response := response.ResponseMessage(200, "Logged out, success", nil, nil)
	c.JSON(http.StatusAccepted, response)
}
