package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type WalletUseCase interface {
	GetWalletHistory(userID int) ([]response.WalletTransactionHistory, error)
	GetUserWallet(userID int) (response.Wallet, error)
	CreateUserWallet(userID int) error
	ValidateWalletPayment(userID int) error
	UpdateWalletHistory(userID int, amount float32, transactionType string) error
	UpdateWallet(userID int, amount float32, transactionType string) error
}
