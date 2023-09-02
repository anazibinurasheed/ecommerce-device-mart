package repository

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type paymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentDatabase{
		DB: DB,
	}
}

func (pd *paymentDatabase) GetPaymentMethods() ([]response.PaymentMethod, error) {
	var PaymentMethods = make([]response.PaymentMethod, 0)
	query := `SELECT * FROM payment_methods;`
	err := pd.DB.Raw(query).Scan(&PaymentMethods).Error

	return PaymentMethods, err
}

func (pd *paymentDatabase) GetPaymentMethodCodId() (int, error) {
	var codID int
	query := `SELECT id FROM  payment_methods WHERE method_name = cash on delivery ;`
	err := pd.DB.Raw(query).Scan(&codID).Error
	return codID, err
}

func (pd *paymentDatabase) GetPaymentMethodRazorpayId() (int, error) {
	var razorpayID int
	query := `SELECT id FROM  payment_methods WHERE method_name = online payment ;`
	err := pd.DB.Raw(query).Scan(&razorpayID).Error
	return razorpayID, err

}

func (pd *paymentDatabase) FindPaymentMethodById(methodID int) (response.PaymentMethod, error) {
	var PaymentMethod response.PaymentMethod
	query := `SELECT * FROM  payment_methods WHERE id = $1 ;`
	err := pd.DB.Raw(query, methodID).Scan(&PaymentMethod).Error
	return PaymentMethod, err

}
