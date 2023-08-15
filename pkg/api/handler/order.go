package handler

import (
	"fmt"
	"net/http"
	"strconv"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}

// CheckOutPage is the handler function for displaying the checkout details.
//
//	@Summary		Checkout page
//	@Description	Displays the checkout details for the current user.
//	@Tags			checkout
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/checkout [get]
func (oh *OrderHandler) CheckOutPage(c *gin.Context) {
	userID, _ := helper.GetUserIDFromContext(c)

	CheckOutDetails, err := oh.orderUseCase.CheckOutDetails(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Succussful.", CheckOutDetails, nil)
	c.JSON(http.StatusOK, response)

}

// ConfirmCodDelivery is the handler function for confirming cash on delivery (COD) delivery.
//
//	@Summary		Confirm COD delivery
//	@Description	Confirms the cash on delivery (COD) delivery for the current user.
//	@Tags			checkout
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/payment/order-cod-confirmed [post]
func (oh *OrderHandler) ConfirmCodDelivery(c *gin.Context) {

	UserID, _ := helper.GetUserIDFromContext(c)
	err := oh.orderUseCase.ConfirmedOrder(UserID, 1) //1 is for  payment cash on delivery

	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success,order placed.", nil, nil)
	c.JSON(http.StatusOK, response)

}

// MakePaymentRazorpay godoc
//
//	@Summary		Make payment razorpay
//	@Description	Make payment using razorpay page .
//	@Tags			checkout
//	@Produce		json
//	@Success		200
//	@Failure		500	{object}	response.Response
//	@Router			/payment/razorpay [get]
func (oh *OrderHandler) MakePaymentRazorpay(c *gin.Context) {
	userID, _ := helper.GetUserIDFromContext(c)
	PaymentDetails, err := oh.orderUseCase.GetRazorPayDetails(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.HTML(200, "razorpay.html", gin.H{
		"username":          PaymentDetails.Username,
		"razorpay_order_id": PaymentDetails.RazorPayOrderId,
		"amount":            PaymentDetails.Amount * 100,
	})
}

// ProccessRazorpayOrder is the handler function for verify  razorpay payment.
//
//	@Summary		Verify razorpay payment
//	@Description	Verify razorpay payment using razorpay credentials .
//	@Tags			checkout
//	@Produce		json
//	@Param			body	body		request.VerifyPayment	true	"Payment details"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		403		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/payment/razorpay/process-order [post]
func (oh *OrderHandler) ProccessRazorpayOrder(c *gin.Context) {
	var body request.VerifyPayment
	if err := c.BindJSON(&body); err != nil {
		response := response.ResponseMessage(403, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetUserIDFromContext(c)

	err := oh.orderUseCase.VerifyRazorPayPayment(body.Signature, body.RazorpayOrderId, body.RazorPayPaymentId)
	if err != nil {
		response := response.ResponseMessage(403, "Failed", nil, err.Error())
		c.JSON(http.StatusForbidden, response)
		return

	}
	err = oh.orderUseCase.ConfirmedOrder(userId, 2) //2 is reffering payment method razorpay(online)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, order placed", nil, nil)
	c.JSON(http.StatusOK, response)
}

// UserOrderHistory is the handler function for retrieving the order history of the current user.
//
//	@Summary		Get order history
//	@Description	Retrieves the order history of the current user.
//	@Tags			user orders
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/my-orders [get]
func (oh *OrderHandler) UserOrderHistory(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetUserIDFromContext(c)

	orderHistory, err := oh.orderUseCase.GetUserOrderHistory(userId, page, count)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(orderHistory) == 0 {
		response := response.ResponseMessage(404, "No data available", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Success", orderHistory, nil)
	c.JSON(http.StatusOK, response)
}

// GetOrderManagementPage godoc
//
//	@Summary		Get order management data
//	@Description	Retrieves order management data.
//	@Tags			admin order management
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/admin/orders/management [get]
func (oh *OrderHandler) GetOrderManagementPage(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	OrderManagementPageDatas, err := oh.orderUseCase.GetOrderManagement(page, count)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", OrderManagementPageDatas, nil)
	c.JSON(http.StatusOK, response)
}

// GetAllOrderOverViewPage godoc
//
//	@Summary		Get all order overview data
//	@Description	Retrieves all order overview data.
//	@Tags			admin order management
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/admin/orders [get]
func (oh *OrderHandler) GetAllOrderOverViewPage(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	AllOrders, err := oh.orderUseCase.AllOrderOverView(page, count)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(AllOrders) == 0 {
		response := response.ResponseMessage(404, "No data available", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Success", AllOrders, nil)
	c.JSON(http.StatusOK, response)
}

// UpdateOrderStatus is the handler function for updating the status of an order.
//
//	@Summary		Update order status
//	@Description	Updates the status of an order with the specified ID.
//	@Tags			admin order management
//	@Produce		json
//	@Param			orderID		path		int	true	"Order ID"
//	@Param			statusID	path		int	true	"Status ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/admin/orders/{orderID}/update-status/{statusID} [put]
func (oh *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(orderID)
	statusID, err := strconv.Atoi(c.Param("statusID"))
	fmt.Println(orderID, statusID)

	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(orderID, statusID)
	err = oh.orderUseCase.UpdateOrderStatus(statusID, orderID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, status updated", nil, nil)
	c.JSON(http.StatusOK, response)

}

// CancelOrder godoc
//
//	@Summary		Cancel an order
//	@Description	Cancel the order. For online payments, the amount will be added to the user's wallet. For cash on delivery orders,will be marked as cancelled.
//	@Description	If the user has used a coupon for the order, the discount amount will be recalculated based on the percentage used and deducted from the refunding amount.
//	@Tags			user orders
//	@Accept			json
//	@Produce		json
//	@Param			orderID	path		int	true	"Order ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/my-orders/cancel/{orderID} [post]
func (oh *OrderHandler) CancelOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = oh.orderUseCase.OrderCancellation(orderID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, order cancelled", nil, nil)
	c.JSON(http.StatusOK, response)
}

// ReturnOrder godoc
//
//	@Summary		Return order
//	@Description	Return the order if the order is valid for return.Amount will be added to the user's wallet.
//	@Description	If the user has used a coupon for the order, the discount amount will be recalculated based on the percentage used and deducted from the refunding amount.
//	@Tags			user orders
//	@Accept			json
//	@Produce		json
//	@Param			orderID	path		int	true	"Order ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/my-orders/return/{orderID} [post]
func (oh *OrderHandler) ReturnOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = oh.orderUseCase.ProcessReturnRequest(orderID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := response.ResponseMessage(200, "Success, order return approved", nil, nil)
	c.JSON(http.StatusOK, response)
}

// DownloadInvoice godoc
//
//	@Summary		Download invoice
//	@Description	Download the invoice as a PDF file.
//	@Tags			user orders
//	@Produce		application/pdf
//	@Param			orderID	path		int	true	"Order ID"
//	@Success		200		{file}		application/pdf
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/my-orders/download-invoice/{orderID} [get]
func (oh *OrderHandler) DownloadInvoice(c *gin.Context) {
	// Sample invoice data (you can replace this with your actual invoice data)
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pdfData, err := oh.orderUseCase.CreateInvoice(orderID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Set the response headers for downloading the file
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=invoice.pdf")

	// Send the PDF data as the response
	c.Data(http.StatusOK, "application/pdf", pdfData)
}

// CreateUserWallet godoc
//
//	@Summary		Create user wallet
//	@Description	Initialize the wallet for the authenticated user.
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/wallet/create [post]
func (oh *OrderHandler) CreateUserWallet(c *gin.Context) {
	userID, _ := helper.GetUserIDFromContext(c)

	err := oh.orderUseCase.CreateUserWallet(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success ,wallet initialized ", nil, nil)
	c.JSON(http.StatusOK, response)

}

// ViewUserWallet godoc
//
//	@Summary		View user wallet
//	@Description	Get the wallet details for the authenticated user.
//	@Tags			wallet
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/wallet/ [get]
func (oh *OrderHandler) ViewUserWallet(c *gin.Context) {
	userID, _ := helper.GetUserIDFromContext(c)

	Wallet, err := oh.orderUseCase.GetUserWallet(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", Wallet, nil)
	c.JSON(http.StatusOK, response)
}

func (od *OrderHandler) WebhookHandler(c *gin.Context) {
	// // Get the raw webhook request body
	// webhookBody, _ := c.GetRawData()

	// // Get the signature from the X-Razorpay-Signature header
	// webhookSignature := c.GetHeader("X-Razorpay-Signature")

	// Validate the webhook signature
	// if !helper.ValidateWebhookSignature(string(webhookBody), webhookSignature) {
	// 	fmt.Println("errr")
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid webhook signature"})
	// 	return
	// }

	// Parse the webhook event data
	var eventData map[string]interface{}
	if err := c.BindJSON(&eventData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid webhook payload"})
		return
	}

	// Check the event type to determine payment status
	eventType, _ := eventData["event"].(string)
	switch eventType {

	case "payment.authorized":
		fmt.Println(eventType)
		paymentID, _ := eventData["razorpay_payment_id"].(string)
		fmt.Println("Payment authorized for Payment ID:", paymentID)
		// Payment is authorized, handle success case here
		c.JSON(http.StatusOK, gin.H{"message": "Payment authorized"})
	case "payment.failed":
		fmt.Println(eventType)

		// Payment failed, handle failure case here
		c.JSON(http.StatusOK, gin.H{"message": "Payment failed"})
	default:
		// Handle other events if needed
		c.JSON(http.StatusOK, gin.H{"message": "Event received"})
	}
}

// WalletPayment godoc
//
//	@Summary		Wallet payment
//	@Description	User can purchase using wallet
//	@Tags			checkout
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/payment/wallet [post]
func (od *OrderHandler) WalletPayment(c *gin.Context) {
	userID, _ := helper.GetUserIDFromContext(c)
	err := od.orderUseCase.ValidateWalletPayment(userID)
	if err != nil {
		response := response.ResponseMessage(400, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = od.orderUseCase.ConfirmedOrder(userID, 3) // 3 refers wallet payment
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (od *OrderHandler) WalletPaymentHistory(c *gin.Context) {

}
