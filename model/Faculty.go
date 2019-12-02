package model

import "github.com/jinzhu/gorm"

type Faculty struct {
	gorm.Model    `gorm:"AUTO INCREMENT"`
	Name string  `gorm:"type:VARCHAR(100); NOT NULL;" json:"faculty_name"`
}