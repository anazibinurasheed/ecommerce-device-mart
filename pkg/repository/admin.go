package repository

import (
	"errors"
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

// from .env
// error not setting as admin
// not working properly
func (ud *adminDatabase) SaveAdminOnDatabase(admin request.SignUpData) (response.UserData, error) {
	query := `INSERT INTO users (user_name,  email, phone, password,created_at,is_admin) VALUES ($1,$2,$3,$4,$5,$6) RETURNING * ;`
	CreatedAt := time.Now()
	var AdminData response.UserData
	IsAdmin := true
	err := ud.DB.Raw(query, admin.UserName, admin.Email, admin.Phone, admin.Password, CreatedAt, IsAdmin).Scan(&AdminData).Error

	fmt.Println("REPOSITORY")
	fmt.Printf("%#v", AdminData)
	fmt.Println("")
	return AdminData, err
}

func (ad *adminDatabase) FindAdminLoginCredentials() (config.AdminCredentials, error) {
	var adminCredentials = config.GetAdminCredentials()
	if adminCredentials.AdminUsername == "" || adminCredentials.AdminPassword == "" {
		return adminCredentials, errors.New("Failed to find admin credentials")
	}

	return adminCredentials, nil
}

func (ad *adminDatabase) GetAllUserDataFromDatabase() ([]response.UserData, error) {
	var ListOfAllUsers = make([]response.UserData, 0)
	query := "SELECT Id,user_name ,email,phone ,is_blocked FROM users ORDER BY Id"
	err := ad.DB.Raw(query).Scan(&ListOfAllUsers).Error

	return ListOfAllUsers, err

}

func (ad *adminDatabase) BlockUserOnDatabase(id int) error {
	var BlockedUser response.UserData
	status := true
	query := "UPDATE Users SET Is_blocked =$1  WHERE Id =$2 RETURNING *"
	err := ad.DB.Raw(query, status, id).Scan(&BlockedUser).Error
	fmt.Println(BlockedUser)
	return err
}

func (ad *adminDatabase) UnBlockUserOnDatabase(id int) error {
	var BlockedUser response.UserData
	status := false
	query := "UPDATE Users SET Is_blocked =$1 WHERE id =$2 RETURNING *"
	err := ad.DB.Raw(query, status, id).Scan(&BlockedUser).Error
	fmt.Println(BlockedUser)
	return err

}
func (ad *adminDatabase) FindUserByNameFromDatabase(name string) ([]response.UserData, error) {
	var users []response.UserData
	query := "SELECT * FROM Users WHERE User_name ILIKE `%$1%`"
	err := ad.DB.Raw(query, name).Scan(&users).Error
	return users, err

}
