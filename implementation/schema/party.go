package schema

import "errors"

// Party resource
type Party struct {
	ID                uint   `gorm:"primary_key" json:"-"`
	AccountName       string `json:"account_name,omitempty"`
	AccountNumber     string `gorm:"unique_index:bankacc" json:"account_number,omitempty"`
	AccountNumberCode string `json:"account_number_code,omitempty"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address,omitempty"`
	BankID            string `gorm:"unique_index:bankacc" json:"bank_id,omitempty"`
	BankIDCode        string `gorm:"unique_index:bankacc" json:"bank_id_code,omitempty"`
	Name              string `json:"name,omitempty"`
}

func validateParties(parties ...*Party) []error {
	validationErrors := []error{}
	for _, party := range parties {
		validationErrors = append(validationErrors, validateParty(party)...)
	}
	return validationErrors
}

func validateParty(party *Party) []error {
	validationErrors := []error{}
	if party.AccountNumber == "" {
		validationErrors = append(validationErrors, errors.New("AccountNumber cannot be empty"))
	}
	if party.BankID == "" {
		validationErrors = append(validationErrors, errors.New("BankID cannot be empty"))
	}
	if party.BankIDCode == "" {
		validationErrors = append(validationErrors, errors.New("BankIDCode cannot be empty"))
	}
	return validationErrors
}

// IsSameParty checks equality of party key properties
func IsSameParty(partyA *Party, partyB *Party) bool {
	return partyA != nil &&
		partyB != nil &&
		partyA.AccountNumber == partyB.AccountNumber &&
		partyA.BankID == partyB.BankID &&
		partyA.BankIDCode == partyB.BankIDCode
}
