package repo

import (
	"fmt"

	"github.com/anazibinurasheed/project-device-mart/pkg/domain"
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
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
	var result response.Category
	query := `INSERT INTO Categories (Category_Name) VALUES($1) RETURNING *;`
	err := pd.DB.Raw(query, category.CategoryName).Scan(&result).Error

	return result, err
}

func (pd *productDatabase) ReadCategory(startIndex int, endIndex int) ([]response.Category, error) {
	var ListOfAllCategories = make([]response.Category, 0)

	query := `SELECT * FROM Categories ORDER BY category_name OFFSET $1 FETCH NEXT $2 ROW ONLY ;`

	err := pd.DB.Raw(query, startIndex, endIndex).Scan(&ListOfAllCategories).Error
	fmt.Println(ListOfAllCategories)
	return ListOfAllCategories, err
}

func (pd *productDatabase) UpdateCategory(categoryID int, category request.Category) error {

	query := `UPDATE Categories SET Category_Name = $1 WHERE ID = $2 ;`

	err := pd.DB.Exec(query, category.CategoryName, categoryID).Error

	return err
}

func (pd *productDatabase) BlockCategoryByID(categoryID int) error {
	query := `Update Categories SET Is_blocked = $1 WHERE ID = $2`
	err := pd.DB.Exec(query, true, categoryID).Error

	return err
}

func (pd *productDatabase) UnBlockCategoryByID(categoryID int) error {
	query := `Update Categories SET Is_blocked = $1 WHERE ID = $2`
	err := pd.DB.Exec(query, false, categoryID).Error
	return err

}

func (pd *productDatabase) FindCategoryByName(name string) (response.Category, error) {
	var ResultOfFinding response.Category
	query := `SELECT * FROM Categories WHERE Category_Name = $1`
	err := pd.DB.Raw(query, name).Scan(&ResultOfFinding).Error
	return ResultOfFinding, err

}

func (pd *productDatabase) FindCategoryByID(categoryID int) (response.Category, error) {
	var ResultOfFinding response.Category
	query := `SELECT * FROM Categories WHERE Id = $1 `
	err := pd.DB.Raw(query, categoryID).Scan(&ResultOfFinding).Error
	return ResultOfFinding, err
}

func (pd *productDatabase) CreateProduct(product request.Product) (response.Product, error) {
	var result response.Product
	query := `INSERT INTO Products (Category_ID,Product_Name,Price,Product_Description, Brand,Sku,is_blocked) Values($1,$2,$3,$4,$5,$6,$7) returning *;`
	err := pd.DB.Raw(query, product.CategoryID, product.ProductName, product.Price, product.ProductDescription, product.Brand, product.SKU, product.IsBlocked).Scan(&result).Error
	return result, err
}

func (pd *productDatabase) ViewAllProductsToAdmin(startIndex, endIndex int) ([]response.Product, error) {
	ListOfAllProducts := []response.Product{}
	query := `SELECT * FROM Products OFFSET $1 FETCH NEXT $2 ROW ONLY ;`
	err := pd.DB.Raw(query, startIndex, endIndex).Scan(&ListOfAllProducts).Error
	return ListOfAllProducts, err
}

func (pd *productDatabase) UpdateProduct(productID int, updations request.UpdateProduct) error {
	query := `Update Products SET Category_ID = $1 ,Product_Name = $2 ,Product_Description = $3 , Price = $4  WHERE ID = $5`
	err := pd.DB.Exec(query, updations.CategoryID, updations.ProductName, updations.ProductDescription, updations.Price, productID).Error
	return err
}

func (pd *productDatabase) BlockProduct(productID int) error {
	status := true
	query := `UPDATE Products SET Is_Blocked = $1 WHERE ID = $2;`
	err := pd.DB.Exec(query, status, productID).Error
	return err
}

func (pd *productDatabase) UnblockProduct(productID int) error {
	status := false
	query := `UPDATE Products SET is_blocked = $1 WHERE id = $2 RETURNING *;`
	err := pd.DB.Exec(query, status, productID).Error
	return err
}

func (pd *productDatabase) FindProductByName(productName string) (response.Product, error) {
	var product response.Product
	query := `SELECT * FROM Products WHERE Product_name = $1  FETCH FIRST 1 ROW ONLY`
	err := pd.DB.Raw(query, productName).Scan(&product).Error
	return product, err
}

func (pd *productDatabase) ViewIndividualProduct(userID, productID int) (response.Product, error) {
	var product response.Product
	query := `SELECT p.*, EXISTS (SELECT 1 FROM wishlists WHERE user_id = $1 AND product_id = $2) AS is_wishlisted
	FROM products p WHERE P.id = $2 FETCH FIRST 1 ROW ONLY`
	err := pd.DB.Raw(query, userID,productID).Scan(&product).Error
	return product, err
}

func (pd *productDatabase) ViewAllProductsToUser(userID, startIndex, endIndex int) ([]response.Product, error) {
	ListOfAllProducts := []response.Product{}
	query := `SELECT p.*, EXISTS (SELECT 1 FROM wishlists WHERE user_id = $1 AND product_id = p.id) AS is_wishlisted
	FROM products p OFFSET $2 FETCH NEXT $3 ROW ONLY`
	err := pd.DB.Raw(query, userID,startIndex, endIndex).Scan(&ListOfAllProducts).Error
	return ListOfAllProducts, err
}

func (pd *productDatabase) FindProductByID(productID int) (response.Product, error) {
	var Product response.Product
	query := "SELECT * FROM Products WHERE Id = $1  FETCH FIRST 1 ROW ONLY"
	err := pd.DB.Raw(query, productID).Scan(&Product).Error
	return Product, err
}

func (pd *productDatabase) FindUserRatingOnProduct(userID, productID int) (response.Rating, error) {
	var Rating response.Rating

	query := `SELECT r.id,r.rating,r.description,u.user_name FROM ratings r 
	INNER JOIN users u ON r.user_id = u.id WHERE r.user_id = $1 AND r.product_id = $2;`
	err := pd.DB.Raw(query, userID, productID).Scan(&Rating).Error
	return Rating, err
}

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

func (pd *productDatabase) InsertCategoryIMG(urls interface{}, categoryID int) error {

	images := domain.NewJsonB()
	images["urls"] = urls
	query := `update categories set images = $1 where id = $2;`
	return pd.DB.Exec(query, images, categoryID).Error
}

func (pd *productDatabase) InsertProductIMG(urls interface{}, productID int) error {

	images := domain.NewJsonB()
	images["urls"] = urls
	query := `update products set images = $1 where id = $2;`
	return pd.DB.Exec(query, images, productID).Error
}

// wishlist
func (pd *productDatabase) AddToWishList(userID, productID int) error {

	query := `insert into wishlists (user_id, product_id) values($1, $2)`
	return pd.DB.Exec(query, userID, productID).Error
}

func (pd *productDatabase) RemoveFromWishList(userID, productID int) error {

	query := `delete from wishlists where product_id = $1 and user_id = $2;`
	return pd.DB.Exec(query, productID, userID).Error
}

func (pd *productDatabase) ShowWishListProducts(userID, page, count int) ([]response.Product, error) {

	products := []response.Product{}
	query := `select p.* from products p inner join wishlists w on w.product_id = p.id where w.user_id = $1 offset $2 fetch next $3 row only `
	err := pd.DB.Raw(query, userID, page, count).Scan(&products).Error
	return products, err
}
