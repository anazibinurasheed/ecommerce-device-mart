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

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(useCase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: useCase}
}

// CreateAdmin godoc
//
//	@Summary		Create admin
//	@Description	Sudo admin to create new admin account.
//	@Tags			sudo admin
//	@Accept			json
//	@Produce		json
//	@Param			body	body	request.SignUpData	true	"Signup data"
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Router			/admin/create-admin [post]
func (ah *AdminHandler) CreateAdmin(c *gin.Context) {
	var body request.SignUpData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//phoneDataMutex and phoneDataMap declared on the top of common.go file .
	//usecase of these variable also mentioned near to the declaration.
	phoneDataMutex.Lock()
	Phone, ok := phoneDataMap[body.UUID]
	phoneDataMutex.Unlock()
	if !ok {
		response := response.ResponseMessage(500, "Failed", nil, fmt.Errorf("Failed to fetch phone number from phoneDataMap").Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	Number, err := strconv.Atoi(Phone)
	if err != nil {
		response := response.ResponseMessage(200, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	body.Phone = Number
	err = ah.adminUseCase.AdminSignUp(body)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//from here phone details from map not needed .
	phoneDataMutex.Lock()
	delete(phoneDataMap, body.UUID)
	phoneDataMutex.Unlock()

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)

}

// SULogin godoc.
//
//	@Summary		Sudo Admin Login
//	@Description	For sudo admin login.
//	@Tags			sudo admin
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.SudoLoginData	true	"Sudo admin login credentials"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		401		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/admin/su-login [post]
func (ah *AdminHandler) SULogin(c *gin.Context) {
	var body request.SudoLoginData
	if err := c.BindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := ah.adminUseCase.SudoAdminLogin(body)
	if err != nil {
		response := response.ResponseMessage(401, "Failed", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	TokenString, _, err := helper.GenerateJwtToken(001) //for su admin
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	MaxAge := int(time.Now().Add(time.Hour * 24 * 30).Unix())
	c.SetCookie("SudoAdminAuthorization", TokenString, MaxAge, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuthorization", TokenString, MaxAge, "", "", false, true) /////////////////
	c.SetSameSite(http.SameSiteLaxMode)

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// ListUsers	godoc
//
//	@Summary		View all users
//	@Description	List of all users
//	@Tags			admin user management
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200										{object}	response.Response	"Success"
//	@Failure		501										{object}	response.Response	"Failed"
//	@Router			/admin/user-management/view-all-users	[get]
func (ah *AdminHandler) DisplayAllUsers(c *gin.Context) {
	// page, err := strconv.Atoi(c.Query("page"))
	// if err != nil {
	// 	response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// count, err := strconv.Atoi(c.Query("count"))
	// if err != nil {
	// 	response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	ListOfAllUserData, err := ah.adminUseCase.GetAllUserData()
	if err != nil {
		response := response.ResponseMessage(501, "Failed", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response := response.ResponseMessage(200, "Success", ListOfAllUserData, nil)
	c.JSON(http.StatusOK, response)
}

// BlockUser godoc
//
//	@Summary		Block a user
//	@Description	Blocks a user with the specified ID.
//	@Tags			admin user management
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int	true	"User ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/admin/user-management/block-user/{userID} [post]
func (ah *AdminHandler) BlockUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ah.adminUseCase.BlockUserByID(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to block user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
	}

	response := response.ResponseMessage(200, "Success, user has been blocked", nil, nil)
	c.JSON(http.StatusOK, response)
}

// UnblockUser godoc
//
//	@Summary		Unblock a user
//	@Description	Unblocks a user with the specified ID.
//	@Tags			admin user management
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int	true	"User ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/admin/user-management/unblock-user/{userID} [post]
func (ah *AdminHandler) UnblockUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ah.adminUseCase.UnBlockUserByID(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to unblock user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, user has been unblocked", nil, nil)
	c.JSON(http.StatusOK, response)
}
