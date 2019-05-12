package payments

import (
	"errors"
	"fmt"
	"regexp"
)

var knownTypes = []string{"Payment"}

// Payment resource
type Payment struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	Version        int64  `json:"version"`
	OrganisationID string `json:"organisation_id"`
}

// Validate a payment resource
func Validate(payment *Payment) []error {
	validationErrors := []error{}
	if payment.ID == "" {
		validationErrors = append(validationErrors, errors.New("ID cannot be empty"))
	}
	if !isUUID(payment.ID) {
		validationErrors = append(validationErrors, errors.New("ID invalid, must be purely alphanumeric with dashes"))
	}
	if !isKnownType(payment.Type) {
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

func isKnownType(alienType string) bool {
	for _, knownType := range knownTypes {
		if knownType == alienType {
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
