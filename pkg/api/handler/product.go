package handler

import (
	"net/http"
	"strconv"

	"github.com/anazibinurasheed/project-device-mart/pkg/usecase"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase services.ProductUseCase
	subHandler     helper.SubHandler
}

func NewProductHandler(useCase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: useCase}
}

// CreateCategory godoc
//
//	@Summary		Create category
//	@Description	Creates a new category based on the provided category name.
//	@Tags			admin category management
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.Category							true	"Category name"
//	@Success		200		{object}	response.Response{data=response.Category}	"Success, category created"
//	@Failure		400		{object}	response.Response							"Failed to bind JSON inputs from request"
//	@Failure		400		{object}	response.Response							"Failed, input does not meet validation criteria"
//	@Failure		500		{object}	response.Response							"Failed to create new category"
//	@Router			/admin/category/add-category [post]
func (p *ProductHandler) CreateCategory(c *gin.Context) {
	var body request.Category
	if !p.subHandler.BindRequest(c, &body) {
		return
	}

	category, err := p.productUseCase.CreateCategory(body)
	if err != nil {
		statusCode, msg := statusInternalServerError, "Failed to create category"

		if err == usecase.ErrRecordAlreadyExist {
			statusCode, msg = statusConflict, "Failed, category already exist"
		}

		response := response.ResponseMessage(statusCode, msg, nil, err.Error())
		c.JSON(statusCode, response)
		return
	}

	response := response.ResponseMessage(200, "Success, category created", category, nil)
	c.JSON(statusOK, response)
}

// ReadAllCategories godoc
//
//	@Summary		List out all categories
//	@Description	Retrieves all available categories.
//	@Tags			admin category management
//	@Security		Bearer
//	@Produce		json
//	@Param			page	query		int											true	"Page number"				default(1)
//	@Param			count	query		int											true	"Number of items per page"	default(10)
//	@Success		200		{object}	response.Response{data=[]response.Category}	"Success"
//	@Failure		400		{object}	response.Response							"Failed to bind page info from request"
//	@Failure		503		{object}	response.Response							"Failed to retrieve categories"
//	@Router			/admin/category/categories [get]
func (p *ProductHandler) ReadAllCategories(c *gin.Context) {
	page, count, ok := p.subHandler.GetPageNCount(c)
	if !ok {
		return
	}

	categories, err := p.productUseCase.ReadAllCategories(page, count)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to retrieve categories", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success", categories, nil)
	c.JSON(http.StatusOK, response)
}

// Its a duplication of ReadAllCategories, for swagger specification
//
// Categories godoc
//
//	@Summary		List out all categories
//	@Description	Retrieves all available categories.
//	@Tags			category
//	@Security		Bearer
//	@Produce		json
//	@Param			page	query		int											true	"Page number"				default(1)
//	@Param			count	query		int											true	"Number of items per page"	default(10)
//	@Success		200		{object}	response.Response{data=[]response.Category}	"Success"
//	@Failure		400		{object}	response.Response							"Failed to bind page info from request"
//	@Failure		503		{object}	response.Response							"Failed to retrieve categories"
//	@Router			/category/all [get]
func (p *ProductHandler) Categories(c *gin.Context) {
	page, count, ok := p.subHandler.GetPageNCount(c)
	if !ok {
		return
	}

	categories, err := p.productUseCase.ReadAllCategories(page, count)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to retrieve categories", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success", categories, nil)
	c.JSON(http.StatusOK, response)
}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	Updates a category with the specified ID.
//	@Tags			admin category management
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			categoryID	path		int					true	"Category ID"
//	@Param			body		body		request.Category	true	"Category name"
//	@Success		200			{object}	response.Response	"Success, category updated"
//	@Failure		400			{object}	response.Response	"Failed to bind JSON inputs from request"
//	@Failure		400			{object}	response.Response	"Failed, input does not meet validation criteria"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to update category"
//	@Router			/admin/category/update-category/{categoryID} [put]
func (p *ProductHandler) UpdateCategory(c *gin.Context) {
	var body request.Category
	if !p.subHandler.BindRequest(c, &body) {
		return
	}

	categoryID, ok := p.subHandler.ParamInt(c, "categoryID")
	if !ok {
		return
	}

	err := p.productUseCase.UpdateCategoryByID(categoryID, body)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to update category", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success, category updated", nil, nil)
	c.JSON(statusOK, response)
}

