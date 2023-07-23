package handler

import (
	"net/http"
	"strconv"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUseCase services.CartUseCase
}

func NewCartHandler(useCase services.CartUseCase) *CartHandler {
	return &CartHandler{
		cartUseCase: useCase,
	}
}

// AddToCart is the handler function for adding a product to the cart.
//
//	@Summary		Add product to cart
//	@Description	Adds a product to the cart for the authenticated user.
//	@Tags			cart
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/cart/add/{productID} [post]
func (ch *CartHandler) AddToCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetUserIdFromContext(c)

	err = ch.cartUseCase.AddToCart(userID, productID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Successful.", nil, nil)
	c.JSON(http.StatusOK, response)
}

// ViewCart is the handler function for viewing the cart items.
//
//	@Summary		View cart
//	@Description	Retrieves the cart items for the authenticated user.
//	@Tags			cart
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/cart [get]
func (ch *CartHandler) ViewCart(c *gin.Context) {
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
	userID, _ := helper.GetUserIdFromContext(c)

	CartItems, err := ch.cartUseCase.ViewCart(userID) ///////////////
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Successful.", CartItems, nil)
	c.JSON(http.StatusOK, response)

}

// IncrementQuantity is the handler function for incrementing the quantity of a product in the cart.
//
//	@Summary		Increment product quantity in cart
//	@Description	Increments the quantity of a product in the cart for the authenticated user.
//	@Tags			cart
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/cart/{productID}/increment [patch]
func (ch *CartHandler) IncrementQuantity(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetUserIdFromContext(c)

	err = ch.cartUseCase.IncrementQuantity(userID, productID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Successful.", nil, nil)
	c.JSON(http.StatusOK, response)

}

// DecrementQuantity is the handler function for decrementing the quantity of a product in the cart.
//
//	@Summary		Decrement product quantity in cart
//	@Description	Decrements the quantity of a product in the cart for the authenticated user.
//	@Tags			cart
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/cart/{productID}/decrement [patch]
func (ch *CartHandler) DecrementQuantity(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetUserIdFromContext(c)

	err = ch.cartUseCase.DecrementQuantity(userID, productID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Successful.", nil, nil)
	c.JSON(http.StatusOK, response)

}

// RemoveFromCart is the handler function for removing a product from the cart.
//
//	@Summary		Remove product from cart
//	@Description	Removes a product from the cart for the authenticated user.
//	@Tags			cart
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/cart/remove/{productID} [delete]
func (ch *CartHandler) RemoveFromCart(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetUserIdFromContext(c)

	err = ch.cartUseCase.RemoveFromCart(userID, productID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed.", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Successful.", nil, nil)
	c.JSON(http.StatusOK, response)

}
