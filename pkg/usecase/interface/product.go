package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductUseCase interface {
	CreateNewCategory(category request.Category) (response.Category, error)
	ReadAllCategories(page int, count int) ([]response.Category, error)
	UpdateCategoryWithId(ParamId int, category request.Category) (response.Category, error)
	BlockCategoryWithId(ParamId int) (response.Category, error)
	UnBlockCategoryWithId(ParamId int) (response.Category, error)
	CreateNewProduct(product request.Product) (response.Product, error)
	DisplayAllProductsToAdmin(page, count int) ([]response.Product, error)
	DisplayAllAvailabeProductsToUser(page, count int) ([]response.Product, error)
	UpdateProductWithId(paramId int, updations request.Product) (response.Product, error)
	BlockProductWithId(paramId int) (response.Product, error)
	UnBlockProductWithId(paramId int) (response.Product, error)
	ViewProductById(productId int) (response.ProductItem, error)
	ValidateProductRatingRequest(userID, productID int) error
	InsertNewProductRating(userID int, productID int, rating request.Rating) error
	SearchProducts(search string, page, count int) ([]response.Product, error)
	GetProductsByCategory(categoryID int, page, count int) ([]response.Product, error)
}
