package testhelpers

import (
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// APIAddress returns http address of the API under test, populated by env vars API_HOST and API_PORT
func APIAddress(t *testing.T) string {
	t.Helper()

	host, port := os.Getenv("API_HOST"), os.Getenv("API_PORT")
	return fmt.Sprintf("http://%s:%s", host, port)
}

// DBConnection gets a DB connection
func DBConnection(t *testing.T, models ...interface{}) *gorm.DB {
	t.Helper()
	host, port, user, password, database := os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	DB, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database))
	assert.Nil(t, err)
	assert.NotNil(t, DB)

	return DB.AutoMigrate(models...)
}
