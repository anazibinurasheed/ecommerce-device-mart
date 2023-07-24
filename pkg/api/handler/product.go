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

// cap
type ProductHandler struct {
	productUseCase services.ProductUseCase
}

func NewProductHandler(useCase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: useCase}
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	Creates a new category based on the provided category name.
//	@Tags			admin category management
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Category	true	"Category name"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/admin/category/add-category [post]
func (ph *ProductHandler) CreateCategory(c *gin.Context) {
	var body request.Category
	if err := c.BindJSON(&body); err != nil {

		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input ",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	NewCategory, err := ph.productUseCase.CreateNewCategory(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to create new category",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "success,created new category",
		Data:       NewCategory,
		Error:      nil,
	})

}

//////////////////////////////////////////////////////////////////////////////////////////////////

// ReadAllCategories godoc
//
//	@Summary		List out all categories
//	@Description	Retrieves all categories available.
//	@Tags			admin category management
//	@Param			page	query	int	true	"Page number"				default(1)
//	@Param			count	query	int	true	"Number of items per page"	default(10)
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		503	{object}	response.Response
//	@Router			/admin/category/all-category [get]
func (ph *ProductHandler) ReadAllCategories(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ListOfAllCategories, err := ph.productUseCase.ReadAllCategories(page, count)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, response.Response{
			StatusCode: 503,
			Message:    " Failed to retrieve all categories ",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	if len(ListOfAllCategories) == 0 {
		response := response.ResponseMessage(404, "No data available .", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       ListOfAllCategories,
		Error:      nil,
	})

}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	Updates a category with the specified ID.
//	@Tags			admin category management
//	@Accept			json
//	@Param			categoryID	path	int					true	"Category ID"
//	@Param			body		body	request.Category	true	"Category name"
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/admin/category/update-category/{categoryID} [patch]
func (ph *ProductHandler) UpdateCategory(c *gin.Context) {
	var body request.Category
	if err := c.BindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	UpdatedCategory, err := ph.productUseCase.UpdateCategoryWithId(categoryID, body)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success,category updated", UpdatedCategory, nil)
	c.JSON(http.StatusInternalServerError, response)

}

// BlockCategory godoc
//
//	@Summary		Block a category
//	@Description	Blocks a category with the specified ID.
//	@Tags			admin category management
//	@Param			categoryID	path	int	true	"Category ID"
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/admin/category/block-category/{categoryID} [patch]
func (ph *ProductHandler) BlockCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input",
			Data:       nil,
			Error:      nil,
		})
		return
	}
	BlockedCategory, err := ph.productUseCase.BlockCategoryWithId(categoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to block category",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Success,Blocked  category",
		Data:       BlockedCategory,
		Error:      nil,
	})

}

// UnBlockCategory godoc
//
//	@Summary		Unblock a category
//	@Description	Unblocks a category with the specified ID.
//	@Tags			admin category management
//	@Accept			json
//	@Param			categoryID	path	int	true	"Category ID"
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/admin/category/unblock-category/{categoryID} [patch]
func (ph *ProductHandler) UnBlockCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input",
			Data:       nil,
			Error:      nil,
		})
		return
	}
	BlockedCategory, err := ph.productUseCase.BlockCategoryWithId(categoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to block category",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Success,unblocked  category",
		Data:       BlockedCategory,
		Error:      nil,
	})

}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Creates a new product with the specified details.
//	@Tags			admin product management
//	@Accept			json
//	@Produce		json
//	@Param			categoryID	path		int				true	"Category ID"
//	@Param			body		body		request.Product	true	"Product details"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		403			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/admin/products/add-product/{categoryID} [post]
func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var body request.Product
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input",
			Data:       nil,
			Error:      nil,
		})
		return
	}

	body.CategoryID = categoryID
	fmt.Println("BRAND OF PRODUCT ", body.CategoryID)
	Product, err := ph.productUseCase.CreateNewProduct(body)
	if err != nil && Product.CategoryID != 0 {
		c.JSON(http.StatusForbidden, response.Response{
			StatusCode: 403,
			Message:    "Product already exists",
			Data:       Product,
			Error:      err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Something went wrong,Cant create product,please try again later",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Success,Product added",
		Data:       Product,
		Error:      nil,
	})

}

