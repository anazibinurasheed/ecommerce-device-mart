// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/api"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"
	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/anazibinurasheed/project-device-mart/pkg/db"
	"github.com/anazibinurasheed/project-device-mart/pkg/repo"
	"github.com/anazibinurasheed/project-device-mart/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	gormDB, err := db.ConnectToDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repo.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	adminRepository := repo.NewAdminRepository(gormDB)
	adminRepository.SetupDB()
	adminUseCase := usecase.NewAdminUseCase(adminRepository, userRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	productRepository := repo.NewProductRepository(gormDB)
	orderRepository := repo.NewOrderRepository(gormDB)
	productUseCase := usecase.NewProductUseCase(productRepository, orderRepository)
	productHandler := handler.NewProductHandler(productUseCase)
	authUseCase := usecase.NewCommonUseCase(userRepository, adminRepository)
	authHandler := handler.NewAuthHandler(authUseCase)
	cartRepository := repo.NewCartRepository(gormDB)
	couponRepository := repo.NewCouponRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, couponRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)
	paymentRepository := repo.NewPaymentRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(userRepository, cartUseCase, paymentRepository, orderRepository, couponRepository, productRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)
	couponUseCase := usecase.NewCouponUseCase(couponRepository)
	couponHandler := handler.NewCouponHandler(couponUseCase)
	referralRepository := repo.NewReferralRepository(gormDB)
	referralUseCase := usecase.NewReferralUseCase(referralRepository, orderRepository)
	referralHandler := handler.NewReferralHandler(referralUseCase)
	authMiddleware := middleware.NewAuthMiddleware(userUseCase)
	walletRepository := repo.NewWalletRepository(gormDB)
	walletUseCase := usecase.NewWalletUseCase(walletRepository, orderRepository, cartUseCase)
	walletHandler := handler.NewWalletHandler(walletUseCase, orderUseCase)
	razorpayUseCase := usecase.NewRazorpayUseCase(paymentRepository, cartUseCase, userRepository)
	razorpayHandler := handler.NewRazorpayHandler(razorpayUseCase, orderUseCase)
	serverHTTP := api.NewServerHTTP(userHandler, adminHandler, productHandler, authHandler, cartHandler, orderHandler, couponHandler, referralHandler, authMiddleware, walletHandler, razorpayHandler)
	return serverHTTP, nil
}
