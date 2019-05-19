package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateMoney_Returns_Error_When_Currency_Is_Empty(t *testing.T) {
	invalidMoney := ValidMoney()
	invalidMoney.Currency = ""
	errs := validateMoney(&invalidMoney)
	assert.NotEmpty(t, errs)
}

func Test_validateMoney_Returns_Error_When_Amount_Is_Invalid(t *testing.T) {
	invalidMoney := ValidMoney()
	invalidMoney.Amount = "narp"
	errs := validateMoney(&invalidMoney)
	assert.NotEmpty(t, errs)
}

func Test_validateMoney_Returns_No_Error_When_Money_Is_Valid(t *testing.T) {
	validMoney := ValidMoney()
	errs := validateMoney(&validMoney)
	assert.Empty(t, errs)
}
