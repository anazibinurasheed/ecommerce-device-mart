package repository

import (
	"fmt"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{
		DB: DB,
	}

}

func (ud *adminDatabase) CreateAdmin(admin request.SignUpData) (response.UserData, error) {
	query := `INSERT INTO users (user_name,  email, phone, password,created_at,is_admin) VALUES ($1,$2,$3,$4,$5,$6) RETURNING * ;`
	CreatedAt := time.Now()
	var AdminData response.UserData
	isAdmin := true
	err := ud.DB.Raw(query, admin.UserName, admin.Email, admin.Phone, admin.Password, CreatedAt, isAdmin).Scan(&AdminData).Error

	return AdminData, err
}

func (ad *adminDatabase) FindAdminCredentials() (config.AdminCredentials, error) {
	var adminCredentials = config.GetAdminCredentials()
	if adminCredentials.AdminUsername == "" || adminCredentials.AdminPassword == "" {
		return adminCredentials, fmt.Errorf("Failed to fetch admin credentials")
	}

	return adminCredentials, nil
}

func (ad *adminDatabase) FetchAllUserData() ([]response.UserData, error) {
	var ListOfAllUsers = make([]response.UserData, 0)
	query := "SELECT Id,user_name ,email,phone ,is_blocked FROM users ORDER BY Id"
	err := ad.DB.Raw(query).Scan(&ListOfAllUsers).Error
	return ListOfAllUsers, err
}

func (ad *adminDatabase) BlockUserByID(userID int) error {
	var BlockedUser response.UserData
	status := true
	query := "UPDATE Users SET Is_blocked =$1  WHERE Id =$2 RETURNING *"
	err := ad.DB.Raw(query, status, userID).Scan(&BlockedUser).Error
	fmt.Println(BlockedUser)
	return err
}

func (ad *adminDatabase) UnblockUserByID(userID int) error {
	var BlockedUser response.UserData
	status := false
	query := "UPDATE Users SET Is_blocked =$1 WHERE id =$2 RETURNING *"
	err := ad.DB.Raw(query, status, userID).Scan(&BlockedUser).Error
	fmt.Println(BlockedUser)
	return err
}

func (ad *adminDatabase) FindUsersByName(name string) ([]response.UserData, error) {
	var users []response.UserData
	query := "SELECT * FROM Users WHERE User_name ILIKE `%$1%`"
	err := ad.DB.Raw(query, name).Scan(&users).Error
	return users, err
}
