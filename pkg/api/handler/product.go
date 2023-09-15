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
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := ph.productUseCase.CreateNewCategory(body)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to create new category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, created new category", nil, nil)
	c.JSON(http.StatusOK, response)
}

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

	ListOfAllCategories, err := ph.productUseCase.ReadAllCategories(page, count)
	if err != nil {
		response := response.ResponseMessage(503, "Failed to retrieve categories", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if len(ListOfAllCategories) == 0 {
		response := response.ResponseMessage(404, "No data available", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "success", ListOfAllCategories, nil)
	c.JSON(http.StatusOK, response)
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

	err = ph.productUseCase.UpdateCategoryWithID(categoryID, body)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, category updated", nil, nil)
	c.JSON(http.StatusOK, response)
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
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ph.productUseCase.BlockCategoryWithID(categoryID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to block category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, blocked category", nil, err.Error())
	c.JSON(http.StatusOK, response)
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
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ph.productUseCase.BlockCategoryWithID(categoryID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed to block category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, unblocked category", nil, nil)
	c.JSON(http.StatusOK, response)
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
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	body.CategoryID = categoryID

	err = ph.productUseCase.CreateNewProduct(body)
	if err != nil {
		response := response.ResponseMessage(403, "Failed", nil, err.Error())
		c.JSON(http.StatusForbidden, response)
		return
	}
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, added new product", nil, nil)
	c.JSON(http.StatusOK, response)
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

	products, err := ph.productUseCase.DisplayAllProductsToAdmin(page, count)
	if err != nil {
		response := response.ResponseMessage(503, "Failed", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return

	}

	if len(products) == 0 {
		response := response.ResponseMessage(404, "No data available", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Success", products, nil)
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
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ph.productUseCase.UpdateProductWithID(productID, body)
	if err != nil {
		response := response.ResponseMessage(503, "Failed", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response := response.ResponseMessage(200, "Success, product updated", nil, nil)
	c.JSON(http.StatusOK, response)
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
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ph.productUseCase.BlockProductWithID(productID)
	if err != nil {
		response := response.ResponseMessage(503, "Failed", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
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
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ph.productUseCase.UnBlockProductWithID(productID)
	if err != nil {
		response := response.ResponseMessage(503, "Failed to block product", nil, err)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response := response.ResponseMessage(200, "Success, unblocked product", nil, nil)
	c.JSON(http.StatusOK, response)
}

// DisplayAllProductsToUser godoc
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
//	@Router			/products/ [get]
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

	products, err := ph.productUseCase.DisplayAllAvailabeProductsToUser(page, count)
	if err != nil {
		response := response.ResponseMessage(503, "Failed to retrieve products", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if len(products) == 0 {
		response := response.ResponseMessage(404, "No products available.", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Success", products, nil)
	c.JSON(http.StatusOK, response)
}

// ViewIndividualProduct godoc
//
//	@Summary		View a product
//	@Description	Retrieves details of a product with the specified ID.
//	@Tags			products
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		503			{object}	response.Response
//	@Router			/product/{productID} [get]
func (pd *ProductHandler) ViewIndividualProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pd.productUseCase.ViewProductByID(productID)
	if err != nil {
		response := response.ResponseMessage(503, "Failed", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response := response.ResponseMessage(200, "Success", product, nil)
	c.JSON(http.StatusOK, response)
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

	userID, _ := helper.GetUserIDFromContext(c)
	err = pd.productUseCase.ValidateProductRatingRequest(userID, productID)
	if err != nil {
		response := response.ResponseMessage(401, "Failed, user is unauthorized to perform a rating", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	response := response.ResponseMessage(200, "Success, authorized user", nil, nil)
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

	userID, _ := helper.GetUserIDFromContext(c)

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
//	@Router			/product/search [post]
func (ph *ProductHandler) SearchProducts(c *gin.Context) {
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

	search := c.Query("search")

	Products, err := ph.productUseCase.SearchProducts(search, page, count)
	if err != nil {
		response := response.ResponseMessage(403, "Failed", nil, err.Error())
		c.JSON(http.StatusForbidden, response)
		return
	}

	if len(Products) == 0 {
		response := response.ResponseMessage(404, "No data available", nil, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := response.ResponseMessage(200, "Success", Products, nil)
	c.JSON(http.StatusOK, response)
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
//	@Router			/category/{categoryID} [get]
func (ph *ProductHandler) ListProductsByCategory(c *gin.Context) {
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

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
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
