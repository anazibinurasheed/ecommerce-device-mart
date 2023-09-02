package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type ProductUseCase interface {
	CreateNewCategory(category request.Category) error
	ReadAllCategories(page, count int) ([]response.Category, error)
	UpdateCategoryWithID(categoryID int, category request.Category) error
	BlockCategoryWithID(categoryID int) error
	UnBlockCategoryWithID(categoryID int) error
	CreateNewProduct(product request.Product) error
	DisplayAllProductsToAdmin(page, count int) ([]response.Product, error)
	DisplayAllAvailabeProductsToUser(page, count int) ([]response.Product, error)
	UpdateProductWithID(productID int, updated request.Product) error
	BlockProductWithID(productID int) error
	UnBlockProductWithID(productID int) error
	ViewProductByID(productID int) (response.ProductItem, error)
	ValidateProductRatingRequest(userID, productID int) error
	InsertNewProductRating(userID, productID int, rating request.Rating) error
	SearchProducts(search string, page, count int) ([]response.Product, error)
	GetProductsByCategory(categoryID, page, count int) ([]response.Product, error)
}
