package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductRepository interface {
	CreateCategory(category request.Category) (response.Category, error)
	ReadCategory(startIndex int, endIndex int) ([]response.Category, error)
	UpdateCategory(ParamId int, category request.Category) (response.Category, error)
	BlockCategoryFromDatabase(ParamId int) (response.Category, error)
	UnBlockCategoryFromDatabase(ParamId int) (response.Category, error)
	FindCategoryByName(name string) (response.Category, error)
	FindCategoryByID(ID int) (response.Category, error)

	InsertNewProductToDatabase(product request.Product) (response.Product, error)
	ViewAllProductsToAdmin(startIndex, endIndex int) ([]response.Product, error)
	ViewAllProductsToUser(startIndex, endIndex int) ([]response.Product, error)
	UpdateProductToDatabase(paramId int, product request.Product) (response.Product, error)
	BlockProductFromDatabase(paramId int) (response.Product, error)
	UnblockProductFromDatabase(paramId int) (response.Product, error)
	FindProductByName(paramName string) (response.Product, error)
	FindProductById(productid int) (response.Product, error)

	FindUserRatingOnProduct(userID, productID int) (response.Rating, error)
	InsertProductRating(rating request.Rating) error
	GetProductReviews(productID int) ([]response.Rating, error)
	SearchProducts(search string, startIndex, endIndex int) ([]response.Product, error)
	GetProductsByCategory(categoryID int, startIndex, endIndex int) ([]response.Product, error)
}
