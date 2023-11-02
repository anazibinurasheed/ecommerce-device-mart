package handler

import (
	"net/http"
	"strconv"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	coupenUseCase services.CouponUseCase
	subHandler    helper.SubHandler
}

// for wire
func NewCouponHandler(useCase services.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		coupenUseCase: useCase,
	}
}

// CreateCoupon creates a new coupon.
//
//	@Summary		Create a new coupon
//	@Description	Create a new coupon with the provided details
//	@Tags			promotions
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Coupon		true	"Coupon details"
//	@Success		201		{object}	response.Response	"Success, created new coupon"
//	@Failure		400		{object}	response.Response	"Failed to bind JSON inputs from request"
//	@Failure		400		{object}	response.Response	"Failed, input does not meet validation criteria"
//	@Failure		500		{object}	response.Response	"Failed to create coupon"
//	@Router			/admin/promotions/create-coupon [post]
func (ch *CouponHandler) CreateCoupon(c *gin.Context) {
	var body request.Coupon
	if !ch.subHandler.BindRequest(c, &body) {
		return
	}

	err := ch.coupenUseCase.CreateCoupons(body)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to create coupon", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(201, "Success, created new coupon", nil, nil)
	c.JSON(http.StatusCreated, response)
}

// UpdateCoupon  Updates existing  coupon by id .
//
//	@Summary		Updates the existing
//	@Description	Create a new coupon with the provided details
//	@Tags			promotions
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			couponID	path		int					true	"coupon ID"
//	@Param			body		body		request.Coupon		true	"Coupon details"
//	@Success		200			{object}	response.Response	"Successful,coupon updated"
//	@Failure		400			{object}	response.Response	"Failed to bind JSON inputs from request"
//	@Failure		400			{object}	response.Response	"Failed, input does not meet validation criteria"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to update coupon"
//	@Router			/admin/promotions/update-coupon/{couponID}  [put]
func (ch *CouponHandler) UpdateCoupon(c *gin.Context) {
	var body request.Coupon
	if !ch.subHandler.BindRequest(c, &body) {
		return
	}

	couponID, ok := ch.subHandler.ParamInt(c, "couponID")
	if !ok {
		return
	}

	err := ch.coupenUseCase.UpdateCoupon(body, couponID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to update coupon", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Successful,coupon updated", nil, nil)
	c.JSON(statusOK, response)
}

// BlockCoupon  godoc
//
//	@Summary		Block coupon
//	@Description	Block the existing coupon by id.
//	@Tags			promotions
//	@Security		Bearer
//	@Produce		json
//	@Param			couponID	path		int					true	"coupon ID"
//	@Success		200			{object}	response.Response	"Success, coupon blocked"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to block coupon"
//	@Router			/admin/promotions/block-coupon/{couponID}  [put]
func (ch *CouponHandler) BlockCoupon(c *gin.Context) {

	couponID, ok := ch.subHandler.ParamInt(c, "couponID")
	if !ok {
		return
	}

	err := ch.coupenUseCase.BlockCoupon(couponID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to block coupon", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success, coupon blocked", nil, nil)
	c.JSON(statusOK, response)
}

// BlockCoupon  godoc
//
//	@Summary		Unblock coupon
//	@Description	Unblock the existing coupon by id.
//	@Tags			promotions
//	@Security		Bearer
//	@Produce		json
//	@Param			couponID	path		int	true	"coupon ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response
//	@Router			/admin/promotions/unblock-coupon/{couponID}  [put]
func (ch *CouponHandler) UnBlockCoupon(c *gin.Context) {
	couponID, ok := ch.subHandler.ParamInt(c, "couponID")
	if !ok {
		return
	}

	err := ch.coupenUseCase.UnBlockCoupon(couponID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success, coupon unblocked", nil, nil)
	c.JSON(statusOK, response)
}

// ListOutAllCouponsToAdmin  godoc
//
//	@Summary		List out all coupons to admin
//	@Description	List out all the created coupons to the admin.
//	@Security		Bearer
//	@Tags			promotions
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]response.Coupon}
//	@Failure		500	{object}	response.Response
//	@Router			/admin/promotions/all-coupons  [get]
func (ch *CouponHandler) ListOutAllCouponsToAdmin(c *gin.Context) {
	Coupons, err := ch.coupenUseCase.ViewAllCoupons()
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to fetch coupons", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success", Coupons, nil)
	c.JSON(statusOK, response)
}

// ApplyCoupon godoc
//
//	@Summary		Apply coupon
//	@Description	Apply the coupon and if valid provide coupon discount
//	@Tags			coupon
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			body	body		string				true	"Coupon code"
//	@Success		200		{object}	response.Response	"Success"
//	@Failure		400		{object}	response.Response
//	@Failure		403		{object}	response.Response
//	@Router			/coupon/apply [post]
func (ch *CouponHandler) ApplyCoupon(c *gin.Context) {
	var body string
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetIDFromContext(c)

	err := ch.coupenUseCase.ProcessApplyCoupon(body, userID)
	if err != nil {
		response := response.ResponseMessage(403, "Failed", nil, err.Error())
		c.JSON(http.StatusForbidden, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// ListOutAvailableCouponsToUser godoc
//
//	@Summary		List available coupons for the user
//	@Description	Get a list of available coupons for the authenticated user.
//	@Tags			coupon
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]response.Coupon}
//	@Failure		500	{object}	response.Response
//	@Router			/coupon/available [get]
func (ch *CouponHandler) ListOutAvailableCouponsToUser(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

	AvailabeCoupons, err := ch.coupenUseCase.ListOutAvailableCouponsToUser(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(AvailabeCoupons) == 0 {
		response := response.ResponseMessage(404, "No available coupons", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Success", AvailabeCoupons, nil)
	c.JSON(http.StatusOK, response)
}

// RemoveAppliedCoupon godoc
//
//	@Summary		Remove applied coupon
//	@Description	Remove the applied coupon from the user's coupon tracking.
//	@Tags			coupon
//	@Security		Bearer
//	@Produce		json
//	@Param			couponID	path		int					true	"Coupon ID"
//	@Success		200			{object}	response.Response	"Success"
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/coupon/remove/{couponID} [delete]
func (ch *CouponHandler) RemoveAppliedCoupon(c *gin.Context) {
	couponID, err := strconv.Atoi(c.Param("couponID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetIDFromContext(c)
	err = ch.coupenUseCase.RemoveFromCouponTracking(couponID, userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}
