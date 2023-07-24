package usecase

import (
	"errors"
	"fmt"
	"log"

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

func (pu *productUseCase) CreateNewCategory(category request.Category) (response.Category, error) {

	DoCategoryExist, _ := pu.productRepo.FindCategoryByName(category.CategoryName)

	if DoCategoryExist.CategoryName != "" {
		return DoCategoryExist, errors.New("Category already exist")

	}

	NewCategory, err := pu.productRepo.CreateCategory(category)

	if err != nil {
		log.Println(" CREATE CATEGORY FAILED ")
		return response.Category{}, err
	}
	return NewCategory, nil

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
	ListOfAllCategories, err := pu.productRepo.ReadCategory(startIndex, endIndex)
	if err != nil {
		log.Println("  FAILED WHILE READING ALL CATEGORIES ")
		return nil, err
	}
	return ListOfAllCategories, nil

}
func (pu *productUseCase) UpdateCategoryWithId(ParamId int, category request.Category) (response.Category, error) {

	UpdatedCategory, err := pu.productRepo.UpdateCategory(ParamId, category)

	if err != nil {
		log.Println("FAILED WHILE UPDATING CATEGORY ")
		return response.Category{}, err
	}
	if UpdatedCategory.Id == 0 {
		return response.Category{}, errors.New("Category not found")
	}
	return UpdatedCategory, nil

}

func (pu *productUseCase) BlockCategoryWithId(ParamId int) (response.Category, error) {
	BlockedCategory, err := pu.productRepo.BlockCategoryFromDatabase(ParamId)
	if err != nil {
		log.Println(" FAILED WHILE DELETE CATEGORY ")
		return response.Category{}, err
	}
	if BlockedCategory.Id == 0 {

		return response.Category{}, errors.New("Category not found ")
	}

	return BlockedCategory, nil

}

func (pu *productUseCase) UnBlockCategoryWithId(ParamId int) (response.Category, error) {
	UnBlockedCategory, err := pu.productRepo.BlockCategoryFromDatabase(ParamId)
	if err != nil {
		log.Println(" FAILED WHILE DELETE CATEGORY ")
		return response.Category{}, err
	}
	if UnBlockedCategory.Id == 0 {

		return response.Category{}, errors.New("Category not found ")
	}

	return UnBlockedCategory, nil

}

// -----------------------------------------------------------------------------------------------------------------------------
func (pu *productUseCase) CreateNewProduct(product request.Product) (response.Product, error) {

	log.Println("USECASE :", product)

	ResultOfFinding, err := pu.productRepo.FindCategoryById(product.CategoryID)
	if err != nil {
		return response.Product{}, err
	} else if ResultOfFinding.Id == 0 {
		return response.Product{}, errors.New("Category Not found")
	}

	product.Brand = ResultOfFinding.CategoryName
	product.SKU = helper.MakeSku(product.ProductName)
	fmt.Printf("PRODUCT.BRAND =  %s   RESULTOFFINDING.CATEGORYNAME  =  %s ", product.Brand, ResultOfFinding.CategoryName)

	ProductExist, err := pu.productRepo.FindProductByName(product.ProductName)
	if err != nil {
		return response.Product{}, errors.New("Error while Finding product by name ")
	}
	if ProductExist.ID != 0 {
		return response.Product{}, errors.New("Product already exist")
	}

	NewProduct, err := pu.productRepo.InsertNewProductToDatabase(product)
	if err != nil {
		return response.Product{}, err
	}
	if NewProduct.ID == 0 {
		return response.Product{}, errors.New(" Failed to  create product ")
	}
	return NewProduct, nil
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
	ListOfAllProducts, err := pu.productRepo.ViewAllProductsToAdmin(startIndex, endIndex)
	if err != nil {
		return []response.Product{}, err
	}
	return ListOfAllProducts, nil

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
	ListOfAllProducts, err := pu.productRepo.ViewAllProductsToUser(startIndex, endIndex)
	if err != nil {
		return []response.Product{}, err
	}
	return ListOfAllProducts, nil

}

func (pu *productUseCase) UpdateProductWithId(paramId int, updations request.Product) (response.Product, error) {

	UpdatedProduct, err := pu.productRepo.UpdateProductToDatabase(paramId, updations)
	if err != nil {
		return response.Product{}, err
	}
	return UpdatedProduct, nil
}
func (pu *productUseCase) BlockProductWithId(paramId int) (response.Product, error) {
	BlockedProduct, err := pu.productRepo.BlockProductFromDatabase(paramId)
	if err != nil {
		return response.Product{}, errors.New("Failed while blocking Product ")
	}
	return BlockedProduct, nil
}

func (pu *productUseCase) UnBlockProductWithId(paramId int) (response.Product, error) {
	BlockedProduct, err := pu.productRepo.UnblockProductFromDatabase(paramId)
	if err != nil {
		return response.Product{}, errors.New("Failed while unblocking Product failed")
	}
	return BlockedProduct, nil
}

func (pd *productUseCase) ViewProductById(productId int) (response.ProductItem, error) {
	Product, err := pd.productRepo.FindProductById(productId)
	if err != nil {
		return response.ProductItem{}, fmt.Errorf("Failed to find product : %s", err)
	}
	if Product.ID == 0 {
		return response.ProductItem{}, fmt.Errorf("unable to retrieve item from database")
	}

	Ratings, err := pd.productRepo.GetProductReviews(productId)
	if err != nil {
		return response.ProductItem{}, fmt.Errorf("Failed to find product reviews : %s", err)
	}

	var ProductData response.ProductItem
	ProductData.ID = Product.ID
	ProductData.CategoryID = Product.CategoryID
	ProductData.Product_Name = Product.ProductName
	ProductData.Price = Product.Price
	ProductData.SKU = Product.SKU
	ProductData.Brand = Product.Brand
	ProductData.Product_Description = Product.Product_Description
	ProductData.Product_Image = Product.Product_Image
	ProductData.Is_Blocked = Product.IsBlocked
	ProductData.RatingAndReviews = Ratings

	return ProductData, nil

}

func (pu *productUseCase) ValidateProductRatingRequest(userID, productID int) error {
	Rating, err := pu.productRepo.FindUserRatingOnProduct(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to find user rating")
	}
	if Rating.ID != 0 {
		return fmt.Errorf("User already done rating on this product")
	}

	OrderData, err := pu.orderRepo.FindOrderDataByUseridAndProductid(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to find order details")
	}
	if OrderData.ID == 0 {
		return fmt.Errorf("User have not purchased the product")
	}
	status, err := pu.orderRepo.FindOrderStatusById(OrderData.OrderStatusId)
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

	Rating, err := pu.productRepo.FindUserRatingOnProduct(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to find product rating :%s", err)
	}

	if Rating.ID == 0 {
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
	Products, err := pu.productRepo.SearchProducts(search, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to search products  :%s", err)
	}
	if len(Products) == 0 {
		return nil, fmt.Errorf("Product not found")
	}

	return Products, nil
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