// / DisplayAllProductsToAdmin is the handler function for viewing all products by an admin.
//
//	@Summary		Display all products to admin
//	@Description	Retrieves a list of all products including blocked.
//	@Tags			admin product management
//	@Produce		json
//	@Param			page	query		int	true	"Page number"				default(1)
//	@Param			count	query		int	true	"Number of items per page"	default(10)
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		503		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Router			/admin/products/all-products [get]
func (ph *ProductHandler) DisplayAllProductsToAdmin(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ListOfAllProducts, err := ph.productUseCase.DisplayAllProductsToAdmin(page, count)
	if err != nil {
		response := response.ResponseMessage(503, "An error occurred during processing. Please try again.", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return

	}

	if len(ListOfAllProducts) == 0 {
		response := response.ResponseMessage(404, "No data available for the specified page number.", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Successful", ListOfAllProducts, nil)
	c.JSON(http.StatusOK, response)

}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Updates an existing product with the specified ID.
//	@Tags			admin product management
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int				true	"Product ID"
//	@Param			body		body		request.Product	true	"Product object"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		503			{object}	response.Response
//	@Router			/admin/products/update-product/{productID} [patch]
func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	var body request.Product

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input",
			Data:       nil,
			Error:      nil,
		})
		return
	}

	UpdatedProduct, err := ph.productUseCase.UpdateProductWithId(productID, body)

	if err != nil {

		c.JSON(http.StatusServiceUnavailable, response.Response{
			StatusCode: 503,
			Message:    " Our service is currently unavailable due to maintenance or high server load.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       UpdatedProduct,
		Error:      nil,
	})

}

// BlockProduct godoc
//
//	@Summary		Block a product
//	@Description	Blocks a product with the specified ID.
//	@Tags			admin product management
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		503			{object}	response.Response
//	@Router			/admin/products/block-product/{productID} [patch]
func (ph *ProductHandler) BlockProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input ",
			Data:       nil,
			Error:      nil,
		})
		return
	}

	BlockedProduct, err := ph.productUseCase.BlockProductWithId(productID)
	if err != nil {

		c.JSON(http.StatusServiceUnavailable, response.Response{
			StatusCode: 503,
			Message:    "Failed to block product",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       BlockedProduct,
		Error:      nil,
	})

}

// UnBlockProduct godoc
//
//	@Summary		Unblock a product
//	@Description	Unblocks a product with the specified ID.
//	@Tags			admin product management
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		503			{object}	response.Response
//	@Router			/admin/products/unblock-product/{productID} [patch]
func (ph *ProductHandler) UnBlockProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input ",
			Data:       nil,
			Error:      nil,
		})
		return
	}
	fmt.Println(productID)

	UnBlockedProduct, err := ph.productUseCase.UnBlockProductWithId(productID)
	if err != nil {

		c.JSON(http.StatusServiceUnavailable, response.Response{
			StatusCode: 503,
			Message:    " Failed to block product",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       UnBlockedProduct,
		Error:      nil,
	})

}

// DisplayAllProductsToUser is the handler function for displaying all available products to the user.
//
//	@Summary		Display all products to the user
//	@Description	Retrieves all available products for the user.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	true	"Page number"				default(1)
//	@Param			count	query		int	true	"Number of items per page"	default(10)
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Failure		503		{object}	response.Response
//	@Router			/products [get]
func (ph *ProductHandler) DisplayAllProductsToUser(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ListOfAllProducts, err := ph.productUseCase.DisplayAllAvailabeProductsToUser(page, count)
	if err != nil {
		response := response.ResponseMessage(503, "An error occurred during processing. Please try again.", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if len(ListOfAllProducts) == 0 {
		response := response.ResponseMessage(404, "No products available for the specified page number.", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Successful", ListOfAllProducts, nil)
	c.JSON(http.StatusOK, response)
}

// ViewProductItem is the handler function for viewing a product by ID.
//
//	@Summary		View a product
//	@Description	Retrieves details of a product with the specified ID.
//	@Tags			products
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		503			{object}	response.Response
//	@Router			/product-item/{productID} [get]
func (pd *ProductHandler) ViewProductItem(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid input ",
			Data:       nil,
			Error:      nil,
		})
		return
	}

	product, err := pd.productUseCase.ViewProductById(productID)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, response.Response{
			StatusCode: 503,
			Message:    "An error occurred during processing. Please try again.",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successful",
		Data:       product,
		Error:      nil,
	})
}

