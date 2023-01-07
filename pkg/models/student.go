package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name        string    `json:"name"`
	Classroom   Classroom `json:"classroom"`
	ClassroomID uint
}
