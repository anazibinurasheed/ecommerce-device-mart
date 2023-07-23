package repository

import (
	"fmt"
	"log"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB: DB}
}

func (pd *productDatabase) CreateCategory(category request.Category) (response.Category, error) {
	var NewCategory response.Category
	query := "INSERT INTO Categories (Category_Name) VALUES($1) RETURNING *;"
	err := pd.DB.Raw(query, category.CategoryName).Scan(&NewCategory).Error

	return NewCategory, err
}

func (pd *productDatabase) ReadCategory(startIndex int, endIndex int) ([]response.Category, error) {
	var ListOfAllCategories = make([]response.Category, 0)

	query := "SELECT * FROM Categories ORDER BY Category_Name OFFSET $1 FETCH NEXT $2 ROW ONLY ;"

	err := pd.DB.Raw(query, startIndex, endIndex).Scan(&ListOfAllCategories).Error
	return ListOfAllCategories, err
}

func (pd *productDatabase) UpdateCategory(ParamId int, category request.Category) (response.Category, error) {
	var UpdatedCategory response.Category

	query := "UPDATE Categories SET Category_Name = $1 WHERE ID = $2 RETURNING *;"

	err := pd.DB.Raw(query, category.CategoryName, ParamId).Scan(&UpdatedCategory).Error

	return UpdatedCategory, err
}

func (pd *productDatabase) BlockCategoryFromDatabase(ParamId int) (response.Category, error) {
	var BlockedCategory response.Category
	query := "Update Categories SET Is_blocked = true WHERE ID = ? RETURNING *;"
	err := pd.DB.Raw(query, ParamId).Scan(&BlockedCategory).Error
	return BlockedCategory, err

}

func (pd *productDatabase) UnBlockCategoryFromDatabase(ParamId int) (response.Category, error) {
	var UnBlockedCategory response.Category
	query := "Update Categories SET Is_blocked = false WHERE ID = ? RETURNING *;"
	err := pd.DB.Raw(query, ParamId).Scan(&UnBlockedCategory).Error
	return UnBlockedCategory, err

}

//-------------------------------------------------------------------------------------------------------------
func (pd *productDatabase) FindCategoryByName(name string) (response.Category, error) {
	var ResultOfFinding response.Category
	query := "SELECT * FROM Categories WHERE Category_Name = $1"
	err := pd.DB.Raw(query, name).Scan(&ResultOfFinding).Error
	return ResultOfFinding, err

}
func (pd *productDatabase) FindCategoryById(id int) (response.Category, error) {
	var ResultOfFinding response.Category
	log.Println("FIND CATEGORY BY ID :", id)
	query := "SELECT * FROM Categories WHERE Id = $1 "
	err := pd.DB.Raw(query, id).Scan(&ResultOfFinding).Error
	return ResultOfFinding, err
}

//product
//--------------------------------------------------------------------------------------------------------------
func (pd *productDatabase) InsertNewProductToDatabase(product request.Product) (response.Product, error) {
	log.Println("REPOSITORY :", product)
	fmt.Println("REPOSITORY :", product)
	var NewProduct response.Product
	query := "INSERT INTO Products (Category_ID,Product_Name,Price,Product_Description,Product_Image,Brand,Sku,is_blocked) Values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *"
	err := pd.DB.Raw(query, product.CategoryID, product.ProductName, product.Price, product.Product_Description, product.ProductImage, product.Brand, product.SKU, product.IsBlocked).Scan(&NewProduct).Error
	return NewProduct, err
}

func (pd *productDatabase) ViewAllProductsToAdmin(startIndex, endIndex int) ([]response.Product, error) {
	var ListOfAllProducts = make([]response.Product, 0)

	query := "SELECT * FROM Products OFFSET $1 FETCH NEXT $2 ROW ONLY ;"
	err := pd.DB.Raw(query, startIndex, endIndex).Scan(&ListOfAllProducts).Error
	return ListOfAllProducts, err

}

func (pd *productDatabase) UpdateProductToDatabase(paramId int, updations request.Product) (response.Product, error) {
	var UpdatedProduct response.Product
	query := "Update Products SET Category_ID = $1 ,Product_Name = $2 ,Product_Description = $3 ,Product_Image = $4 , Price = $5 WHERE ID = $6 RETURNING *"
	err := pd.DB.Raw(query, updations.CategoryID, updations.ProductName, updations.Product_Description, updations.ProductImage, updations.Price, paramId).Scan(&UpdatedProduct).Error
	return UpdatedProduct, err

}

