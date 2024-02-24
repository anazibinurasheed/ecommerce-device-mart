// // //go:build wireinject
// // // +build wireinject

package di

// import (
// 	"github.com/anazibinurasheed/project-device-mart/pkg/api"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/config"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/db"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/repo"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/usecase"
// 	"github.com/google/wire"
// )

// func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {

// 	wire.Build(

// 		db.ConnectToDatabase,

// 		middleware.NewAuthMiddleware,

// 		handler.NewAdminHandler,

// 		handler.NewAuthHandler,

// 		handler.NewUserHandler,

// 		handler.NewProductHandler,

// 		handler.NewCartHandler,

// 		handler.NewOrderHandler,

// 		handler.NewCouponHandler,

// 		handler.NewReferralHandler,

// 		handler.NewWalletHandler,

// 		handler.NewRazorpayHandler,

// 		usecase.NewAdminUseCase,

// 		usecase.NewUserUseCase,

// 		usecase.NewRazorpayUseCase,

// 		usecase.NewWalletUseCase,

// 		usecase.NewProductUseCase,

// 		usecase.NewCommonUseCase,

// 		usecase.NewCartUseCase,

// 		usecase.NewOrderUseCase,

// 		usecase.NewCouponUseCase,

// 		usecase.NewReferralUseCase,

// 		repo.NewAdminRepository,

// 		repo.NewUserRepository,

// 		repo.NewProductRepository,

// 		repo.NewCartRepository,

// 		repo.NewPaymentRepository,

// 		repo.NewOrderRepository,

// 		repo.NewCouponRepository,

// 		repo.NewReferralRepository,

// 		repo.NewWalletRepository,

// 		api.NewServerHTTP)

// 	return &api.ServerHTTP{}, nil
// }
