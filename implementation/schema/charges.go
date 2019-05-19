package schema

import (
	"errors"
	"fmt"
)

// "sender_charges": [
//   {
// 	"amount": "5.00",
// 	"currency": "GBP"
//   },
//   {
// 	"amount": "10.00",
// 	"currency": "GBP"
//   }
// ],

// Charges resource
type Charges struct {
	ID                      uint   `gorm:"primary_key" json:"-"`
	BearerCode              string `json:"bearer_code"`
	ReceiverChargesAmount   string `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string `json:"receiver_charges_currency"`
}

func validateCharges(charges *Charges) []error {
	validationErrors := []error{}
	if !isAmount(charges.ReceiverChargesAmount) {
		validationErrors = append(validationErrors, fmt.Errorf("Invalid receiver charges amount: %s", charges.ReceiverChargesAmount))
	}
	if charges.BearerCode == "" {
		validationErrors = append(validationErrors, errors.New("BearerCode cannot be empty"))
	}
	if charges.ReceiverChargesCurrency == "" {
		validationErrors = append(validationErrors, errors.New("ReceiverChargesCurrency cannot be empty"))
	}
	return validationErrors
}