func (pd *productDatabase) BlockProductFromDatabase(paramId int) (response.Product, error) {
	var BlockedProduct response.Product
	status := true
	query := "UPDATE Products SET Is_Blocked = $1 WHERE ID = $2 RETURNING *;"
	err := pd.DB.Raw(query, status, paramId).Scan(&BlockedProduct).Error
	fmt.Println(BlockedProduct)
	return BlockedProduct, err

}

func (pd *productDatabase) UnblockProductFromDatabase(paramId int) (response.Product, error) {
	var Product response.Product
	status := false
	query := "UPDATE Products SET is_blocked = $1 WHERE id = $2 RETURNING *;"
	err := pd.DB.Raw(query, status, paramId).Scan(&Product).Error
	fmt.Println(Product)

	return Product, err

}

func (pd *productDatabase) FindProductByName(productName string) (response.Product, error) {
	var NewProduct response.Product
	query := "SELECT * FROM Products WHERE Product_name = $1  FETCH FIRST 1 ROW ONLY"
	err := pd.DB.Raw(query, productName).Scan(&NewProduct).Error
	return NewProduct, err
}

func (pd *productDatabase) ViewAllProductsToUser(startIndex, endIndex int) ([]response.Product, error) {
	var ListOfAllProducts = make([]response.Product, 0)

	query := "SELECT * FROM Products  OFFSET $1 FETCH NEXT $2 ROW ONLY;"
	err := pd.DB.Raw(query, startIndex, endIndex).Scan(&ListOfAllProducts).Error
	return ListOfAllProducts, err
}

//-------------------------------------------------------------------------------------------------------------------------

func (pd *productDatabase) FindProductById(productid int) (response.Product, error) {
	var Product response.Product
	query := "SELECT * FROM Products WHERE Id = $1  FETCH FIRST 1 ROW ONLY"
	err := pd.DB.Raw(query, productid).Scan(&Product).Error
	return Product, err
}

func (pd *productDatabase) FindUserRatingOnProduct(userID, productID int) (response.Rating, error) {
	var Rating response.Rating

	query := `SELECT r.id,r.rating,r.description,u.user_name FROM ratings r 
	INNER JOIN users u ON r.user_id = u.id WHERE r.user_id = $1 AND r.product_id = $2;`
	err := pd.DB.Raw(query, userID, productID).Scan(&Rating).Error
	return Rating, err
}

// func (pd *productDatabase) InsertProductRating(userID int, productID int, rating request.Rate) error {

// 	query := `INSERT INTO ratings (user_id, product_id, rating, description)
// 	VALUES ($1, $2, $3, $4);`
// 	err := pd.DB.Raw(query, userID, productID, rating.Rating, rating.Description).Error
// 	return err

// }
func (pd *productDatabase) GetProductReviews(productID int) ([]response.Rating, error) {
	var Ratings = make([]response.Rating, 0)

	query := `SELECT r.id,r.rating,r.description,u.user_name FROM ratings r INNER JOIN users u ON u.id = r.user_id WHERE r.product_id =$1;`
	err := pd.DB.Raw(query, productID).Scan(&Ratings).Error
	return Ratings, err
}

func (pd *productDatabase) InsertProductRating(rating request.Rating) error {

	err := pd.DB.Create(&rating).Error
	return err

}

func (pd *productDatabase) SearchProducts(search string, startIndex, endIndex int) ([]response.Product, error) {
	var Products = make([]response.Product, 0)

	query := `SELECT *
	FROM products
	WHERE product_name ILIKE $1 OR
		  brand ILIKE $1 OFFSET  $2 FETCH NEXT $3 ROW ONLY ;`
	search = search + "%"

	err := pd.DB.Raw(query, search, startIndex, endIndex).Scan(&Products).Error
	return Products, err
}
func (pd *productDatabase) GetProductsByCategory(categoryID int, startIndex, endIndex int) ([]response.Product, error) {
	var Products = make([]response.Product, 0)

	query := `SELECT *
	FROM products
	WHERE category_id = $1 OFFSET  $2 FETCH NEXT $3 ROW ONLY ;`

	err := pd.DB.Raw(query, categoryID, startIndex, endIndex).Scan(&Products).Error
	return Products, err

}
