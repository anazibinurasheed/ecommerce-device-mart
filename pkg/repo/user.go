package repo

import (
	"time"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

//for wire

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
}

func (ud *userDatabase) FindUserByPhone(phone int) (response.UserData, error) {
	var userData response.UserData
	query := `SELECT * FROM users WHERE  phone = $1`
	err := ud.DB.Raw(query, phone).Scan(&userData).Error

	return userData, err
}

func (ud *userDatabase) FindUserByEmail(email string) (response.UserData, error) {
	var UserData response.UserData
	query := `SELECT * FROM users WHERE  email = $1`
	err := ud.DB.Raw(query, email).Scan(&UserData).Error

	return UserData, err
}

func (ud *userDatabase) FindUserByID(id int) (response.UserData, error) {
	var UserData response.UserData
	query := `SELECT * FROM users WHERE  ID= $1`
	err := ud.DB.Raw(query, id).Scan(&UserData).Error

	return UserData, err
}

func (ud *userDatabase) CreateUser(user request.SignUpData) (response.UserData, error) {
	query := `INSERT INTO users (user_name,  email, phone, password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id,user_name,email,phone;`
	var userData response.UserData
	err := ud.DB.Raw(query, user.UserName, user.Email, user.Phone, user.Password, time.Now(), time.Now()).Scan(&userData).Error

	return userData, err
}

func (ud *userDatabase) ReadCategories() ([]response.Category, error) {
	var ListOfAllCategories = make([]response.Category, 0)
	query := `SELECT * FROM Categories WHERE Is_blocked = false  ORDER BY Category_Name;`
	err := ud.DB.Raw(query).Scan(&ListOfAllCategories).Error
	return ListOfAllCategories, err
}

func (ud *userDatabase) GetListOfStates() ([]response.States, error) {
	var ListOfStates = make([]response.States, 0)
	query := `SELECT * FROM States`
	err := ud.DB.Raw(query).Scan(&ListOfStates).Error
	return ListOfStates, err
}

func (ud *userDatabase) AddAddress(userID int, address request.Address) (response.Address, error) {
	var NewAddress response.Address

	query := `INSERT INTO Addresses 
	(name , phone_number,pincode,locality,address_line,district,state_id,landmark,alternative_phone,user_id)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING * ; `

	err := ud.DB.Raw(query, address.Name, address.PhoneNumber, address.Pincode, address.Locality,
		address.AddressLine, address.District, address.StateID, address.Landmark, address.AlternativePhone, userID).Scan(&NewAddress).Error

	return NewAddress, err

}

func (ud *userDatabase) FindDefaultAddress(userID int) (response.Address, error) {
	var DefaultAddress response.Address
	query := `SELECT * FROM Addresses WHERE is_default = true AND user_id = $1 FETCH FIRST 1 ROW ONLY ; `
	err := ud.DB.Raw(query, userID).Scan(&DefaultAddress).Error
	return DefaultAddress, err
}

func (ud *userDatabase) SetDefaultAddressStatus(status bool, addressId int, userId int) (response.Address, error) {

	var DefaultAddress response.Address
	query := `UPDATE Addresses SET is_default = $1 WHERE user_id = $2  AND id = $3 RETURNING * ;`
	err := ud.DB.Raw(query, status, userId, addressId).Scan(&DefaultAddress).Error

	return DefaultAddress, err
}

func (ud *userDatabase) GetAllUserAddresses(userID int) ([]response.Address, error) {

	var ListOfUserAddress = make([]response.Address, 0)
	query := `SELECT a.id,a.user_id,a.name  , a.phone_number , a.pincode ,a.locality , a.address_line,a.district , a.state_id , a.landmark ,a.alternative_phone,a.is_default ,s.name AS state FROM Addresses a  INNER JOIN states s ON a.state_id = s.id  WHERE a.user_id = $1 ORDER BY is_default ; `
	err := ud.DB.Raw(query, userID).Scan(&ListOfUserAddress).Error

	return ListOfUserAddress, err
}

func (ud *userDatabase) UpdateAddress(address request.Address, addressID int, userID int) (response.Address, error) {
	var UpdatedAddress response.Address

	query := `UPDATE addresses SET name = $1 ,phone_number = $2 , pincode = $3 ,locality = $4 , address_line = $5 ,district = $6 , state_id = $7 , landmark = $8 , alternative_phone = $9 WHERE id = $10 AND user_id = $11 RETURNING *;`

	err := ud.DB.Raw(query, address.Name, address.PhoneNumber, address.Pincode, address.Locality, address.AddressLine, address.District, address.StateID, address.Landmark, address.AddressLine, addressID, userID).Scan(&UpdatedAddress).Error

	return UpdatedAddress, err
}

func (ud *userDatabase) DeleteAddress(addressID int) (response.Address, error) {
	var DeletedAddress response.Address
	query := `DELETE FROM Addresses WHERE Id = $1 RETURNING * ; `
	err := ud.DB.Raw(query, addressID).Scan(&DeletedAddress).Error
	return DeletedAddress, err
}

func (ud *userDatabase) ChangePassword(userId int, newPassword string) error {
	var user response.UserData
	query := `UPDATE users SET password = $1 WHERE Id = $2  ; `
	err := ud.DB.Raw(query, newPassword, userId).Scan(&user).Error
	return err
}

func (ud *userDatabase) FindAddressByID(addressID int) (response.Address, error) {

	var UserAddress response.Address
	query := `SELECT * FROM addresses WHERE id = $1 ;  `

	err := ud.DB.Raw(query, addressID).Scan(&UserAddress).Error

	return UserAddress, err
}

func (ud *userDatabase) FindUserAddress(userID int) (response.Address, error) {

	var UserAddress response.Address
	query := `SELECT * FROM addresses WHERE User_id = $1 FETCH FIRST 1 ROW ONLY ;  `

	err := ud.DB.Raw(query, userID).Scan(&UserAddress).Error

	return UserAddress, err
}

func (ud *userDatabase) UpdateUserName(name string, userID int) (response.UserData, error) {
	var user response.UserData
	query := `UPDATE users SET user_name = $1 WHERE Id = $2  RETURNING * ; `

	err := ud.DB.Raw(query, name, userID).Scan(&user).Error
	return user, err
}