// BlockCategory godoc
//
//	@Summary		Block a category
//	@Description	Blocks a category with the specified ID.
//	@Tags			admin category management
//	@Security		Bearer
//	@Produce		json
//	@Param			categoryID	path		int					true	"Category ID"
//	@Success		200			{object}	response.Response	"Success, category has been blocked"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to block category"
//	@Router			/admin/category/block-category/{categoryID} [put]
func (ph *ProductHandler) BlockCategory(c *gin.Context) {
	categoryID, ok := ph.subHandler.ParamInt(c, "categoryID")
	if !ok {
		return
	}

	err := ph.productUseCase.BlockCategoryByID(categoryID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to block category", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, category has been blocked", nil, nil)
	c.JSON(statusOK, response)
}

// UnBlockCategory godoc
//
//	@Summary		Unblock a category
//	@Description	Unblocks a category with the specified ID.
//	@Tags			admin category management
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			categoryID	path		int					true	"Category ID"
//	@Success		200			{object}	response.Response	"Success, category unblocked"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to block category"
//	@Router			/admin/category/unblock-category/{categoryID} [put]
func (ph *ProductHandler) UnBlockCategory(c *gin.Context) {
	categoryID, ok := ph.subHandler.ParamInt(c, "categoryID")
	if !ok {
		return
	}

	err := ph.productUseCase.UnBlockCategoryByID(categoryID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to block category", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success, category unblocked", nil, nil)
	c.JSON(statusOK, response)
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Creates a new product with the specified details.
//	@Tags			admin product management
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			categoryID	path		int											true	"Category ID"
//	@Param			body		body		request.Product								true	"Product details"
//	@Success		200			{object}	response.Response{data=response.Product}	"Success, added new product"
//	@Failure		400			{object}	response.Response							"Failed to bind JSON inputs from request"
//	@Failure		400			{object}	response.Response							"Failed, input does not meet validation criteria"
//	@Failure		400			{object}	response.Response							"Failed to retrieve param from URL"
//	@Failure		400			{object}	response.Response							"Failed, category not found"
//	@Failure		409			{object}	response.Response							"Failed, product already exist with same name"
//	@Failure		500			{object}	response.Response							"Failed to create product"
//	@Router			/admin/product/add-product/{categoryID} [post]
func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var body request.Product
	if !ph.subHandler.BindRequest(c, &body) {
		return
	}

	categoryID, ok := ph.subHandler.ParamInt(c, "categoryID")
	if !ok {
		return
	}

	body.CategoryID = categoryID

	product, err := ph.productUseCase.CreateProduct(body)

	if err != nil {
		func() {
			status, msg := statusInternalServerError, "Failed to create product"
			switch {
			case err == usecase.ErrCategoryNotFound:
				status = statusBadRequest
				msg = "Failed, category not found"
			case err == usecase.ErrRecordAlreadyExist:
				status = statusConflict
				msg = "Failed, product already exist with same name"

			}
			response := response.ResponseMessage(status, msg, nil, err.Error())
			c.JSON(status, response)

		}()
		return
	}

	response := response.ResponseMessage(statusOK, "Success, added new product", product, nil)
	c.JSON(statusOK, response)
}

