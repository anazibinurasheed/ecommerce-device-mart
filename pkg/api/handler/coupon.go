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
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Coupon	true	"Coupon details"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/admin/promotions/create-coupon [post]
func (ch *CouponHandler) CreateCoupon(c *gin.Context) {
	var body request.Coupon
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := helper.ValidateStruct(body)
	if err != nil {
		response := response.ResponseMessage(400, "Unable to process without filling up required credentials", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	err = ch.coupenUseCase.CreateCoupons(body)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
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
//	@Accept			json
//	@Produce		json
//	@Param			couponID	path		int				true	"coupon ID"
//	@Param			body		body		request.Coupon	true	"Coupon details"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/admin/promotions/update-coupon/{couponID}  [put]
func (ch *CouponHandler) UpdateCoupon(c *gin.Context) {
	var body request.Coupon
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid Input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := helper.ValidateStruct(body)
	if err != nil {
		response := response.ResponseMessage(400, "Unable to process without filling up required credentials", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	couponID, err := strconv.Atoi(c.Param("couponID"))
	if err != nil {
		response := response.ResponseMessage(500, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ch.coupenUseCase.UpdateCoupon(body, couponID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Successful,coupon updated", nil, nil)
	c.JSON(http.StatusOK, response)
}

// BlockCoupon  godoc
//
//	@Summary		Block coupon
//	@Description	Block the existing coupon by id.
//	@Tags			promotions
//	@Produce		json
//	@Param			couponID	path		int	true	"coupon ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/admin/promotions/block-coupon/{couponID}  [put]
func (ch *CouponHandler) BlockCoupon(c *gin.Context) {
	couponID, err := strconv.Atoi(c.Param("couponID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ch.coupenUseCase.BlockCoupon(couponID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, coupon blocked", nil, nil)
	c.JSON(http.StatusOK, response)
}

// BlockCoupon  godoc
//
//	@Summary		Unblock coupon
//	@Description	Unblock the existing coupon by id.
//	@Tags			promotions
//	@Produce		json
//	@Param			couponID	path		int	true	"coupon ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/admin/promotions/unblock-coupon/{couponID}  [put]
func (ch *CouponHandler) UnBlockCoupon(c *gin.Context) {
	couponID, err := strconv.Atoi(c.Param("couponID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ch.coupenUseCase.UnBlockCoupon(couponID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, coupon unblocked", nil, nil)
	c.JSON(http.StatusOK, response)
}

// ListOutAllCouponsToAdmin  godoc
//
//	@Summary		List out all coupons to admin
//	@Description	List out all the created coupons to the admin.
//	@Tags			promotions
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/admin/promotions/all-coupons  [get]
func (ch *CouponHandler) ListOutAllCouponsToAdmin(c *gin.Context) {
	Coupons, err := ch.coupenUseCase.ViewAllCoupons()
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", Coupons, nil)
	c.JSON(http.StatusOK, response)
}

// ApplyCoupon godoc
//
//	@Summary		Apply coupon
//	@Description	Apply the coupon and if valid provide coupon discount
//	@Tags			coupon
//	@Accept			json
//	@Produce		json
//	@Param			body	body		string	true	"Coupon code"
//	@Success		200		{object}	response.Response
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

	userID, _ := helper.GetUserIDFromContext(c)

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
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/coupon/available [get]
func (ch *CouponHandler) ListOutAvailableCouponsToUser(c *gin.Context) {
	userID, _ := helper.GetUserIDFromContext(c)

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
//	@Produce		json
//	@Param			couponID	path		int	true	"Coupon ID"
//	@Success		200			{object}	response.Response
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

	userID, _ := helper.GetUserIDFromContext(c)
	err = ch.coupenUseCase.RemoveFromCouponTracking(couponID, userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}
