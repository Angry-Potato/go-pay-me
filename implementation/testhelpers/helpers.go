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
	DB, err := db.Initialise(&schema.Payment{}, &schema.PaymentAttributes{}, &schema.Party{}, &schema.Party{}, &schema.CurrencyExchange{}, &schema.Charges{}, &schema.Money{})

	//this is bad, where should I do this?
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("internal_payment_id", "payments(id)", "CASCADE", "CASCADE")
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("beneficiary_party_id", "parties(id)", "SET NULL", "CASCADE")
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("debtor_party_id", "parties(id)", "SET NULL", "CASCADE")
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("sponsor_party_id", "parties(id)", "SET NULL", "CASCADE")
	DB.Model(&schema.CurrencyExchange{}).AddForeignKey("payment_attributes_id", "payment_attributes(id)", "CASCADE", "CASCADE")
	DB.Model(&schema.Charges{}).AddForeignKey("payment_attributes_id", "payment_attributes(id)", "CASCADE", "CASCADE")
	DB.Model(&schema.Money{}).AddForeignKey("charges_id", "charges(id)", "CASCADE", "CASCADE")

	assert.Nil(t, err)
	assert.NotNil(t, DB)
	DB.LogMode(false)
	return DB
}
