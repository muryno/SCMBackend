package model

import "github.com/jinzhu/gorm"

type Student struct {

	gorm.Model
	FirstName string `gorm:"type:varchar(130);not null;" json:"fisrt_name"`
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Age       string    `gorm:"not null" json:"age"`
	LastName  string `gorm:"type:varchar(130);not null" json:"last_name"`
	Phone  string `gorm:"type:varchar(100);not null" json:"phone"`
}








