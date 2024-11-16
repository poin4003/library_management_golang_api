package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"library_management/global"
	"library_management/models"
)

func InitPostgresql() {
	dsn := "host=localhost user=postgres password=Nvidia_geforce940mx dbname=social_media port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to PostgreSQL")

	if err := DBMigrator(db); err != nil {
		fmt.Println(err)
	}

	global.PDB = db
}

func DBMigrator(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Book{},
	)
	return err
}
