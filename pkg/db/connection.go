package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/anazibinurasheed/project-device-mart/pkg/config"
	domain "github.com/anazibinurasheed/project-device-mart/pkg/domain"
)

//this function is only for intialize the database and do migrate tables
// it is also used in the di / wire.go

func ConnectToDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	db.Set("gorm.singular_table_names", true)

	db.Debug() //to log the query

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
	); err != nil {
		log.Fatal("FAILED TO CONNECT WITH DATABASE ", err)
		return nil, err
	}

	return db, dbErr

}
