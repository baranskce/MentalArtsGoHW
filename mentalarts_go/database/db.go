package database

import (
	"mentalarts_go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	dbURL := "postgres://root:root@localhost:5432/test_db"

	Db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		return nil, err
	}
	Db.AutoMigrate(&models.Musician{}, &models.Album{})

	return Db, nil
}
