package repo

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type walletDatabase struct {
	DB *gorm.DB
}

func NewWalletRepository(DB *gorm.DB) interfaces.WalletRepository {
	return &walletDatabase{
		DB: DB,
	}
}

func (wd *walletDatabase) InsertIntoWallet(userID int, amount float32) (response.Wallet, error) {
	var InsertedRecord response.Wallet

	query := `INSERT INTO wallets (user_id,amount)VALUES($1,$2) RETURNING *;`
	err := wd.DB.Raw(query, userID, amount).Scan(&InsertedRecord).Error
	return InsertedRecord, err
}

// wishlist
func (wd *walletDatabase) AddToWishList(userID, productID int) error {

	query := `insert into wishlists (user_id, product_id) values($1, $2)`
	return wd.DB.Exec(query, userID, productID).Error
}

func (wd *walletDatabase) RemoveFromWishList(userID, productID int) error {

	query := `delete from wishlists where product_id = $1 and user_id = $2;`
	return wd.DB.Exec(query, productID, userID).Error
}

func (wd *walletDatabase) ShowWishListProducts(userID, page, count int) ([]response.Product, error) {

	products := []response.Product{}
	query := `select p.* from products p inner join wishlists w on w.product_id = p.id where w.user_id = $1 offset $2 fetch next $3 row only `
	err := wd.DB.Raw(query, userID, page, count).Scan(&products).Error
	return products, err
}
