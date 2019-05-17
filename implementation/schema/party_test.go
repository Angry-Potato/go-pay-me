package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateParty_Returns_Error_When_AccountNumber_Is_Empty(t *testing.T) {
	invalidParty := ValidParty()
	invalidParty.AccountNumber = ""
	errs := validateParty(&invalidParty)
	assert.NotEmpty(t, errs)
}

func Test_validateParty_Returns_Error_When_BankID_Is_Empty(t *testing.T) {
	invalidParty := ValidParty()
	invalidParty.BankID = ""
	errs := validateParty(&invalidParty)
	assert.NotEmpty(t, errs)
}

func Test_validateParty_Returns_Error_When_BankIDCode_Is_Empty(t *testing.T) {
	invalidParty := ValidParty()
	invalidParty.BankIDCode = ""
	errs := validateParty(&invalidParty)
	assert.NotEmpty(t, errs)
}

func Test_validateParty_Returns_No_Error_When_Party_Valid(t *testing.T) {
	validParty := ValidParty()
	errs := validateParty(&validParty)
	assert.Empty(t, errs)
}

func Test_validateParties_Returns_Error_When_Parties_Invalid(t *testing.T) {
	invalidPartyA, invalidPartyB, validParty := ValidParty(), ValidParty(), ValidParty()
	invalidPartyA.BankIDCode = ""
	invalidPartyB.BankID = ""
	errs := validateParties(&invalidPartyA, &invalidPartyB, &validParty)
	assert.NotEmpty(t, errs)
}

func Test_validateParties_Returns_No_Error_When_Parties_Valid(t *testing.T) {
	validPartyA, validPartyB, validPartyC := ValidParty(), ValidParty(), ValidParty()
	errs := validateParties(&validPartyA, &validPartyB, &validPartyC)
	assert.Empty(t, errs)
}

func Test_IsSameParty_Returns_True_When_Parties_Are_Equal(t *testing.T) {
	partyA, partyB := ValidParty(), ValidParty()
	assert.Equal(t, partyA, partyB)
	assert.True(t, IsSameParty(&partyA, &partyB))
	partyA.AccountName = "Her account"
	partyA.AccountNumberCode = "45678"
	partyA.AccountType = 90
	partyA.Address = "321 street"
	partyA.Name = "Miss Lady"
	assert.True(t, IsSameParty(&partyA, &partyB))
}

func Test_IsSameParty_Returns_False_When_Parties_Are_Not_Equal(t *testing.T) {
	partyA, partyB := ValidParty(), ValidParty()
	assert.Equal(t, partyA, partyB)
	partyA.AccountNumber = "some new account"
	assert.False(t, IsSameParty(&partyA, &partyB))
	partyA, partyB = ValidParty(), ValidParty()
	assert.Equal(t, partyA, partyB)
	partyA.BankID = "NationSlide"
	assert.False(t, IsSameParty(&partyA, &partyB))
	partyA, partyB = ValidParty(), ValidParty()
	assert.Equal(t, partyA, partyB)
	partyA.BankIDCode = "n1"
	assert.False(t, IsSameParty(&partyA, &partyB))
	assert.False(t, IsSameParty(nil, &partyB))
	assert.False(t, IsSameParty(nil, nil))
}
