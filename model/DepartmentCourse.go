package model

import "github.com/jinzhu/gorm"

type DepartmentCourse struct {
	gorm.Model
	DepartmentID uint `gorm:"not null" json:"code" `
	CourseID uint
	Department Department
	Course Course
}
