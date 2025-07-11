package dao

import (
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) {
	db.AutoMigrate(
		SystemUser{},
	)
}
