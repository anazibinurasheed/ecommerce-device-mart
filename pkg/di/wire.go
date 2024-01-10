// //go:build wireinject
// // +build wireinject

package di

// func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {

// 	wire.Build(

// 		db.ConnectToDatabase,

// 		handler.NewAdminHandler,

// 		handler.NewUserHandler,

// 		handler.NewProductHandler,

// 		handler.NewCartHandler,

// 		handler.NewCommonHandler,

// 		handler.NewOrderHandler,

// 		handler.NewCouponHandler,

// 		handler.NewReferralHandler,

// 		usecase.NewAdminUseCase,

// 		usecase.NewUserUseCase,

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

// 		api.NewServerHTTP)

// 	return &api.ServerHTTP{}, nil
// }
