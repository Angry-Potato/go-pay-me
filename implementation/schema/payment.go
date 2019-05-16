package schema

import (
	"errors"
	"fmt"
	"regexp"
)

var types = []string{"Payment"}

// Payment resource
type Payment struct {
	ID             string            `json:"id"`
	Type           string            `json:"type"`
	Version        int64             `json:"version"`
	OrganisationID string            `json:"organisation_id"`
	Attributes     PaymentAttributes `gorm:"foreignkey:InternalPaymentID" json:"attributes"`
}

// Validate a payment resource
func Validate(payment *Payment) []error {
	return append(validatePayment(payment), validatePaymentAttributes(&payment.Attributes)...)
}

func validatePayment(payment *Payment) []error {
	validationErrors := []error{}
	if err := ValidateID(payment.ID); err != nil {
		validationErrors = append(validationErrors, err)
	}
	if !contains(types, payment.Type) {
		validationErrors = append(validationErrors, fmt.Errorf("Unknown payment type: %s", payment.Type))
	}
	if payment.Version < 0 {
		validationErrors = append(validationErrors, errors.New("Version cannot be less than zero"))
	}
	if !isUUID(payment.OrganisationID) {
		validationErrors = append(validationErrors, errors.New("OrganisationID invalid, must be purely alphanumeric with dashes"))
	}
	return validationErrors
}

// ValidateID of a payment
func ValidateID(ID string) error {
	if ID == "" {
		return errors.New("ID cannot be empty")
	}
	if !isUUID(ID) {
		return errors.New("ID invalid, must be purely alphanumeric with dashes")
	}
	return nil
}

func contains(haystack []string, needle string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}

func isUUID(uuid string) bool {
	if matches, err := regexp.MatchString("^[a-zA-Z0-9-]+$", uuid); !matches || err != nil {
		return false
	}
	return true
}
