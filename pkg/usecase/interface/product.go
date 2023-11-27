package interfaces

import (
	"mime/multipart"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductUseCase interface {
	CreateCategory(category request.Category) error
	ReadAllCategories(page, count int) ([]response.Category, error)
	UpdateCategoryByID(categoryID int, category request.Category) error
	BlockCategoryByID(categoryID int) error
	UnBlockCategoryByID(categoryID int) error
	CreateProduct(product request.Product) error
	DisplayAllProductsToAdmin(page, count int) ([]response.Product, error)
	DisplayAllAvailableProductsToUser(page, count int) ([]response.Product, error)
	UpdateProductByID(productID int, updated request.Product) error
	BlockProductByID(productID int) error
	UnBlockProductByID(productID int) error
	ViewProductByID(productID int) (response.ProductItem, error)
	ValidateProductRatingRequest(userID, productID int) error
	InsertNewProductRating(userID, productID int, rating request.Rating) error
	SearchProducts(search string, page, count int) ([]response.Product, error)
	GetProductsByCategory(categoryID, page, count int) ([]response.Product, error)
	UploadImage(files []*multipart.FileHeader, imageUUID string) error
}
