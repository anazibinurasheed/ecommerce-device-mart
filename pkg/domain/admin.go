package domain

// this table will not migrated because it is created for follow the same pattern .
type Admin struct {
	UserName string `json:"username" gorm:"not null" binding:"required,min=3,max=25"`
	Password string `json:"password" gorm:"not null" binding:"required"`
}
