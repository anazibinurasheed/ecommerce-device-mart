package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductRepository interface {
	CreateCategory(category request.Category) (response.Category, error)
	ReadCategory(startIndex int, endIndex int) ([]response.Category, error)
	UpdateCategory(categoryID int, category request.Category) error
	BlockCategoryByID(categoryID int) error
	UnBlockCategoryByID(categoryID int) error
	FindCategoryByName(name string) (response.Category, error)
	FindCategoryByID(categoryID int) (response.Category, error)

	CreateProduct(product request.Product) (response.Product, error)
	ViewAllProductsToAdmin(startIndex, endIndex int) ([]response.Product, error)
	ViewAllProductsToUser(userID, startIndex, endIndex int) ([]response.Product, error)
	UpdateProduct(productID int, product request.UpdateProduct) error
	BlockProduct(productID int) error
	UnblockProduct(productID int) error
	FindProductByName(paramName string) (response.Product, error)
	FindProductByID(productID int) (response.Product, error)
	ViewIndividualProduct(userID, productID int) (response.Product, error)

	FindUserRatingOnProduct(userID, productID int) (response.Rating, error)
	InsertProductRating(rating request.Rating) error
	GetProductReviews(productID int) ([]response.Rating, error)
	SearchProducts(search string, startIndex, endIndex int) ([]response.Product, error)
	GetProductsByCategory(categoryID int, startIndex, endIndex int) ([]response.Product, error)

	InsertCategoryIMG(urls interface{}, categoryID int) error
	InsertProductIMG(urls interface{}, productID int) error

	AddToWishList(userID, productID int) error
	RemoveFromWishList(userID, productID int) error
	ShowWishListProducts(userID, page, count int) ([]response.Product, error)
}
