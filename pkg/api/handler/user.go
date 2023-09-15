package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

// for wire
func NewUserHandler(useCase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: useCase,
	}
}

// UserSignUp is the handler function for user sign-up.
//
//	@Summary		User Sign-Up after otp validation
//	@Description	Creates a new user account.
//	@Tags			common
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.SignUpData	true	"User Sign-Up Data"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Router			/sign-up [post]
func (u *UserHandler) UserSignUp(c *gin.Context) {
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
		response := response.ResponseMessage(500, "Failed.", nil, fmt.Errorf("Failed to fetch phone number from phoneDataMap").Error())
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

	err = u.userUseCase.SignUp(body)
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
//	@Tags			common
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.LoginData	true	"User login data"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		401		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/login [post]
func (uh *UserHandler) UserLogin(c *gin.Context) {
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	UserData, err := uh.userUseCase.ValidateUserLoginCredentials(body)
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

// GetAddAddressPage godoc
//
//	@Summary		Get the page for adding an address
//	@Description	Retrieves the list of states for address selection.
//	@Tags			user profile
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/profile/add-address [get]
func (u *UserHandler) GetAddAddressPage(c *gin.Context) {
	listOfStates, err := u.userUseCase.DisplayListOfStates()
	if err != nil {
		response := response.ResponseMessage(500, "No states found", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", listOfStates, nil)
	c.JSON(http.StatusOK, response)
}

// AddAddress godoc
//
//	@Summary		Add a new address
//	@Description	Adds a new address for the user.
//	@Tags			user profile
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Address	true	"Address details"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/profile/add-address [post]
func (uh *UserHandler) AddAddress(c *gin.Context) {
	var body request.Address
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetUserIDFromContext(c)

	err := uh.userUseCase.AddNewAddress(userId, body)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to add address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// UpdateAddress godoc
//
//	@Summary		Update an address
//	@Description	Updates an existing address for the user.
//	@Tags			user profile
//	@Accept			json
//	@Produce		json
//	@Param			addressID	path		int				true	"Address ID"
//	@Param			body		body		request.Address	true	"Address updation details"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/profile/update-address/{addressID} [put]
func (uh *UserHandler) UpdateAddress(c *gin.Context) {
	var body request.Address
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	addressID, err := strconv.Atoi(c.Param("addressID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetUserIDFromContext(c)

	err = uh.userUseCase.UpdateUserAddress(body, addressID, userId)
	if err != nil {
		response := response.ResponseMessage(500, "Update address failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusBadRequest, response)
}

// DeleteAddress godoc
//
//	@Summary		Delete an address
//	@Description	Deletes an address by its ID.
//	@Tags			user profile
//	@Produce		json
//	@Param			addressID	path		int	true	"Address ID"
//	@Success		200			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/profile/delete-address/{addressID} [delete]
func (uh *UserHandler) DeleteAddress(c *gin.Context) {
	addressID, err := strconv.Atoi(c.Param("addressID"))
	if err != nil {
		response := response.ResponseMessage(500, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = uh.userUseCase.DeleteUserAddress(addressID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to delete address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// GetAllAddresses godoc
//
//	@Summary		Get all addresses
//	@Description	Retrieves all addresses for the user.
//	@Tags			user profile
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/profile/addresses [get]
func (uh *UserHandler) GetAllAddresses(c *gin.Context) {
	userId, _ := helper.GetUserIDFromContext(c)

	ListOfAddresses, err := uh.userUseCase.GetUserAddresses(userId)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", ListOfAddresses, nil)
	c.JSON(http.StatusOK, response)
}

// Profile godoc
//
//	@Summary		Get user profile
//	@Description	Retrieves the profile information for the authenticated user.
//	@Tags			user profile
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/profile [get]
func (uh *UserHandler) Profile(c *gin.Context) {
	userId, _ := helper.GetUserIDFromContext(c)

	UserProfile, err := uh.userUseCase.GetProfile(userId)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", UserProfile, nil)
	c.JSON(http.StatusOK, response)
}

// ChangePasswordRequest handles the request to change user password.
//
//	@Summary		Change user password request
//	@Description	validate the user password based on the provided old password and give access to password change api.
//	@Tags			user profile
//	@Accept			json
//	@Produce		json
//	@Param			body	body	request.OldPassword	true	"User's old password"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/profile/verify-password [post]
func (uh *UserHandler) ChangePasswordRequest(c *gin.Context) {
	var body request.OldPassword
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(500, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetUserIDFromContext(c)

	err := uh.userUseCase.CheckUserOldPassword(body, userId)
	if err != nil {
		response := response.ResponseMessage(400, "Failed to change user password", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	TokenString, _, err := helper.GenerateJwtToken(userId)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to generate jwt token", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	MaxAge := int(time.Now().Add(time.Minute * 30).Unix())
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("PassChangeAuthorization", TokenString, MaxAge, "", "", false, true)

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// ChangePassword is used to change the password of the authenticated user.
//
//	@Summary		Change user password
//	@Description	Change the password of the authenticated user
//	@Tags			user profile
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.ChangePassword	true	"Change password request body"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/profile/change-password [post]
func (uh *UserHandler) ChangePassword(c *gin.Context) {
	var body request.ChangePassword
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetUserIDFromContext(c)

	err := uh.userUseCase.ChangeUserPassword(body, userId, c)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to change password", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	helper.DeleteCookie("PassChangeAuth", c)

	response := response.ResponseMessage(200, "Success, password changed", nil, nil)
	c.JSON(http.StatusOK, response)
}

// SetDefaultAddress is the handler function for setting an address as the default address for the user.
//
//	@Summary		Set default address
//	@Description	Sets the specified address as the default address for the authenticated user.
//	@Tags			user profile
//	@Produce		json
//	@Param			addressID	path		int	true	"Address ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/profile/address-default/{addressID} [put]
func (uh *UserHandler) SetDefaultAddress(c *gin.Context) {
	addressID, err := strconv.Atoi(c.Param("addressID"))
	if err != nil || addressID == 0 {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetUserIDFromContext(c)

	err = uh.userUseCase.SetDefaultAddress(userID, addressID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to set address to default", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, updated default address", nil, nil)
	c.JSON(http.StatusOK, response)
}

// EditUserName is used to edit the username of the authenticated user.
//
//	@Summary		Edit user username
//	@Description	Edit the username of the authenticated user
//	@Tags			user profile
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			body	body		string	true	"New username"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/profile/edit-username [post]
func (uh *UserHandler) EditUserName(c *gin.Context) {
	var body request.Name
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetUserIDFromContext(c)

	err := uh.userUseCase.UpdateUserName(body.Name, userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to update username", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "success, username has been changed", nil, nil)
	c.JSON(http.StatusOK, response)
}
