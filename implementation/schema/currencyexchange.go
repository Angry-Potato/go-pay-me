package schema

import (
	"errors"
	"fmt"
)

// CurrencyExchange resource
type CurrencyExchange struct {
	ID                  uint   `gorm:"primary_key" json:"-"`
	ContractReference   string `json:"contract_reference"`
	ExchangeRate        string `json:"exchange_rate"`
	OriginalAmount      string `json:"original_amount"`
	OriginalCurrency    string `json:"original_currency"`
	PaymentAttributesID uint   `gorm:"unique;not null" json:"-"`
}

func validateCurrencyExchange(currencyExchange *CurrencyExchange) []error {
	validationErrors := []error{}
	if !isAmount(currencyExchange.OriginalAmount) {
		validationErrors = append(validationErrors, fmt.Errorf("Invalid original amount: %s", currencyExchange.OriginalAmount))
	}
	if currencyExchange.ContractReference == "" {
		validationErrors = append(validationErrors, errors.New("ContractReference cannot be empty"))
	}
	if currencyExchange.ExchangeRate == "" {
		validationErrors = append(validationErrors, errors.New("ExchangeRate cannot be empty"))
	}
	if currencyExchange.OriginalCurrency == "" {
		validationErrors = append(validationErrors, errors.New("OriginalCurrency cannot be empty"))
	}
	return validationErrors
}
