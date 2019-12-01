package model

import "github.com/jinzhu/gorm"

type RegisteredCourse struct {
	gorm.Model
	CourseID uint `gorm:"not null" json:"course_id"`
	StudentID   uint `gorm:"not null" json:"student_id"`
	Course   Course
	Student     Student
}