// ValidateRatingRequest is the handler function for validating a product rating request.
//
//	@Summary		Validate product rating request
//	@Description	Validates if the user is authorized to rate a product.
//	@Tags			user orders
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		401			{object}	response.Response
//	@Router			/product/rating/{productID} [get]
func (pd *ProductHandler) ValidateRatingRequest(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))

	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}
	userID, _ := helper.GetUserIdFromContext(c)
	err = pd.productUseCase.ValidateProductRatingRequest(userID, productID)
	if err != nil {
		response := response.ResponseMessage(401, "Failed", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	response := response.ResponseMessage(200, "Success,authorized user", nil, nil)
	c.JSON(http.StatusOK, response)

}

// AddProductRating is the handler function for adding a product rating.
//
//	@Summary		Add product rating
//	@Description	Adds a new rating for a product.
//	@Tags			user orders
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int				true	"Product ID"
//	@Param			rating		body		request.Rating	true	"Rating details"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		403			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/product/rating/{productID} [post]
func (pd *ProductHandler) AddProductRating(c *gin.Context) {
	var body request.Rating
	if err := c.BindJSON(&body); err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.Rating == 0 || body.Description == "" {
		response := response.ResponseMessage(403, "Fields not be empty", nil, nil)
		c.JSON(http.StatusForbidden, response)
		return
	}
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	userID, _ := helper.GetUserIdFromContext(c)
	err = pd.productUseCase.InsertNewProductRating(userID, productID, body)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to add rating", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := response.ResponseMessage(200, "Success ", nil, nil)
	c.JSON(http.StatusOK, response)
}

// SearchProducts searches for products based on the given input.
//
//	@Summary		Search Products
//	@Description	Searches for products based on the provided search input
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			search	query		string	true	"Search input"
//	@Param			page	query		int		true	"Page number"				default(1)
//	@Param			count	query		int		true	"Number of items per page"	default(10)
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		403		{object}	response.Response
//	@Router			/products/search [post]
func (ph *ProductHandler) SearchProducts(c *gin.Context) {
	// var body string
	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	c.JSON(http.StatusBadRequest, response.Response{
	// 		StatusCode: 400,
	// 		Message:    "Invalid input. ",
	// 		Data:       nil,
	// 		Error:      err.Error(),
	// 	})
	// 	return
	// }
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	search := c.Query("search")

	Products, err := ph.productUseCase.SearchProducts(search, page, count)
	if err != nil {
		c.JSON(http.StatusForbidden, response.Response{
			StatusCode: 403,
			Message:    "Failed. ",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	if len(Products) == 0 {
		response := response.ResponseMessage(404, "No data available.", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Success. ",
		Data:       Products,
		Error:      nil,
	})

}

// ListProductsByCategory lists products by category ID.
//
//	@Summary		List products by category
//	@Description	Lists products based on the provided category ID.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			categoryID	path		int	true	"Category ID"
//	@Param			page		query		int	true	"Page number"				default(1)
//	@Param			count		query		int	true	"Number of items per page"	default(10)
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/products-by-category/{categoryID} [get]
func (ph *ProductHandler) ListProductsByCategory(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry.", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}
	Products, err := ph.productUseCase.GetProductsByCategory(categoryID, page, count)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := response.ResponseMessage(200, "Success", Products, nil)
	c.JSON(http.StatusBadRequest, response)

}
