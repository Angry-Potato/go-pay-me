package payments

import (
	"fmt"
	"strings"

	"github.com/Angry-Potato/go-pay-me/implementation/schema"
	"github.com/jinzhu/gorm"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ValidationError describes the incorrectness of a payment operation
type ValidationError struct {
	err string
}

func (e *ValidationError) Error() string {
	return e.err
}

// All payments
func All(DB *gorm.DB) ([]schema.Payment, error) {
	allPayments := []schema.Payment{}
	err := DB.Preload("Attributes").Find(&allPayments).Error
	return allPayments, err
}

// Create a new payment
func Create(DB *gorm.DB, payment *schema.Payment) (*schema.Payment, error) {
	validationErrors := schema.Validate(payment)
	if len(validationErrors) == 0 {
		if err := DB.Create(&payment).Error; err != nil {
			return nil, err
		}
		return payment, nil
	}

	return nil, consolidateValidationErrors(validationErrors, fmt.Sprintf("Validation errors whilst creating payment with id %s", payment.ID))
}

// DeleteAll payments
func DeleteAll(DB *gorm.DB) error {
	allPayments := []schema.Payment{}
	err := DB.Delete(&allPayments).Error
	return err
}

// SetAll payments
func SetAll(DB *gorm.DB, payments []schema.Payment) ([]schema.Payment, error) {
	var consolidatedValidation []error
	for _, payment := range payments {
		validationErrors := schema.Validate(&payment)
		if len(validationErrors) != 0 {
			consolidatedValidation = append(consolidatedValidation, consolidateValidationErrors(validationErrors, fmt.Sprintf("Validation errors whilst creating payment with id %s", payment.ID)))
		}
	}

	if len(consolidatedValidation) == 0 {
		DB.Lock()
		defer DB.Unlock()
		err := DeleteAll(DB)
		if err != nil {
			return nil, err
		}

		successes := 0
		errs := []error{}
		for _, payment := range payments {
			if _, err = Create(DB, &payment); err != nil {
				errs = append(errs, err)
			} else {
				successes++
			}
		}

		if len(errs) != 0 {
			return nil, consolidateErrors(errs, "Error(s) batch inserting: ")
		} else if successes != len(payments) {
			return nil, fmt.Errorf("Could only insert %d out of %d", successes, len(payments))
		}
		newPayments, err := All(DB)
		if err != nil {
			return nil, err
		}
		return newPayments, nil
	}

	return nil, consolidateValidationErrors(consolidatedValidation, "Errors")
}

// Get a payment by ID
func Get(DB *gorm.DB, ID string) (*schema.Payment, error) {
	err := schema.ValidateID(ID)
	if err != nil {
		return nil, &ValidationError{err.Error()}
	}

	payment := schema.Payment{}
	if err = DB.Preload("Attributes").Where(&schema.Payment{ID: ID}).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

// Delete a payment by ID
func Delete(DB *gorm.DB, ID string) error {
	err := schema.ValidateID(ID)
	if err != nil {
		return &ValidationError{err.Error()}
	}

	DB = DB.Delete(&schema.Payment{ID: ID})
	if err = DB.Error; err != nil {
		return err
	}
	if DB.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Update a payment by ID
func Update(DB *gorm.DB, ID string, payment *schema.Payment) (*schema.Payment, error) {
	err := schema.ValidateID(ID)
	if err != nil {
		return nil, &ValidationError{err.Error()}
	}
	errs := schema.Validate(payment)
	if len(errs) != 0 {
		return nil, consolidateValidationErrors(errs, fmt.Sprintf("Validation errors whilst Updating payment %s", payment.ID))
	}

	existingPayment, err := Get(DB, ID)
	if err != nil {
		return nil, err
	}
	if *existingPayment == *payment {
		return nil, nil
	}

	existingPayment.ID = payment.ID
	existingPayment.Type = payment.Type
	existingPayment.Version = payment.Version
	existingPayment.OrganisationID = payment.OrganisationID
	existingPayment.Attributes = payment.Attributes
	//Preload("Attributes")
	DB = DB.Save(existingPayment).P
	if err = DB.Error; err != nil {
		return nil, err
	}
	if DB.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return payment, nil
}

func consolidateErrors(errs []error, prefix string) error {
	var errstrings []string
	for _, err := range errs {
		errstrings = append(errstrings, err.Error())
	}
	return fmt.Errorf("%s: %s", prefix, strings.Join(errstrings, ", "))
}

func consolidateValidationErrors(errs []error, prefix string) error {
	var errstrings []string
	for _, err := range errs {
		errstrings = append(errstrings, err.Error())
	}
	return &ValidationError{fmt.Sprintf("%s: %s", prefix, strings.Join(errstrings, ", "))}
}
