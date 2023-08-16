package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
	orderRepo   interfaces.OrderRepository
}

func NewProductUseCase(productRepo interfaces.ProductRepository, orderRepo interfaces.OrderRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}

}

func (pu *productUseCase) CreateNewCategory(category request.Category) error {
	doCategoryExist, err := pu.productRepo.FindCategoryByName(category.CategoryName)
	if doCategoryExist.ID != 0 {
		return fmt.Errorf("Category already exist with this name")
	}

	NewCategory, err := pu.productRepo.CreateCategory(category)
	if err != nil {
		return fmt.Errorf("Failed to create category : %s", err)
	}

	if NewCategory.ID == 0 {
		return fmt.Errorf("Failed to verify created category")
	}

	return nil
}

func (pu *productUseCase) ReadAllCategories(page int, count int) ([]response.Category, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}

	startIndex := (page - 1) * count
	endIndex := startIndex + count

	listOfAllCategories, err := pu.productRepo.ReadCategory(startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to find categories :%s", err)
	}

	return listOfAllCategories, nil
}

func (pu *productUseCase) UpdateCategoryWithID(productID int, category request.Category) error {
	updatedCategory, err := pu.productRepo.UpdateCategory(productID, category)
	if err != nil {
		return fmt.Errorf("Failed to update category :%s", err)
	}

	if updatedCategory.ID == 0 {
		return fmt.Errorf("Failed to verify updated category")
	}

	return nil
}

func (pu *productUseCase) BlockCategoryWithID(categoryID int) error {
	blockedCategory, err := pu.productRepo.BlockCategoryFromDatabase(categoryID)
	if err != nil {
		return fmt.Errorf("Failed to block category :%s", err)
	}

	if blockedCategory.ID == 0 {
		return fmt.Errorf("Failed to verify the blocked category")
	}

	return nil
}

func (pu *productUseCase) UnBlockCategoryWithID(categoryID int) error {
	unBlockedCategory, err := pu.productRepo.BlockCategoryFromDatabase(categoryID)
	if err != nil {
		return fmt.Errorf("Failed to block category :%s", err)
	}

	if unBlockedCategory.ID == 0 {
		return fmt.Errorf("Failed to verify blocked category")
	}

	return nil
}
func (pu *productUseCase) CreateNewProduct(product request.Product) error {
	category, err := pu.productRepo.FindCategoryByID(product.CategoryID)
	if err != nil {
		return fmt.Errorf("failed to find category: %s", err)
	}
	if category.ID == 0 {
		return fmt.Errorf("category not found")
	}

	product.Brand = category.CategoryName
	product.SKU = helper.MakeSKU(product.ProductName)

	existingProduct, err := pu.productRepo.FindProductByName(product.ProductName)
	if err != nil {
		return fmt.Errorf("failed to find product by name: %s", err)
	}
	if existingProduct.ID != 0 {
		return fmt.Errorf("product already exists with the same name")
	}

	newProduct, err := pu.productRepo.InsertNewProductToDatabase(product)
	if err != nil {
		return fmt.Errorf("failed to create new product: %s", err)
	}
	if newProduct.ID == 0 {
		return fmt.Errorf("failed to verify created product")
	}

	return nil
}

func (pu *productUseCase) DisplayAllProductsToAdmin(page, count int) ([]response.Product, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}

	startIndex := (page - 1) * count
	endIndex := startIndex + count

	listOfAllProducts, err := pu.productRepo.ViewAllProductsToAdmin(startIndex, endIndex)
	if err != nil {
		return []response.Product{}, err
	}

	return listOfAllProducts, nil
}

func (pu *productUseCase) DisplayAllAvailabeProductsToUser(page, count int) ([]response.Product, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}

	startIndex := (page - 1) * count
	endIndex := startIndex + count

	listOfAllProducts, err := pu.productRepo.ViewAllProductsToUser(startIndex, endIndex)
	if err != nil {
		return []response.Product{}, err
	}

	return listOfAllProducts, nil
}

