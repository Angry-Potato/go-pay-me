package db

import (
	"os"

	"github.com/jinzhu/gorm"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Initialise the db with the given models to automigrate
func Initialise(models ...interface{}) (*gorm.DB, error) {
	DB, err := Connect(
		os.Getenv("DATABASE_URL"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		return nil, err
	}
	DB = DB.AutoMigrate(models...)

	DB.LogMode(true)
	return DB, nil
}
