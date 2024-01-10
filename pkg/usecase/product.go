package usecase

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/google/uuid"
)

var (
	ErrRecordAlreadyExist = errors.New("record already exist")
	ErrCategoryNotFound   = errors.New("referenced category not found")
	ErrNoRecord           = errors.New("record not found")
	ErrInProcessing       = errors.New("currently processing the order, not completed yet")
)

const (
	statusDelivered = "Delivered"
	statusCancelled = "Cancelled"
	statusReturned  = "Returned"
)

const (
	category = "category"
	product  = "product"
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

func (pu *productUseCase) CreateCategory(category request.Category) (response.Category, error) {
	existingCategory, err := pu.productRepo.FindCategoryByName(category.CategoryName)
	if existingCategory.ID != 0 {
		return response.Category{}, ErrRecordAlreadyExist
	}

	result, err := pu.productRepo.CreateCategory(category)
	if err != nil {
		return response.Category{}, fmt.Errorf("Failed to create category : %s", err)
	}

	return result, nil
}

func (pu *productUseCase) ReadAllCategories(page int, count int) ([]response.Category, error) {

	startIndex, endIndex := helper.Paginate(page, count)

	listOfAllCategories, err := pu.productRepo.ReadCategory(startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to find categories :%s", err)
	}

	return listOfAllCategories, nil
}

func (pu *productUseCase) UpdateCategoryByID(productID int, category request.Category) error {
	err := pu.productRepo.UpdateCategory(productID, category)
	if err != nil {
		return fmt.Errorf("Failed to update category :%s", err)
	}

	return nil
}

func (pu *productUseCase) BlockCategoryByID(categoryID int) error {
	err := pu.productRepo.BlockCategoryByID(categoryID)
	if err != nil {
		return fmt.Errorf("Failed to block category :%s", err)
	}

	return nil
}

func (pu *productUseCase) UnBlockCategoryByID(categoryID int) error {
	err := pu.productRepo.UnBlockCategoryByID(categoryID)
	if err != nil {
		return fmt.Errorf("Failed to block category :%s", err)
	}

	return nil
}

func (pu *productUseCase) CreateProduct(product request.Product) (response.Product, error) {
	category, err := pu.productRepo.FindCategoryByID(product.CategoryID)
	if err != nil {
		return response.Product{}, fmt.Errorf("failed to find category: %s", err)
	}
	if category.ID == 0 {
		return response.Product{}, ErrCategoryNotFound
	}

	product.Brand = category.Category_Name
	product.SKU = helper.MakeSKU(product.ProductName)

	existingProduct, err := pu.productRepo.FindProductByName(product.ProductName)
	if err != nil {
		return response.Product{}, fmt.Errorf("failed to find product by name: %s", err)
	}
	if existingProduct.ID != 0 {
		return response.Product{}, ErrRecordAlreadyExist
	}

	result, err := pu.productRepo.CreateProduct(product)
	if err != nil {
		return response.Product{}, fmt.Errorf("failed to create new product: %s", err)
	}

	return result, nil
}

func (pu *productUseCase) DisplayAllProductsToAdmin(page, count int) ([]response.Product, error) {
	startIndex, endIndex := helper.Paginate(page, count)

	products, err := pu.productRepo.ViewAllProductsToAdmin(startIndex, endIndex)
	if err != nil {
		return []response.Product{}, err
	}

	return products, nil
}

func (pu *productUseCase) DisplayAllProductsToUser(userID, page, count int) ([]response.Product, error) {
	startIndex, endIndex := helper.Paginate(page, count)

	listOfAllProducts, err := pu.productRepo.ViewAllProductsToUser(userID, startIndex, endIndex)
	if err != nil {
		return []response.Product{}, err
	}

	return listOfAllProducts, nil
}

func (pu *productUseCase) UpdateProductByID(productID int, update request.UpdateProduct) error {
	err := pu.productRepo.UpdateProduct(productID, update)
	if err != nil {
		return fmt.Errorf("Failed to update product :%s", err)
	}

	return nil
}

func (pu *productUseCase) BlockProductByID(productID int) error {
	err := pu.productRepo.BlockProduct(productID)
	if err != nil {
		return fmt.Errorf("Failed to block product :%s", err)
	}

	return nil
}

func (pu *productUseCase) UnBlockProductByID(productID int) error {
	err := pu.productRepo.UnblockProduct(productID)
	if err != nil {
		return fmt.Errorf("Failed unblock product :%s", err)
	}
	return nil
}

func (pd *productUseCase) ViewIndividualProduct(userID, productID int) (response.ProductItem, error) {

	product, err := pd.productRepo.ViewIndividualProduct(userID, productID)
	if err != nil {
		return response.ProductItem{}, fmt.Errorf("Failed to find product :%s", err)
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
		Images:              product.Images,
		IsWishlisted:        product.IsWishlisted,
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
		return ErrRecordAlreadyExist
	}

	orderData, err := pu.orderRepo.FindOrderByUserIDAndProductID(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to find order details")
	}

	if orderData.ID == 0 {
		return ErrNoRecord
	}

	status, err := pu.orderRepo.FindOrderStatusByID(orderData.OrderStatusID)
	if err != nil {
		return fmt.Errorf("Failed to find order status")
	}

	if status != statusDelivered || status != statusReturned {
		return ErrInProcessing
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
	startIndex, endIndex := helper.Paginate(page, count)

	products, err := pu.productRepo.SearchProducts(search, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to search products  :%s", err)
	}

	return products, nil
}

func (pu *productUseCase) GetProductsByCategory(categoryID int, page, count int) ([]response.Product, error) {
	startIndex, endIndex := helper.Paginate(page, count)

	products, err := pu.productRepo.GetProductsByCategory(categoryID, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to get products by category : %s", err)
	}

	return products, nil

}

func (pu *productUseCase) UploadCategoryImage(files []*multipart.FileHeader, categoryID int) error {

	err := pu.uploadImage(files, category, categoryID)
	if err != nil {
		return err
	}
	return nil
}

func (pu *productUseCase) UploadProductImage(files []*multipart.FileHeader, productID int) error {

	err := pu.uploadImage(files, product, productID)
	if err != nil {
		return err
	}
	return nil
}

func (pu *productUseCase) uploadImage(files []*multipart.FileHeader, imageFor string, ID int) error {

	imageURLs := make([]string, len(files))
	for i, f := range files {
		d, err := f.Open()
		if err != nil {
			return err
		}

		data, err := io.ReadAll(d)
		if err != nil {
			return err
		}
		filename := uuid.New().String() + f.Filename
		locationURL, err := helper.UploadMediaToS3(data, filename)
		if err != nil {
			return err
		}
		imageURLs[i] = locationURL
	}

	if imageFor == category {
		err := pu.productRepo.InsertCategoryIMG(imageURLs, ID)
		if err != nil {
			return err
		}
	}

	if imageFor == product {
		err := pu.productRepo.InsertProductIMG(imageURLs, ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pu *productUseCase) AddToWishList(userID, productID int) error {
	return pu.productRepo.AddToWishList(userID, productID)
}

func (pu *productUseCase) RemoveFromWishList(userID, productID int) error {
	return pu.productRepo.RemoveFromWishList(userID, productID)
}

func (pu *productUseCase) ShowWishListProducts(userID, page, count int) ([]response.Product, error) {
	startIndex, endIndex := helper.Paginate(page, count)
	return pu.productRepo.ShowWishListProducts(userID, startIndex, endIndex)
}
