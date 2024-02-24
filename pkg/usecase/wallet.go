package usecase

import (
	"fmt"
	"time"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
)

type walletUseCase struct {
	walletRepo  interfaces.WalletRepository
	orderRepo   interfaces.OrderRepository
	cartUseCase services.CartUseCase
}

func NewWalletUseCase(walletRepo interfaces.WalletRepository,
	orderRepo interfaces.OrderRepository,
	cartUseCase services.CartUseCase) services.WalletUseCase {
	return &walletUseCase{
		walletRepo:  walletRepo,
		orderRepo:   orderRepo,
		cartUseCase: cartUseCase,
	}
}

func (ou *walletUseCase) GetWalletHistory(userID int) ([]response.WalletTransactionHistory, error) {
	walletHistory, err := ou.orderRepo.GetWalletHistoryByUserID(userID)
	if err != nil {
		return walletHistory, err
	}

	return walletHistory, nil
}

func (ou *walletUseCase) GetUserWallet(userID int) (response.Wallet, error) {
	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
	if err != nil {
		return response.Wallet{}, fmt.Errorf("Failed to get user wallet :%s", err)
	}
	if wallet.ID == 0 {
		return response.Wallet{}, ErrNoWallet
	}
	return wallet, nil
}

func (ou *walletUseCase) CreateUserWallet(userID int) error {

	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
	if err != nil {
		return fmt.Errorf("Failed to check user wallet :%s", err)
	}

	if wallet.ID != 0 {
		return fmt.Errorf("User already have wallet")
	}

	newWallet, err := ou.orderRepo.InitializeNewUserWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to initialize wallet for user %d : %s", userID, err)
	}
	if newWallet.ID == 0 {
		return fmt.Errorf("Failed to verify new wallet for user id  %d", userID)
	}
	return nil
}

func (ou *walletUseCase) ValidateWalletPayment(userID int) error {
	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	}
	if wallet.ID == 0 {
		return fmt.Errorf("User don't have a wallet ")
	}

	userCart, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return fmt.Errorf("Failed to fetch user cart : %s", err)
	}
	if userCart.Total > wallet.Amount {
		return fmt.Errorf("Insufficient balance")
	}
	return nil
}

func (ou *walletUseCase) UpdateWalletHistory(userID int, amount float32, transactionType string) error {

	walletHistory, err := ou.orderRepo.UpdateWalletTransactionHistory(
		request.WalletTransactionHistory{
			TransactionTime: time.Now(),
			UserID:          userID,
			Amount:          amount,
			TransactionType: transactionType,
		},
	)

	if err != nil || walletHistory.ID == 0 {
		return func() error {
			if err != nil {
				return err
			}
			return fmt.Errorf("Failed to verify the updated history")
		}()

	}
	return nil
}

func (ou *walletUseCase) UpdateWallet(userID int, amount float32, transactionType string) error {
	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
	if err != nil {
		return err
	}

	var updtAmount float32

	if transactionType == credit {
		updtAmount = (wallet.Amount + amount)

	}

	if transactionType == debit {
		updtAmount = (wallet.Amount - amount)

	}

	wallet, err = ou.orderRepo.UpdateUserWalletBalance(userID, updtAmount)
	if err != nil {
		return err
	}
	err = ou.UpdateWalletHistory(userID, amount, transactionType)
	if err != nil {
		return err
	}
	return nil
}
