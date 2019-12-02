package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"schoolManagementAPI/config"
	u "schoolManagementAPI/utils"
	"strings"
)


type Token struct {
	UserId uint
	jwt.StandardClaims
}
type Student struct {

	gorm.Model   `gorm:"AUTO INCREMENT"`
	FirstName string `gorm:"type:varchar(130);not null;" json:"fisrt_name"`
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Age       string    `gorm:"not null" json:"age"`
	LastName  string `gorm:"type:varchar(130);not null" json:"last_name"`
	Phone  string `gorm:"type:varchar(100);not null" json:"phone"`
	DepartmentId  uint `gorm:"not null" json:"department_id"`
	Department Department
}

func (student *Student) validateSignUp() (map[string]interface{}, bool)  {

	if student.FirstName == ""{
		return u.Message(false, "First name is required"), false
	}
	if student.LastName == ""{
		return u.Message(false, "Last name is required"), false
	}
	if student.Phone == ""{
		return u.Message(false, "Phone number is required"), false
	}

	if !strings.Contains(student.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(student.Password) < 6 {
		return u.Message(false, "invalid Password!"), false
	}



	err := config.GetDB().Table("student").Find(&Student{}).Where("email=?",student.Email).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if err!= nil{
		return u.Message(false,"Email address already in use by another user."),false
	}

	return u.Message(false, "Requirement passed"), true


}









