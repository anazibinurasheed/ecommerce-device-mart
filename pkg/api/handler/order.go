package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anazibinurasheed/project-device-mart/pkg/usecase"
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
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response{data=response.Checkout}
//	@Failure		500	{object}	response.Response
//	@Router			/checkout [get]
func (oh *OrderHandler) CheckOutPage(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

	CheckOutDetails, err := oh.orderUseCase.CheckOutDetails(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success.", CheckOutDetails, nil)
	c.JSON(http.StatusOK, response)

}

// ConfirmCodDelivery is the handler function for confirming cash on delivery (COD) delivery.
//
//	@Summary		Confirm COD delivery
//	@Description	Confirms the cash on delivery (COD) delivery for the current user.
//	@Tags			checkout
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/payment/cod-confirm [post]
func (oh *OrderHandler) ConfirmCodDelivery(c *gin.Context) {

	UserID, _ := helper.GetIDFromContext(c)
	err := oh.orderUseCase.ConfirmedOrder(UserID, 1) //1 is for  payment cash on delivery

	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success,order placed.", nil, nil)
	c.JSON(http.StatusOK, response)

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
func (oh *OrderHandler) GetOnlinePayment(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)
	PaymentDetails, err := oh.orderUseCase.GetRazorPayDetails(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.HTML(200, "razorpay.html", gin.H{
		"username":          PaymentDetails.Username,
		"razorpay_order_id": PaymentDetails.RazorPayOrderID,
		"amount":            PaymentDetails.Amount * 100,
	})
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
func (oh *OrderHandler) ProcessOnlinePayment(c *gin.Context) {
	var body request.VerifyPayment
	if err := c.BindJSON(&body); err != nil {
		response := response.ResponseMessage(403, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, _ := helper.GetIDFromContext(c)

	err := oh.orderUseCase.VerifyRazorPayPayment(body.Signature, body.RazorpayOrderID, body.RazorPayPaymentID)
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

// UserOrderHistory is the handler function for retrieving the order history of the current user.
//
//	@Summary		Get order history
//	@Description	Retrieves the order history of the current user.
//	@Tags			user orders
//	@Security		Bearer
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]response.Orders}
//	@Failure		500	{object}	response.Response
//	@Router			/orders [get]
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

	userId, _ := helper.GetIDFromContext(c)

	orderHistory, err := oh.orderUseCase.GetUserOrderHistory(userId, page, count)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
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
//	@Security		Bearer
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response{data=response.OrderManagement}
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
//	@Security		Bearer
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]response.Orders}
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


	response := response.ResponseMessage(200, "Success", AllOrders, nil)
	c.JSON(http.StatusOK, response)
}

// UpdateOrderStatus is the handler function for updating the status of an order.
//
//	@Summary		Update order status
//	@Description	Updates the status of an order with the specified ID.
//	@Tags			admin order management
//	@Security		Bearer
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
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			orderID	path		int	true	"Order ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/orders/cancel/{orderID} [post]
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
//	@Description	If the user has used a coupon for the order, the discount amount will be recalculated based
//	@Description	on the percentage used and deducted from the refunding amount.
//	@Security		Bearer
//	@Tags			user orders
//	@Accept			json
//	@Produce		json
//	@Param			orderID	path		int	true	"Order ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/orders/return/{orderID} [post]
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

// CreateInvoice godoc
//
//	@Summary		Download invoice
//	@Description	Download the invoice as a PDF file.
//	@Tags			user orders
//	@Security		Bearer
//	@Produce		application/pdf
//	@Param			orderID	path		int	true	"Order ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/orders/invoice/{orderID} [get]
func (oh *OrderHandler) CreateInvoice(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	invoiceDetails, err := oh.orderUseCase.CreateInvoice(orderID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, invoiceDetails)

}

// CreateUserWallet godoc
//
//	@Summary		Create user wallet
//	@Description	Initialize the wallet for the authenticated user.
//	@Tags			wallet
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/wallet/create [post]
func (oh *OrderHandler) CreateUserWallet(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

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
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response{response.Wallet}
//	@Failure		500	{object}	response.Response
//	@Router			/wallet/ [get]
func (oh *OrderHandler) ViewUserWallet(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

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

// PayUsingWallet godoc
//
//	@Summary		Pay using wallet
//	@Description	User can purchase using wallet
//	@Tags			checkout
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/payment/wallet [post]
func (od *OrderHandler) PayUsingWallet(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)
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

// WalletTransactionHistory godoc
//
//	@Summary		User wallet transaction history
//	@Description	This endpoint will show all the wallet transaction history of the user.
//	@Tags			wallet
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/wallet/history [get]
func (od *OrderHandler) WalletTransactionHistory(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

	walletHistory, err := od.orderUseCase.GetWalletHistory(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to get wallet history", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "success", walletHistory, nil)
	c.JSON(http.StatusOK, response)
}

// MonthlySalesReport godoc
//
//	@Summary		Monthly sales report
//	@Description	Sales report of last 30 days from the requested time
//	@Tags			sales-report
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	response.Response{data=response.MonthlySalesReport}	"Success"
//	@Success		200	{object}	response.Response{data=response.MonthlySalesReport}	"No orders created yet"
//	@Failure		500	{object}	response.Response									"Failed to generate the sales report"
//	@Router			/admin/sales-report [get]
func (od *OrderHandler) MonthlySalesReport(c *gin.Context) {
	salesReport, err := od.orderUseCase.MonthlySalesReport()

	if err == usecase.ErrNoOrders {
		response := response.ResponseMessage(statusOK, "No orders created yet", salesReport, nil)
		c.JSON(statusOK, response)
		return
	}

	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to generate the sales report", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success", salesReport, nil)
	c.JSON(statusOK, response)
}
