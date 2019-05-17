package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateCurrencyExchange_Returns_Error_When_ContractReference_Is_Empty(t *testing.T) {
	invalidCurrencyExchange := ValidCurrencyExchange()
	invalidCurrencyExchange.ContractReference = ""
	errs := validateCurrencyExchange(&invalidCurrencyExchange)
	assert.NotEmpty(t, errs)
}

func Test_validateCurrencyExchange_Returns_Error_When_ExchangeRate_Is_Empty(t *testing.T) {
	invalidCurrencyExchange := ValidCurrencyExchange()
	invalidCurrencyExchange.ExchangeRate = ""
	errs := validateCurrencyExchange(&invalidCurrencyExchange)
	assert.NotEmpty(t, errs)
}

func Test_validateCurrencyExchange_Returns_Error_When_OriginalCurrency_Is_Empty(t *testing.T) {
	invalidCurrencyExchange := ValidCurrencyExchange()
	invalidCurrencyExchange.OriginalCurrency = ""
	errs := validateCurrencyExchange(&invalidCurrencyExchange)
	assert.NotEmpty(t, errs)
}

func Test_validateCurrencyExchange_Returns_Error_When_OriginalAmount_Is_Invalid(t *testing.T) {
	invalidCurrencyExchange := ValidCurrencyExchange()
	invalidCurrencyExchange.OriginalAmount = ".0.0."
	errs := validateCurrencyExchange(&invalidCurrencyExchange)
	assert.NotEmpty(t, errs)
}
