package interfaces

import (
	"mime/multipart"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductUseCase interface {
	CreateCategory(category request.Category) (response.Category, error)
	ReadAllCategories(page, count int) ([]response.Category, error)
	UpdateCategoryByID(categoryID int, category request.Category) error
	BlockCategoryByID(categoryID int) error
	UnBlockCategoryByID(categoryID int) error
	CreateProduct(product request.Product) (response.Product, error)
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
	
	UploadCategoryImage(files []*multipart.FileHeader, ID int) error
	UploadProductImage(files []*multipart.FileHeader, ID int) error
	GetCategoryImage(categoryID int) (response.CategoryImage, error)
	GetProductImages(productID int) ([]response.ProductImages, error)
}
