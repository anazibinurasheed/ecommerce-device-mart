package repo

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartDatabase{
		DB: DB,
	}
}

func (cd *cartDatabase) AddToCart(userID int, ProductID int) (response.Cart, error) {
	var CartItem response.Cart
	qty := 1
	query := `INSERT INTO carts (user_id,product_id,qty)VALUES($1,$2,$3) RETURNING *;`
	err := cd.DB.Raw(query, userID, ProductID, qty).Scan(&CartItem).Error

	return CartItem, err

}

func (cd *cartDatabase) ViewCart(userID int) ([]response.Cart, error) {
	var CartItem = make([]response.Cart, 0)

	query := `SELECT c.id , c.product_id,c.qty,p.product_name , p.brand,p.price FROM carts c INNER JOIN products p ON c.product_id = p.id WHERE c.user_id = $1 `
	err := cd.DB.Raw(query, userID).Scan(&CartItem).Error

	return CartItem, err

}

func (cd *cartDatabase) RemoveFromCart(userID int, productID int) (response.Cart, error) {
	var CartItem response.Cart

	query := `DELETE FROM carts WHERE user_id = $1 AND product_id =$2 RETURNING * ; `
	err := cd.DB.Raw(query, userID, productID).Scan(&CartItem).Error

	return CartItem, err

}

func (cd *cartDatabase) IncrementQuantity(qty int, userID int, productID int) (response.Cart, error) {
	var CartItem response.Cart

	query := `UPDATE carts SET qty = $1 WHERE user_id = $2 AND product_id =$3 RETURNING * ; `
	err := cd.DB.Raw(query, qty, userID, productID).Scan(&CartItem).Error

	return CartItem, err

}

func (cd *cartDatabase) DecrementQuantity(qty int, userID int, productID int) (response.Cart, error) {
	var CartItem response.Cart

	query := `UPDATE carts SET qty = $1 WHERE user_id = $2 AND product_id =$3 RETURNING * ; `
	err := cd.DB.Raw(query, qty, userID, productID).Scan(&CartItem).Error

	return CartItem, err

}
func (cd *cartDatabase) GetCartItem(userID int, productID int) (response.Cart, error) {
	var CartItem response.Cart

	query := `SELECT * FROM carts WHERE user_id = $1 AND product_id = $2 ; `
	err := cd.DB.Raw(query, userID, productID).Scan(&CartItem).Error

	return CartItem, err

}
func (cd *cartDatabase) DeleteCart(userID int) (response.Cart, error) {
	var DeletedCart response.Cart

	query := `DELETE FROM carts WHERE user_id = $1 RETURNING *; `
	err := cd.DB.Raw(query, userID).Scan(&DeletedCart).Error

	return DeletedCart, err

}
