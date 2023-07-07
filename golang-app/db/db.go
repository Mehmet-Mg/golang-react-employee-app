package db

import (
	"example/employee-app/config"
	"example/employee-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.ConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(models.Employee{}); err != nil {
		panic("failed to migrate database")
	}
}
