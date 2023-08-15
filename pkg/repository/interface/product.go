package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductRepository interface {
	CreateCategory(category request.Category) (response.Category, error)
	ReadCategory(startIndex int, endIndex int) ([]response.Category, error)
	UpdateCategory(categoryID int, category request.Category) (response.Category, error)
	BlockCategoryFromDatabase(categoryID int) (response.Category, error)
	UnBlockCategoryFromDatabase(categoryID int) (response.Category, error)
	FindCategoryByName(name string) (response.Category, error)
	FindCategoryByID(categoryID int) (response.Category, error)

	InsertNewProductToDatabase(product request.Product) (response.Product, error)
	ViewAllProductsToAdmin(startIndex, endIndex int) ([]response.Product, error)
	ViewAllProductsToUser(startIndex, endIndex int) ([]response.Product, error)
	UpdateProductToDatabase(productID int, product request.Product) (response.Product, error)
	BlockProductFromDatabase(productID int) (response.Product, error)
	UnblockProductFromDatabase(productID int) (response.Product, error)
	FindProductByName(paramName string) (response.Product, error)
	FindProductById(productID int) (response.Product, error)

	FindUserRatingOnProduct(userID, productID int) (response.Rating, error)
	InsertProductRating(rating request.Rating) error
	GetProductReviews(productID int) ([]response.Rating, error)
	SearchProducts(search string, startIndex, endIndex int) ([]response.Product, error)
	GetProductsByCategory(categoryID int, startIndex, endIndex int) ([]response.Product, error)
}
