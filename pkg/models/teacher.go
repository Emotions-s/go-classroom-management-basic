package models

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name      string      `json:"name"`
	Classroom []Classroom `gorm:"foreighkey:TeacherID" json:"classroom"`
}
