package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type WalletRepository interface {
	InsertIntoWallet(userID int, amount float32) (response.Wallet, error)
}
