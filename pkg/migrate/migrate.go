package migrate

import (
	"github.com/Emotions-s/go-classroom-management-basic/pkg/config"
	"github.com/Emotions-s/go-classroom-management-basic/pkg/models"
)

func Migrate() {
	config.Loadenviranment()
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&models.Classroom{}, &models.Student{}, &models.Teacher{})
}
