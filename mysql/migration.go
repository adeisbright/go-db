package mysql

import (
	"github.com/adeisbright/go-db/mysql/entities"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{},
		&entities.Post{},
	)
}
