package handler

import (
	"net/http"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type ReferralHandler struct {
	referralUseCase services.ReferralUseCase
}

// for wire
func NewReferralHandler(useCase services.ReferralUseCase) *ReferralHandler {
	return &ReferralHandler{
		referralUseCase: useCase,
	}
}

// GetReferralCode godoc
//
//	@Summary		Get the referral code for the current user
//	@Description	Get the referral code of the currently logged-in user
//	@Tags			referral
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response	"Success."
//	@Failure		500	{object}	response.Response	"Failed."
//	@Router			/referral/get-code [get]
func (rh *ReferralHandler) GetReferralCode(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)
	referralCode, err := rh.referralUseCase.GetUserReferralCode(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success.", referralCode, nil)
	c.JSON(http.StatusOK, response)

}

// ApplyReferralCode godoc
//
//	@Summary		Apply referral code for referral bonus.
//	@Description	Apply a  referral code  to get wallet money,
//	@Tags			referral
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			body	body	string	true	"Referral code to apply"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	response.Response	"Success, bonus amount updated in wallet."
//	@Failure		400	{object}	response.Response	"Invalid Input."
//	@Failure		403	{object}	response.Response	"Failed."
//	@Failure		500	{object}	response.Response	"Failed."
//	@Router			/referral/claim [post]
func (rh *ReferralHandler) ApplyReferralCode(c *gin.Context) {
	var body string
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid Input.", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetIDFromContext(c)
	codeOwnerID, err := rh.referralUseCase.VerifyReferralCode(body, userID)

	if err != nil {
		response := response.ResponseMessage(400, "Failed.", nil, err.Error())
		c.JSON(http.StatusForbidden, response)
		return
	}

	err = rh.referralUseCase.ClaimReferralBonus(userID, codeOwnerID)

	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, bonus amount updated in wallet.", nil, nil)
	c.JSON(http.StatusOK, response)

}
