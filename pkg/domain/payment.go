package domain

// changed int to uint
type PaymentMethod struct {
	ID         uint   `gorm:"primaryKey;AutoIncrement;unique"`
	MethodName string `gorm:"not null;unique"`
}
