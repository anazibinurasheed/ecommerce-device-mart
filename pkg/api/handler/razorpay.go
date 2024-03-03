package handler

import (
	"net/http"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type RazorpayHandler struct {
	razorpayUseCase services.RazorpayUseCase
	orderUseCase    services.OrderUseCase
}

func NewRazorpayHandler(razorpayUseCase services.RazorpayUseCase, orderUseCase services.OrderUseCase) *RazorpayHandler {
	return &RazorpayHandler{
		razorpayUseCase: razorpayUseCase,
		orderUseCase:    orderUseCase,
	}
}

// GetOnlinePayment godoc
//
//	@Summary		Make payment razorpay
//	@Description	Make payment using razorpay page .
//	@Tags			checkout
//	@Security		Bearer
//	@Produce		json
//	@Success		200
//	@Failure		500	{object}	response.Response
//	@Router			/payment/online [get]
func (oh *RazorpayHandler) GetOnlinePayment(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)
	PaymentDetails, err := oh.razorpayUseCase.GetRazorPayDetails(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// c.HTML(200, "razorpay.html", gin.H{
	// 	"username":          PaymentDetails.Username,
	// 	"razorpay_order_id": PaymentDetails.RazorPayOrderID,
	// 	"amount":            PaymentDetails.Amount * 100,
	// })

	paymentDetails := response.PaymentDetails{
		Username:        PaymentDetails.Username,
		RazorPayOrderID: PaymentDetails.RazorPayOrderID,
		Amount:          PaymentDetails.Amount * 100,
	}

	response := response.ResponseMessage(statusOK, "success", paymentDetails, nil)
	c.JSON(statusOK, response)
}

// ProcessOnlinePayment is the handler function for verify  razorpay payment.
//
//	@Summary		Verify razorpay payment
//	@Description	Verify razorpay payment using razorpay credentials .
//	@Tags			checkout
//	@Security		Bearer
//	@Produce		json
//	@Param			body	body		request.VerifyPayment	true	"Payment details"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		403		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/payment/online/process [post]
func (oh *RazorpayHandler) ProcessOnlinePayment(c *gin.Context) {
	var body request.VerifyPayment
	if err := c.BindJSON(&body); err != nil {
		response := response.ResponseMessage(403, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetIDFromContext(c)

	err := oh.razorpayUseCase.VerifyRazorPayPayment(body.Signature, body.RazorpayOrderID, body.RazorPayPaymentID)
	if err != nil {
		response := response.ResponseMessage(403, "Failed", nil, err.Error())
		c.JSON(http.StatusForbidden, response)
		return

	}

	err = oh.orderUseCase.ConfirmedOrder(userId, 2) //2 is referring payment method razorpay(online)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, order placed", nil, nil)
	c.JSON(http.StatusOK, response)
}