// / ShowProductsToAdmin is the handler function for viewing all products by an admin.
//
//	@Summary		Display  products to admin
//	@Description	Retrieves a list of all products including blocked.
//	@Tags			admin product management
//	@Security		Bearer
//	@Produce		json
//	@Param			page	query		int											true	"Page number"				default(1)
//	@Param			count	query		int											true	"Number of items per page"	default(10)
//	@Success		200		{object}	response.Response{data=[]response.Product}	"Success"
//	@Failure		400		{object}	response.Response							"Failed to bind page info from request"
//	@Failure		500		{object}	response.Response							"Failed to fetch products"
//	@Router			/admin/product/products [get]
func (ph *ProductHandler) ShowProductsToAdmin(c *gin.Context) {
	page, count, ok := ph.subHandler.GetPageNCount(c)
	if !ok {
		return
	}

	products, err := ph.productUseCase.DisplayAllProductsToAdmin(page, count)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to fetch products", nil, err.Error())
		c.JSON(http.StatusServiceUnavailable, response)
		return

	}

	response := response.ResponseMessage(statusOK, "Success", products, nil)
	c.JSON(http.StatusOK, response)
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Updates an existing product with the specified ID.
//	@Tags			admin product management
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int						true	"Product ID"
//	@Param			body		body		request.UpdateProduct	true	"Product details"
//	@Success		200			{object}	response.Response		"Success, product updated"
//	@Failure		400			{object}	response.Response		"Failed to bind JSON inputs from request"
//	@Failure		400			{object}	response.Response		"Failed, input does not meet validation criteria"
//	@Failure		400			{object}	response.Response		"Failed, input does not meet validation criteria"
//	@Failure		500			{object}	response.Response		"Failed update product"
//	@Router			/admin/product/update-product/{productID} [put]
func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	var body request.UpdateProduct
	if !ph.subHandler.BindRequest(c, &body) {
		return
	}

	productID, ok := ph.subHandler.ParamInt(c, "productID")
	if !ok {
		return
	}

	err := ph.productUseCase.UpdateProductByID(productID, body)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed update product", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success, product updated", nil, nil)
	c.JSON(statusOK, response)
}

// BlockProduct godoc
//
//	@Summary		Block a product
//	@Description	Blocks a product with the specified ID.
//	@Tags			admin product management
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int					true	"Product ID"
//	@Success		200			{object}	response.Response	"Success, product blocked"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to block product"
//	@Router			/admin/product/block-product/{productID} [put]
func (ph *ProductHandler) BlockProduct(c *gin.Context) {
	productID, ok := ph.subHandler.ParamInt(c, "productID")
	if !ok {
		return
	}

	err := ph.productUseCase.BlockProductByID(productID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to block product", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success, product blocked", nil, nil)
	c.JSON(http.StatusOK, response)
}

