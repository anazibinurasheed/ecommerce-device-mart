package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/anazibinurasheed/project-device-mart/pkg/config"
	domain "github.com/anazibinurasheed/project-device-mart/pkg/domain"
)

// for feature isolation
var dbInstance *gorm.DB

func ConnectToDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	db.Debug()

	if err := db.AutoMigrate(

		&domain.User{},
		&domain.Category{},
		&domain.Product{},
		&domain.Addresses{},
		&domain.State{},
		&domain.Cart{},
		&domain.PaymentMethod{},
		&domain.OrderLine{},
		&domain.OrderStatus{},
		&domain.Rating{},
		&domain.Coupon{},
		&domain.CouponTracking{},
		&domain.Wallet{},
		&domain.Referral{},
		&domain.WalletTransactionHistory{},
		&domain.Wishlist{},
	); err != nil {
		log.Fatal("Failed to connect with DB", err)
		return nil, err
	}
	dbInstance = db
	return db, dbErr

}

func GetDBInstance() *gorm.DB {
	return dbInstance
}
