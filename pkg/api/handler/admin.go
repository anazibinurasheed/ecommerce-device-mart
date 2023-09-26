package handler

import (
	"net/http"
	"strconv"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
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
//	@Router			/admin/user-management/block-user/{userID} [put]
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
//	@Router			/admin/user-management/unblock-user/{userID} [put]
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
