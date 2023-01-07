package models

import (
	"gorm.io/gorm"
)

type Classroom struct {
	gorm.Model
	Name      string    `json:"name"`
	Student   []Student `gorm:"foreighkey:ClassroomID" json:"stdents"`
	Teacher   Teacher   `json:"teacher_name"`
	TeacherID uint
}