func (pu *productUseCase) UpdateProductWithID(productID int, update request.Product) error {
	updatedProduct, err := pu.productRepo.UpdateProductToDatabase(productID, update)
	if err != nil {
		return fmt.Errorf("Failed to update product :%s", err)
	}
	if updatedProduct.ID == 0 {
		return fmt.Errorf("Failed to verify the updated product")
	}
	return nil
}

func (pu *productUseCase) BlockProductWithID(productID int) error {
	blockedProduct, err := pu.productRepo.BlockProductFromDatabase(productID)
	if err != nil {
		return fmt.Errorf("Failed to block product :%s", err)
	}
	if blockedProduct.ID == 0 {
		return fmt.Errorf("Failed to verify updated product")
	}
	return nil
}

func (pu *productUseCase) UnBlockProductWithID(productID int) error {
	unBlockedProduct, err := pu.productRepo.UnblockProductFromDatabase(productID)
	if err != nil {
		return fmt.Errorf("Failed unblock product :%s", err)
	}
	if unBlockedProduct.ID == 0 {
		return fmt.Errorf("Failed to verify unblocked product")
	}
	return nil
}

func (pd *productUseCase) ViewProductByID(productID int) (response.ProductItem, error) {
	product, err := pd.productRepo.FindProductById(productID)
	if err != nil {
		return response.ProductItem{}, fmt.Errorf("Failed to find product :%s", err)
	}
	if product.ID == 0 {
		return response.ProductItem{}, fmt.Errorf("Failed to fetch product")
	}

	ratings, err := pd.productRepo.GetProductReviews(productID)
	if err != nil {
		return response.ProductItem{}, fmt.Errorf("Failed to find product reviews :%s", err)
	}

	return response.ProductItem{
		ID:                  product.ID,
		CategoryID:          product.CategoryID,
		Product_Name:        product.ProductName,
		Price:               product.Price,
		SKU:                 product.SKU,
		Brand:               product.Brand,
		Product_Description: product.Product_Description,
		Product_Image:       product.Product_Image,
		Is_Blocked:          product.IsBlocked,
		RatingAndReviews:    ratings,
	}, nil

}

func (pu *productUseCase) ValidateProductRatingRequest(userID, productID int) error {
	rating, err := pu.productRepo.FindUserRatingOnProduct(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to find user rating")
	}
	if rating.ID != 0 {
		return fmt.Errorf("User already done rating on this product")
	}

	orderData, err := pu.orderRepo.FindOrderDataByUseridAndProductid(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to find order details")
	}
	if orderData.ID == 0 {
		return fmt.Errorf("User have not purchased the product")
	}
	status, err := pu.orderRepo.FindOrderStatusById(orderData.OrderStatusId)
	if err != nil {
		return fmt.Errorf("Failed to find order status")
	}

	if status != "Delivered" {
		return fmt.Errorf("Only delivered purchase can do rating.")
	}
	return nil
}

func (pu *productUseCase) InsertNewProductRating(userID int, productID int, rating request.Rating) error {
	rating.UserID = userID
	rating.ProductID = productID
	err := pu.productRepo.InsertProductRating(rating)
	if err != nil {
		return fmt.Errorf("Failed to insert product rating :%s", err)
	}

	userRating, err := pu.productRepo.FindUserRatingOnProduct(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to find product rating :%s", err)
	}

	if userRating.ID == 0 {
		return fmt.Errorf("Inserted product rating not found ")
	}
	return nil
}

func (pu *productUseCase) SearchProducts(search string, page, count int) ([]response.Product, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}
	startIndex := (page - 1) * count
	endIndex := startIndex + count

	products, err := pu.productRepo.SearchProducts(search, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to search products  :%s", err)
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("Product not found")
	}

	return products, nil
}

func (pu *productUseCase) GetProductsByCategory(categoryID int, page, count int) ([]response.Product, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}
	startIndex := (page - 1) * count
	endIndex := startIndex + count
	products, err := pu.productRepo.GetProductsByCategory(categoryID, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to get products by category : %s", err)
	}
	if len(products) == 0 {
		return nil, fmt.Errorf(" no product found  ")
	}

	return products, nil

}
