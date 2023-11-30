package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase services.AuthUseCase
	token       middleware.TokenManager
	subHandler  helper.SubHandler
}

func NewAuthHandler(useCase services.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: useCase,
	}
}

var contact = helper.NewPhone()

// AdminLogin godoc.
//
//	@Summary		Admin Login
//	@Description	Admin can login using username and password.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.AdminLogin						true	"Admin login credentials"
//	@Success		200		{object}	response.Response{data=response.Token}	"Login success"
//	@Failure		400		{object}	response.Response						"Failed to bind JSON inputs from request"
//	@Failure		400		{object}	response.Response						"Failed, input does not meet validation criteria"
//	@Failure		401		{object}	response.Response						"Invalid credentials"
//	@Failure		500		{object}	response.Response						"Failed to generate token"
//	@Router			/admin/login [post]
func (a *AuthHandler) AdminLogin(c *gin.Context) {
	var body request.AdminLogin
	if !a.subHandler.BindRequest(c, &body) {
		return
	}

	err := a.authUseCase.AdminLogin(body)
	if err != nil {
		response := response.ResponseMessage(statusUnauthorized, "Invalid credentials", nil, err.Error())
		c.JSON(statusUnauthorized, response)
		return
	}

	TokenString, err := a.token.GenerateAdminToken()
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to generate token", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	a.token.SetTokenHeader(c, TokenString)
	token := &response.Token{Tkn: TokenString}

	response := response.ResponseMessage(statusOK, "Login success", token, nil)
	c.JSON(statusOK, response)
}

// SendOTP godoc
//
//	@Summary		Send sign up OTP to Phone
//	@Description	Sends an OTP to the provided phone number. Should take the uuid and verify the otp using verify otp api then take the uuid and include it also in the sign up credentials. Else will not work
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Phone	true	"Phone number"
//	@Success		200		{object}	response.Response{data=response.Uuid}
//	@Failure		400		{object}	response.Response	"Failed to bind JSON inputs from request"
//	@Failure		400		{object}	response.Response	"Failed, input does not meet validation criteria"
//	@Router			/send-otp [post]
func (a *AuthHandler) SendOTP(c *gin.Context) {
	var body request.Phone
	if !a.subHandler.BindRequest(c, &body) {
		return
	}

	phone, err := a.authUseCase.ValidateSignUpRequest(body)
	if err != nil {
		response := response.ResponseMessage(statusBadRequest, "Failed", nil, err.Error())
		c.JSON(statusBadRequest, response)
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

	data := response.Uuid{Uuid: uuid}
	response := response.ResponseMessage(statusOK, "Success, otp sended.The otp will be expire within 3 minute.", data, nil)
	c.JSON(statusOK, response)
}

// VerifyOTP godoc
//
//	@Summary		Verify sign up  OTP
//	@Description	Validates the provided OTP for a phone number. Provide the accurate uuid and otp = 0000(predefined).
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Otp								true	"OTP"
//	@Success		200		{object}	response.Response{data=response.Uuid}	"Success, verified phone number"
//	@Failure		400		{object}	response.Response						"Failed to bind JSON inputs from request"
//	@Failure		400		{object}	response.Response						"Failed, input does not meet validation criteria"
//	@Failure		400		{object}	response.Response						"Failed to verify otp"
//	@Failure		400		{object}	response.Response						"Incorrect otp"
//	@Failure		500		{object}	response.Response						"OTP expired"
//	@Router			/verify-otp [post]
func (a *AuthHandler) VerifyOTP(c *gin.Context) {
	var body request.Otp
	if !a.subHandler.BindRequest(c, &body) {
		return
	}

	phone, ok, _ := contact.Get(body.UUID)
	if ok {
		contact.Verified(body.UUID, phone)
	}

	if !ok {
		response := response.ResponseMessage(statusInternalServerError, "OTP expired", nil, "unable to find phone number")
		c.JSON(statusInternalServerError, response)
		return
	}

	status, err := helper.CheckOtp(phone, body.Otp)

	if err != nil {
		contact.NotVerified(body.UUID, phone)
		response := response.ResponseMessage(statusBadRequest, "Failed to verify otp", nil, err.Error())
		c.JSON(statusBadRequest, response)
		return
	}

	if status == "incorrect" {
		response := response.ResponseMessage(statusBadRequest, "Incorrect otp", nil, nil)
		c.JSON(statusBadRequest, response)
		return
	}

	data := response.Uuid{
		Uuid: body.UUID,
	}
	response := response.ResponseMessage(statusOK, "Success, verified phone number", data, nil)
	c.JSON(statusOK, response)
}

// UserSignUp is the handler function for user sign-up.
//
//	@Summary		User Sign-Up after otp validation
//	@Description	Creates a new user account.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.SignUpData	true	"User Sign-Up Data"
//	@Success		201		{object}	response.Response	"Success, account created"
//
//	@Failure		400		{object}	response.Response	"Failed to bind JSON inputs from request"
//
//	@Failure		400		{object}	response.Response	"Failed, input does not meet validation criteria"
//	@Router			/sign-up [post]
func (u *AuthHandler) UserSignUp(c *gin.Context) {
	var body request.SignUpData
	if !u.subHandler.BindRequest(c, &body) {
		return
	}

	phoneStr, ok, verified := contact.Get(body.Uuid)

	switch {
	case !ok:
		response := response.ResponseMessage(statusUnauthorized, "User not verified OTP", nil, "phone not found")
		c.JSON(statusUnauthorized, response)
		return

	case !verified:
		response := response.ResponseMessage(statusUnauthorized, "Failed not verified OTP", nil, "invalid try, user not verified otp")
		c.JSON(statusUnauthorized, response)
		return

	}

	phone, err := strconv.Atoi(phoneStr)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed unable to convert type", nil, err.Error())
		c.JSON(statusInternalServerError, response)
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

	response := response.ResponseMessage(statusCreated, "Success, account created", nil, nil)
	c.JSON(statusCreated, response)
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
// @Description	Logs out user and removes token from the header.
// @Security		Bearer
// @Tags			auth
// @Accept			json
// @Produce		json
// @Success		202	{object}	response.Response	"Logged out, success"
// @Router			/logout [post]
func (ah *AuthHandler) Logout(c *gin.Context) {
	ah.token.RemoveToken(c)

	response := response.ResponseMessage(statusAccepted, "Log out, success", nil, nil)
	c.JSON(statusAccepted, response)
}
