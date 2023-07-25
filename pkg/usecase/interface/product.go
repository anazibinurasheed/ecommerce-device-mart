package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductUseCase interface {
	CreateNewCategory(category request.Category) error
	ReadAllCategories(page int, count int) ([]response.Category, error)
	UpdateCategoryWithId(ParamId int, category request.Category) error
	BlockCategoryWithId(ParamId int) error
	UnBlockCategoryWithId(ParamId int) error
	CreateNewProduct(product request.Product) error
	DisplayAllProductsToAdmin(page, count int) ([]response.Product, error)
	DisplayAllAvailabeProductsToUser(page, count int) ([]response.Product, error)
	UpdateProductWithId(paramId int, updations request.Product) error
	BlockProductWithId(paramId int) error
	UnBlockProductWithId(paramId int) error
	ViewProductById(productId int) (response.ProductItem, error)
	ValidateProductRatingRequest(userID, productID int) error
	InsertNewProductRating(userID int, productID int, rating request.Rating) error
	SearchProducts(search string, page, count int) ([]response.Product, error)
	GetProductsByCategory(categoryID int, page, count int) ([]response.Product, error)
}
