package schema

import (
	"errors"
	"fmt"
)

// 	"charges_information": {
// 	  "$ref": "#/definitions/Charges/example"
// 	},

var paymentTypes = []string{"Credit"}
var schemePaymentSubTypes = []string{"InternetBanking"}
var schemePaymentTypes = []string{"ImmediatePayment"}

// PaymentAttributes resource
type PaymentAttributes struct {
	ID                   uint             `gorm:"primary_key" json:"-"`
	Amount               string           `json:"amount,omitempty"`
	Currency             string           `json:"currency,omitempty"`
	EndToEndReference    string           `json:"end_to_end_reference,omitempty"`
	NumericReference     string           `json:"numeric_reference,omitempty"`
	PaymentID            string           `json:"payment_id,omitempty"`
	PaymentPurpose       string           `json:"payment_purpose,omitempty"`
	PaymentScheme        string           `json:"payment_scheme,omitempty"`
	PaymentType          string           `json:"payment_type,omitempty"`
	ProcessingDate       string           `json:"processing_date,omitempty"`
	Reference            string           `json:"reference,omitempty"`
	SchemePaymentSubType string           `json:"scheme_payment_sub_type,omitempty"`
	SchemePaymentType    string           `json:"scheme_payment_type,omitempty"`
	InternalPaymentID    string           `gorm:"unique;not null" json:"-"`
	BeneficiaryParty     Party            `json:"beneficiary_party,omitempty"`
	BeneficiaryPartyID   uint             `json:"-"`
	DebtorParty          Party            `json:"debtor_party,omitempty"`
	DebtorPartyID        uint             `json:"-"`
	SponsorParty         Party            `json:"sponsor_party,omitempty"`
	SponsorPartyID       uint             `json:"-"`
	ForeignExchange      CurrencyExchange `json:"fx,omitempty"`
	ChargesInformation   Charges          `json:"charges_information,omitempty"`
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
		validationErrors = append(validationErrors, errors.New("Currency cannot be empty"))
	}
	if attributes.EndToEndReference == "" {
		validationErrors = append(validationErrors, errors.New("EndToEndReference cannot be empty"))
	}
	if attributes.NumericReference == "" {
		validationErrors = append(validationErrors, errors.New("NumericReference cannot be empty"))
	}
	if attributes.PaymentID == "" {
		validationErrors = append(validationErrors, errors.New("Amount cannot be empty"))
	}
	if attributes.PaymentPurpose == "" {
		validationErrors = append(validationErrors, errors.New("Currency cannot be empty"))
	}
	if attributes.PaymentScheme == "" {
		validationErrors = append(validationErrors, errors.New("EndToEndReference cannot be empty"))
	}
	if attributes.ProcessingDate == "" {
		validationErrors = append(validationErrors, errors.New("NumericReference cannot be empty"))
	}
	if attributes.Reference == "" {
		validationErrors = append(validationErrors, errors.New("NumericReference cannot be empty"))
	}
	if attributes.InternalPaymentID != "" && !isUUID(attributes.InternalPaymentID) {
		validationErrors = append(validationErrors, errors.New("InternalPaymentID invalid, must be purely alphanumeric with dashes"))
	}
	validationErrors = append(validationErrors, validateParties(&attributes.BeneficiaryParty, &attributes.DebtorParty, &attributes.SponsorParty)...)
	validationErrors = append(validationErrors, validateCurrencyExchange(&attributes.ForeignExchange)...)
	return append(validationErrors, validateCharges(&attributes.ChargesInformation)...)
}
