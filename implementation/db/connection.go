package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect to Postgres database
func Connect(host, port, user, password, database string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connectionString(host, port, user, password, database))

	return db, err
}

func connectionString(host, port, user, password, database string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
}
