package payments

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// ValidationError describes the incorrectness of a payment operation
type ValidationError struct {
	err string
}

func (e *ValidationError) Error() string {
	return e.err
}

// All payments
func All(DB *gorm.DB) ([]Payment, error) {
	allPayments := []Payment{}
	err := DB.Find(&allPayments).Error
	return allPayments, err
}

// Create a new payment
func Create(DB *gorm.DB, payment *Payment) (*Payment, error) {
	validationErrors := Validate(payment)
	if len(validationErrors) == 0 {
		if err := DB.Create(&payment).Error; err != nil {
			return nil, err
		}
		return payment, nil
	}

	var errstrings []string
	for _, validationError := range validationErrors {
		errstrings = append(errstrings, validationError.Error())
	}
	return nil, &ValidationError{strings.Join(errstrings, ", ")}
}
