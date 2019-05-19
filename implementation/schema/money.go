package schema

import (
	"errors"
	"fmt"
)

// Money makes the world go round
type Money struct {
	ID        uint   `gorm:"primary_key" json:"-"`
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	ChargesID uint   `gorm:"not null" json:"-"`
}

func validateMoney(money *Money) []error {
	validationErrors := []error{}
	if !isAmount(money.Amount) {
		validationErrors = append(validationErrors, fmt.Errorf("Invalid amount: %s", money.Amount))
	}
	if money.Currency == "" {
		validationErrors = append(validationErrors, errors.New("Currency cannot be empty"))
	}
	return validationErrors
}
