package handler

import (
	"net/http"

	"github.com/anazibinurasheed/project-device-mart/pkg/usecase"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletUseCase services.WalletUseCase
	orderUseCase  services.OrderUseCase
}

func NewWalletHandler(walletUseCase services.WalletUseCase,
	orderUseCase services.OrderUseCase) *WalletHandler {
	return &WalletHandler{
		walletUseCase: walletUseCase,
		orderUseCase:  orderUseCase,
	}
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
func (oh *WalletHandler) CreateUserWallet(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

	err := oh.walletUseCase.CreateUserWallet(userID)
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
//	@Router			/wallet [get]
func (oh *WalletHandler) ViewUserWallet(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

	Wallet, err := oh.walletUseCase.GetUserWallet(userID)

	if err == usecase.ErrNoWallet {
		response := response.ResponseMessage(204, "user does not have wallet", nil, err.Error())
		c.JSON(http.StatusNoContent, response)
		return
	}
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", Wallet, nil)
	c.JSON(http.StatusOK, response)
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
func (od *WalletHandler) PayUsingWallet(c *gin.Context) {
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
func (od *WalletHandler) WalletTransactionHistory(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)

	walletHistory, err := od.walletUseCase.GetWalletHistory(userID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to get wallet history", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "success", walletHistory, nil)
	c.JSON(http.StatusOK, response)
}
