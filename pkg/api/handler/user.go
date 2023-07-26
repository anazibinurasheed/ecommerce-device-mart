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
	//usecases of these variable also mentioned top of  the declaration.
	phoneDataMutex.Lock()
	Phone, ok := phoneDataMap[body.Id]
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
	delete(phoneDataMap, body.Id)
	phoneDataMutex.Unlock()

	response := response.ResponseMessage(200, "Success, account created", nil, err.Error())
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
//	@Success		200		{object}	response.Response{}
//	@Failure		400		{object}	response.Response{}
//	@Failure		401		{object}	response.Response{}
//	@Failure		500		{object}	response.Response{}
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

	TokenString, RefreshTokenString, err := helper.GenerateJwtToken(UserData.Id)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to generate jwt token", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var CoockieName string

	if UserData.IsAdmin {
		CoockieName = "AdminAuthorization"
	} else {
		CoockieName = "UserAuthorization"
	}

	MaxAge := int(time.Now().Add(time.Hour * 24 * 30).Unix())
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(CoockieName, TokenString, MaxAge, "", "", false, true)
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
	ListOfStates, err := u.userUseCase.DisplayListOfStates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "No states found",
			Data:       nil,
			Error:      err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successful",
		Data:       ListOfStates,
		Error:      nil,
	})

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
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input. ",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	userId, _ := helper.GetUserIdFromContext(c)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, response.Response{
	// 		StatusCode: 500,
	// 		Message:    "Failed to identify user.",
	// 		Data:       nil,
	// 		Error:      err.Error(),
	// 	})
	// 	return
	// }

	err := uh.userUseCase.AddNewAddress(userId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to add address.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successful",
		Data:       nil,
		Error:      nil,
	})
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
		response := response.ResponseMessage(400, "Invalid input.", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	addressID, err := strconv.Atoi(c.Param("addressID"))
	if err != nil {
		response := response.ResponseMessage(400, "Failed,Invalid entry.", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userId, _ := helper.GetUserIdFromContext(c)
	err = uh.userUseCase.UpdateUserAddress(body, addressID, userId)

	if err != nil {
		response := response.ResponseMessage(500, "Update address failed .", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.ResponseMessage(200, "Success.", nil, nil)
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
	if err != nil || addressID == 0 {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Invalid input",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err = uh.userUseCase.DeleteUserAddress(addressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Delete address failed .",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Seccusfull.",
		Data:       nil,
		Error:      nil,
	})
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
func (uh *UserHandler) GetAllAdresses(c *gin.Context) {
	userId, err := helper.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to identify user.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	ListOfAddresses, err := uh.userUseCase.GetUserAddresses(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Succesful.",
		Data:       ListOfAddresses,
		Error:      nil,
	})

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
	userId, err := helper.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to identify user.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	UserProfile, err := uh.userUseCase.GetProfile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Succesful.",
		Data:       UserProfile,
		Error:      nil,
	})

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
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input. ",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	userId, err := helper.GetUserIdFromContext(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to identify user.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err = uh.userUseCase.CheckUserOldPassword(body, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to change user password .",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	TokenString, _, err := helper.GenerateJwtToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to generate jwt token",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	MaxAge := int(time.Now().Add(time.Minute * 30).Unix())
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("PassChangeAuthorization", TokenString, MaxAge, "", "", false, true)

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Seccusfull.",
		Data:       nil,
		Error:      nil,
	})

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
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input. ",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	userId, err := helper.GetUserIdFromContext(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to identify user.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err = uh.userUseCase.ChangeUserPassword(body, userId, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Change password failed .",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	helper.DeleteCookie("PassChangeAuth", c)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successful,password changed .",
		Data:       nil,
		Error:      nil,
	})

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
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	userID, _ := helper.GetUserIdFromContext(c)
	err = uh.userUseCase.SetDefaultAddress(userID, addressID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
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
	var body string
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input. ",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	userId, _ := helper.GetUserIdFromContext(c)

	err := uh.userUseCase.UpdateUserName(body, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    " failed .",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successful,username changed .",
		Data:       nil,
		Error:      nil,
	})

}
