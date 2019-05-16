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

func Test_IsSameParty_Returns_True_When_Parties_Are_Equal(t *testing.T) {
	partyA, partyB := validParty(), validParty()
	assert.Equal(t, partyA, partyB)
	assert.True(t, IsSameParty(partyA, partyB))
	partyA.AccountName = "Her account"
	partyA.AccountNumberCode = "45678"
	partyA.AccountType = 90
	partyA.Address = "321 street"
	partyA.Name = "Miss Lady"
	assert.True(t, IsSameParty(partyA, partyB))
}

func Test_IsSameParty_Returns_False_When_Parties_Are_Not_Equal(t *testing.T) {
	partyA, partyB := validParty(), validParty()
	assert.Equal(t, partyA, partyB)
	partyA.AccountNumber = "some new account"
	assert.False(t, IsSameParty(partyA, partyB))
	partyA, partyB = validParty(), validParty()
	assert.Equal(t, partyA, partyB)
	partyA.BankID = "NationSlide"
	assert.False(t, IsSameParty(partyA, partyB))
	partyA, partyB = validParty(), validParty()
	assert.Equal(t, partyA, partyB)
	partyA.BankIDCode = "n1"
	assert.False(t, IsSameParty(partyA, partyB))
	assert.False(t, IsSameParty(nil, partyB))
	assert.False(t, IsSameParty(nil, nil))
}
