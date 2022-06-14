package migrations

import (
	"github.com/luccasalves/apigo-sample/models"
	"gorm.io/gorm"
)

func RunMigs(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
