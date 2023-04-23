package migrations

import (
	"github.com/dg943/MyProject/backend/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	models.CreateUserTable(db)
}
