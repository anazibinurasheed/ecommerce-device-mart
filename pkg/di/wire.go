//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/anazibinurasheed/project-device-mart/pkg/api"
	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/anazibinurasheed/project-device-mart/pkg/db"
	"github.com/anazibinurasheed/project-device-mart/pkg/repository"
	"github.com/anazibinurasheed/project-device-mart/pkg/usecase"

	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {

	wire.Build(

		db.ConnectToDatabase,

		handler.NewAdminHandler,

		handler.NewUserHandler,

		handler.NewProductHandler,

		handler.NewCartHandler,

		handler.NewCommonHandler,

		handler.NewOrderHandler,

		handler.NewCouponHandler,

		handler.NewRefferalHandler,

		usecase.NewAdminUseCase,

		usecase.NewUserUseCase,

		usecase.NewProductUseCase,

		usecase.NewCommonUseCase,

		usecase.NewCartUseCase,

		usecase.NewOrderUseCase,

		usecase.NewCouponUseCase,

		usecase.NewRefferalUseCase,

		repository.NewAdminRepository,

		repository.NewUserRepository,

		repository.NewProductRepository,

		repository.NewCartRepository,

		repository.NewPaymentRepository,

		repository.NewOrderRepository,

		repository.NewCouponRepository,

		repository.NewRefferalRepository,

		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
