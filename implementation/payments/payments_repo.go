package payments

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

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
	return nil, fmt.Errorf(strings.Join(errstrings, ", "))
}
