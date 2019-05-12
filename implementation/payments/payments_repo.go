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
