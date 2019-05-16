package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validParty() *Party {
	return &Party{
		AccountName:       "My account",
		AccountNumber:     "12345",
		AccountNumberCode: "93847",
		AccountType:       5,
		Address:           "123 lane",
		BankID:            "best bank",
		BankIDCode:        "12AFR",
		Name:              "Mr Man",
	}
}

func Test_validateParty_Returns_Error_When_AccountNumber_Is_Empty(t *testing.T) {
	invalidParty := validParty()
	invalidParty.AccountNumber = ""
	errs := validateParty(invalidParty)
	assert.NotEmpty(t, errs)
}

func Test_validateParty_Returns_Error_When_BankID_Is_Empty(t *testing.T) {
	invalidParty := validParty()
	invalidParty.BankID = ""
	errs := validateParty(invalidParty)
	assert.NotEmpty(t, errs)
}

func Test_validateParty_Returns_Error_When_BankIDCode_Is_Empty(t *testing.T) {
	invalidParty := validParty()
	invalidParty.BankIDCode = ""
	errs := validateParty(invalidParty)
	assert.NotEmpty(t, errs)
}
