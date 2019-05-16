package testhelpers

import (
	"fmt"
	"os"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/db"
	"github.com/Angry-Potato/go-pay-me/implementation/schema"
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
	DB, err := db.Initialise(&schema.Payment{}, &schema.PaymentAttributes{})
	assert.Nil(t, err)
	assert.NotNil(t, DB)
	DB.LogMode(false)
	return DB
}
