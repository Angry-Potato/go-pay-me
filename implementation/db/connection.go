package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect to Postgres database
func Connect(url, host, port, user, password, database string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connectionString(url, host, port, user, password, database))

	return db, err
}

func connectionString(url, host, port, user, password, database string) string {
	if url != "" {
		return url
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
}
