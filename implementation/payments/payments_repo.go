package payments

import (
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
	err := DB.Save(&payment).Error
	return payment, err
}
