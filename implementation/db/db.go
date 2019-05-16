package db

import (
	"os"

	"github.com/Angry-Potato/go-pay-me/implementation/schema"
	"github.com/jinzhu/gorm"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("internal_payment_id", "payments(id)", "CASCADE", "CASCADE")
	DB.LogMode(true)
	return DB, nil
}
