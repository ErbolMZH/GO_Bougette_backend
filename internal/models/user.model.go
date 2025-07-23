package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	FirstName *string `gorm:"type:varchar(255)"`
	LastName  *string `gorm:"type:varchar(255)"`
	Email     string  `gorm:"type:varchar(255)"`
	Gender    *string `gorm:"type:varchar(50)"`
	Password  string  `gorm:"type:varchar(255) not null"`
}

func (receiver UserModel) TableName() string {
	return "users"
}
