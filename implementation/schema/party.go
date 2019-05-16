package schema

import "errors"

// Party resource
type Party struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `gorm:"primary_key" json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `gorm:"primary_key" json:"bank_id"`
	BankIDCode        string `gorm:"primary_key" json:"bank_id_code"`
	Name              string `json:"name"`
}

func validateParty(party *Party) []error {
	validationErrors := []error{}
	if party.AccountNumber == "" {
		validationErrors = append(validationErrors, errors.New("AccountNumber cannot be empty."))
	}
	if party.BankID == "" {
		validationErrors = append(validationErrors, errors.New("BankID cannot be empty."))
	}
	if party.BankIDCode == "" {
		validationErrors = append(validationErrors, errors.New("BankIDCode cannot be empty."))
	}
	return validationErrors
}
