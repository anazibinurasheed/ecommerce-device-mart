package usecase

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
)

type walletUseCase struct {
	walletRepo interfaces.WalletRepository
}

func NewWalletUseCase(useCase interfaces.WalletRepository) services.WalletUseCase {
	return &walletUseCase{
		walletRepo: useCase,
	}
}
