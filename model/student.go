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

type Faculty struct {
	gorm.Model
     Name string  `gorm:"type:VARCHAR(100); NOT NULL;" json:"faculty_name"`
}

type Department struct {
	gorm.Model
	Name string  `gorm:"type:VARCHAR(100); NOT NULL;" json:"department_name"`
}

type Course struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(100); NOT NULL;" json:"department_name"`
	Code string `gorm:"not null" json:"code"`
}

type RegisteredCourse struct {
	gorm.Model
	CourseID uint `gorm:"not null" json:"course_id"`
	StudentID   uint `gorm:"not null" json:"student_id"`
	Course   Course
	Student     Student
}

type DepartmentCourse struct {
	gorm.Model
	DepartmentID uint `gorm:"not null" json:"code" `
	CourseID uint
	Department Department
	Course Course
}