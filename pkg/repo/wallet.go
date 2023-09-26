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

//not using this file
