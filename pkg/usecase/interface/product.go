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
	DisplayAllProductsToUser(userID, page, count int) ([]response.Product, error)
	UpdateProductByID(productID int, updated request.UpdateProduct) error
	BlockProductByID(productID int) error
	UnBlockProductByID(productID int) error
	ValidateProductRatingRequest(userID, productID int) error
	InsertNewProductRating(userID, productID int, rating request.Rating) error
	SearchProducts(search string, page, count int) ([]response.Product, error)
	GetProductsByCategory(categoryID, page, count int) ([]response.Product, error)

	UploadCategoryImage(files []*multipart.FileHeader, ID int) error
	UploadProductImage(files []*multipart.FileHeader, ID int) error
	ViewIndividualProduct(userID, productID int) (response.ProductItem, error)

	AddToWishList(userID, productID int) error
	RemoveFromWishList(userID, productID int) error
	ShowWishListProducts(userID, page, count int) ([]response.Product, error)
}