// UnBlockProduct godoc
//
//	@Summary		Unblock a product
//	@Description	Unblocks a product with the specified ID.
//	@Tags			admin product management
//	@Security		Bearer
//	@Produce		json
//	@Param			productID	path		int					true	"Product ID"
//	@Success		200			{object}	response.Response	"Success, unblocked product"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to unblock product"
//	@Router			/admin/product/unblock-product/{productID} [put]
func (ph *ProductHandler) UnBlockProduct(c *gin.Context) {
	productID, ok := ph.subHandler.ParamInt(c, "productID")
	if !ok {
		return
	}

	err := ph.productUseCase.UnBlockProductByID(productID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to unblock product", nil, err)
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success, unblocked product", nil, nil)
	c.JSON(statusOK, response)
}

// DisplayAllProductsToUser godoc
//
//	@Summary		Display all products to the user
//	@Description	Retrieves all available products for the user.
//	@Tags			products
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int											true	"Page number"				default(1)
//	@Param			count	query		int											true	"Number of items per page"	default(10)
//	@Success		200		{object}	response.Response{data=response.Product}	"Success"
//	@Failure		400		{object}	response.Response							"Failed to bind page info from request"
//	@Failure		500		{object}	response.Response							"Failed to retrieve products"
//	@Router			/product/all [get]
func (ph *ProductHandler) DisplayAllProductsToUser(c *gin.Context) {
	page, count, ok := ph.subHandler.GetPageNCount(c)
	if !ok {
		return
	}

	userID, _ := helper.GetIDFromContext(c)

	products, err := ph.productUseCase.DisplayAllProductsToUser(userID, page, count)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to retrieve products", nil, err.Error())
		c.JSON(statusInternalServerError, response)
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
//	@Security		Bearer
//	@Produce		json
//	@Param			productID	path		int												true	"Product ID"
//	@Success		200			{object}	response.Response{data=response.ProductItem}	"Success"
//	@Failure		400			{object}	response.Response								"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response								"Failed to fetch product"
//	@Router			/product/{productID} [get]
func (pd *ProductHandler) ViewIndividualProduct(c *gin.Context) {
	productID, ok := pd.subHandler.ParamInt(c, "productID")
	if !ok {
		return
	}

	userID, _ := helper.GetIDFromContext(c)
	product, err := pd.productUseCase.ViewIndividualProduct(userID, productID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "Failed to fetch product", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success", product, nil)
	c.JSON(statusOK, response)
}

// ValidateRatingRequest is the handler function for validating a product rating request.
//
//	@Summary		Validate product rating request
//	@Description	Validates if the user is authorized to rate a product.
//	@Tags			user orders
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int					true	"Product ID"
//	@Success		200			{object}	response.Response	"Success, user is valid for rating the product"
//	@Failure		400			{object}	response.Response	"Failed to retrieve param from URL"
//	@Failure		500			{object}	response.Response	"Failed to retrieve user id from context"
//	@Failure		409			{object}	response.Response	"User already rated this product"
//	@Failure		401			{object}	response.Response	"Order is not delivered or returned"
//	@Failure		401			{object}	response.Response	"User doesn't purchased this product"
//	@Router			/product/rating/{productID} [get]
func (pd *ProductHandler) ValidateRatingRequest(c *gin.Context) {
	productID, ok := pd.subHandler.ParamInt(c, "productID")
	if !ok {
		return
	}

	userID, ok := pd.subHandler.GetUserID(c)
	if !ok {
		return
	}

	err := pd.productUseCase.ValidateProductRatingRequest(userID, productID)
	if err != nil {

		status, msg := statusInternalServerError, "Failed to validate the rating request"

		switch {
		case err == usecase.ErrRecordAlreadyExist:
			status, msg = statusConflict, "User already rated this product"
		case err == usecase.ErrInProcessing:
			status, msg = statusUnauthorized, "Order is not delivered or returned"
		case err == usecase.ErrNoRecord:
			status, msg = statusUnauthorized, "User doesn't purchased this product"
		}

		response := response.ResponseMessage(status, msg, nil, err.Error())
		c.JSON(status, response)
		return
	}

	response := response.ResponseMessage(statusOK, "Success, user is valid for rating the product", nil, nil)
	c.JSON(statusOK, response)
}

// AddProductRating is the handler function for adding a product rating.
//
//	@Summary		Add product rating
//	@Description	Adds a new rating for a product.
//	@Tags			user orders
//	@Security		Bearer
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

	userID, _ := helper.GetIDFromContext(c)

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
//	@Security		Bearer
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

	response := response.ResponseMessage(200, "Success", Products, nil)
	c.JSON(http.StatusOK, response)
}

// ListProductsByCategory lists products by category ID.
//
//	@Summary		List products by category
//	@Description	Lists products based on the provided category ID.
//	@Tags			products
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			categoryID	path		int	true	"Category ID"
//	@Param			page		query		int	true	"Page number"				default(1)
//	@Param			count		query		int	true	"Number of items per page"	default(10)
//	@Success		200			{object}	response.Response{data=[]response.Product}
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/product/category/{categoryID} [get]
func (ph *ProductHandler) ListProductsByCategoryUser(c *gin.Context) {
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
	c.JSON(statusOK, response)
}

// ListProductsByCategoryAdmin lists products by category ID.
//
//	@Summary		List products by category
//	@Description	Lists products based on the provided category ID.
//	@Tags			admin product management
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			categoryID	path		int	true	"Category ID"
//	@Param			page		query		int	true	"Page number"				default(1)
//	@Param			count		query		int	true	"Number of items per page"	default(10)
//	@Success		200			{object}	response.Response{data=[]response.Product}
//	@Failure		400			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/admin/product/category/{categoryID} [get]
func (ph *ProductHandler) ListProductsByCategoryAdmin(c *gin.Context) {
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
	c.JSON(statusOK, response)
}

// @Summary		UploadCategoryImage
// @Description	Upload category images.
// @Tags			admin category management
// @Security		Bearer
// @Accept			mpfd
// @Produce		json
// @Param			categoryID		path		int					true	"Category ID"
// @Param			category-image	formData	file				true	"Image file to upload"
// @Success		201				{object}	response.Response	"success, image uploaded"
// @Failure		400				{object}	response.Response	"failed to get image from file"	or	"no files received to the server"
// @Failure		500				{object}	response.Response	"failed to upload image"
// @Router			/admin/category/add-image/{categoryID} [post]
func (ad *ProductHandler) UploadCategoryImage(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		response := response.ResponseMessage(statusBadRequest, "failed to get file from request", nil, err.Error())
		c.JSON(statusBadRequest, response)
		return
	}

	files := form.File["category-image"]
	if files == nil || len(files) == 0 {
		response := response.ResponseMessage(statusBadRequest, "no files received to the server", nil, "got 0 files for upload")
		c.JSON(statusBadRequest, response)
		return
	}

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
	}

	// Upload the file to specific dst.
	err = ad.productUseCase.UploadCategoryImage(files, categoryID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "failed to upload files", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "success, files uploaded", nil, nil)
	c.JSON(statusCreated, response)
}

