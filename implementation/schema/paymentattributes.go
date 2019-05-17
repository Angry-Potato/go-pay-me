package schema

import (
	"errors"
	"fmt"
)

// 	"beneficiary_party": {
// 	  "$ref": "#/definitions/Party/example"
// 	},
// 	"charges_information": {
// 	  "$ref": "#/definitions/Charges/example"
// 	},
// 	"debtor_party": {
// 	  "$ref": "#/definitions/Party/example"
// 	},
// 	"fx": {
// 	  "$ref": "#/definitions/CurrencyExchange/example"
// 	},
// 	"sponsor_party": {
// 	  "$ref": "#/definitions/Party/example"
// 	}

var paymentTypes = []string{"Credit"}
var schemePaymentSubTypes = []string{"InternetBanking"}
var schemePaymentTypes = []string{"ImmediatePayment"}

// PaymentAttributes resource
type PaymentAttributes struct {
	ID                   uint   `gorm:"primary_key" json:"-"`
	Amount               string `json:"amount"`
	Currency             string `json:"currency"`
	EndToEndReference    string `json:"end_to_end_reference"`
	NumericReference     string `json:"numeric_reference"`
	PaymentID            string `json:"payment_id"`
	PaymentPurpose       string `json:"payment_purpose"`
	PaymentScheme        string `json:"payment_scheme"`
	PaymentType          string `json:"payment_type"`
	ProcessingDate       string `json:"processing_date"`
	Reference            string `json:"reference"`
	SchemePaymentSubType string `json:"scheme_payment_sub_type"`
	SchemePaymentType    string `json:"scheme_payment_type"`
	InternalPaymentID    string `gorm:"unique;not null" json:"-"`
	BeneficiaryParty     Party  `json:"beneficiary_party"`
	BeneficiaryPartyID   uint   `json:"-"`
}

func validatePaymentAttributes(attributes *PaymentAttributes) []error {
	validationErrors := []error{}
	if !contains(paymentTypes, attributes.PaymentType) {
		validationErrors = append(validationErrors, fmt.Errorf("Unknown payment type on PaymentAttributes: %s", attributes.PaymentType))
	}
	if !contains(schemePaymentSubTypes, attributes.SchemePaymentSubType) {
		validationErrors = append(validationErrors, fmt.Errorf("Unknown scheme payment sub type: %s", attributes.SchemePaymentSubType))
	}
	if !contains(schemePaymentTypes, attributes.SchemePaymentType) {
		validationErrors = append(validationErrors, fmt.Errorf("Unknown scheme payment type: %s", attributes.SchemePaymentType))
	}
	if !isAmount(attributes.Amount) {
		validationErrors = append(validationErrors, fmt.Errorf("Invalid amount: %s", attributes.Amount))
	}
	if attributes.Currency == "" {
		validationErrors = append(validationErrors, errors.New("Currency cannot be empty."))
	}
	if attributes.EndToEndReference == "" {
		validationErrors = append(validationErrors, errors.New("EndToEndReference cannot be empty."))
	}
	if attributes.NumericReference == "" {
		validationErrors = append(validationErrors, errors.New("NumericReference cannot be empty."))
	}
	if attributes.PaymentID == "" {
		validationErrors = append(validationErrors, errors.New("Amount cannot be empty."))
	}
	if attributes.PaymentPurpose == "" {
		validationErrors = append(validationErrors, errors.New("Currency cannot be empty."))
	}
	if attributes.PaymentScheme == "" {
		validationErrors = append(validationErrors, errors.New("EndToEndReference cannot be empty."))
	}
	if attributes.ProcessingDate == "" {
		validationErrors = append(validationErrors, errors.New("NumericReference cannot be empty."))
	}
	if attributes.Reference == "" {
		validationErrors = append(validationErrors, errors.New("NumericReference cannot be empty."))
	}
	if attributes.InternalPaymentID != "" && !isUUID(attributes.InternalPaymentID) {
		validationErrors = append(validationErrors, errors.New("InternalPaymentID invalid, must be purely alphanumeric with dashes"))
	}
	return append(validationErrors, validateParty(&attributes.BeneficiaryParty)...)
}
