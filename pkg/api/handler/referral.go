package handler

import (
	"net/http"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type RefferalHandler struct {
	refferalUseCase services.RefferalUseCase
}

// for wire
func NewRefferalHandler(useCase services.RefferalUseCase) *RefferalHandler {
	return &RefferalHandler{
		refferalUseCase: useCase,
	}
}

// GetRefferalCode godoc
//
//	@Summary		Get the referral code for the current user
//	@Description	Get the referral code of the currently logged-in user
//	@Tags			referral
//	@Produce		json
//	@Success		200	{object}	response.Response	"Success."
//	@Failure		500	{object}	response.Response	"Failed."
//	@Router			/referral/get-code [get]
func (rh *RefferalHandler) GetRefferalCode(c *gin.Context) {
	userID, _ := helper.GetUserIDFromContext(c)
	RefferalCode, err := rh.refferalUseCase.GetUserRefferalCode(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success.", RefferalCode, nil)
	c.JSON(http.StatusOK, response)

}

// ApplyRefferalCode godoc
//
//	@Summary		Apply referral code for referral bonus.
//	@Description	Apply a  referral code  to get wallet money,
//	@Tags			referral
//	@Accept			json
//	@Produce		json
//	@Param			body	body	string	true	"Referral code to apply"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	response.Response	"Success, bonus amount updated in wallet."
//	@Failure		400	{object}	response.Response	"Invalid Input."
//	@Failure		403	{object}	response.Response	"Failed."
//	@Failure		500	{object}	response.Response	"Failed."
//	@Router			/referral/claim [post]
func (rh *RefferalHandler) ApplyRefferalCode(c *gin.Context) {
	var body string
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid Input.", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetUserIDFromContext(c)
	CodeOwnerID, err := rh.refferalUseCase.VerifyRefferalCode(body, userID)

	if err != nil {
		response := response.ResponseMessage(400, "Failed.", nil, err.Error())
		c.JSON(http.StatusForbidden, response)
		return
	}

	err = rh.refferalUseCase.ClaimRefferalBonus(userID, CodeOwnerID)

	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, bonus amount updated in wallet.", nil, nil)
	c.JSON(http.StatusOK, response)

}