// @Summary		UploadProductImages
// @Description	Upload product images.
// @Tags			admin product management
// @Security		Bearer
// @Accept			mpfd
// @Produce		json
// @Param			productID		path		int					true	"Product ID"
// @Param			product-image	formData	file				true	"Image file to upload"
// @Success		201				{object}	response.Response	"Success, images uploaded"
// @Failure		400				{object}	response.Response	"Failed to get image from file"	or	"No files received to the server"	or	"Invalid input"
// @Failure		500				{object}	response.Response	"Failed to upload image"
// @Router			/admin/product/add-images/{productID} [post]
func (ad *ProductHandler) UploadProductImages(c *gin.Context) {

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		response := response.ResponseMessage(statusBadRequest, "failed to get file from request", nil, err.Error())
		c.JSON(statusBadRequest, response)
		return
	}

	files := form.File["product-image"]
	if files == nil || len(files) == 0 {
		response := response.ResponseMessage(statusBadRequest, "no files received to the server", nil, "got 0 files for upload")
		c.JSON(statusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid input", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
	}

	// Upload the file to specific dst.
	err = ad.productUseCase.UploadProductImage(files, productID)
	if err != nil {
		response := response.ResponseMessage(statusInternalServerError, "failed to upload file", nil, err.Error())
		c.JSON(statusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(statusOK, "success, files uploaded", nil, nil)
	c.JSON(statusCreated, response)
}

// @Summary		Add to wishList
// @Description	Adds a product into wishlist.
// @Tags			wishlist
// @Security		Bearer
// @Produce		json
// @Param			productID	path		int	true	"Product ID"
// @Success		200			{object}	response.Response
// @Failure		400			{object}	response.Response
// @Failure		500			{object}	response.Response
// @Router			/wishlist/add/{productID} [post]
func (ph *ProductHandler) AddToWishList(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetIDFromContext(c)

	err = ph.productUseCase.AddToWishList(userID, productID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// @Summary		Add to wishList
// @Description	Adds a product into wishlist.
// @Tags			wishlist
// @Security		Bearer
// @Produce		json
// @Param			productID	path		int	true	"Product ID"
// @Success		200			{object}	response.Response
// @Failure		400			{object}	response.Response
// @Failure		500			{object}	response.Response
// @Router			/wishlist/remove/{productID} [delete]
func (ph *ProductHandler) RemoveFromWishList(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		response := response.ResponseMessage(400, "Invalid entry", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID, _ := helper.GetIDFromContext(c)

	err = ph.productUseCase.RemoveFromWishList(userID, productID)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", nil, nil)
	c.JSON(http.StatusOK, response)
}

// @Summary		Show wishlist products
// @Description	shows the products in users wishlist
// @Tags			wishlist
// @Security		Bearer
// @Produce		json
// @Param			page	query		int											true	"Page number"				default(1)
// @Param			count	query		int											true	"Number of items per page"	default(10)
// @Success		200		{object}	response.Response{data=response.Product}	"Success"
// @Failure		400		{object}	response.Response
// @Failure		500		{object}	response.Response
// @Router			/wishlist [get]
func (ph *ProductHandler) ShowWishListProducts(c *gin.Context) {
	userID, _ := helper.GetIDFromContext(c)
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

	products, err := ph.productUseCase.ShowWishListProducts(userID, page, count)
	if err != nil {
		response := response.ResponseMessage(500, "Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.ResponseMessage(200, "Success", products, nil)
	c.JSON(http.StatusOK, response)
}
