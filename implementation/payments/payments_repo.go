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
	err := DB.Preload("Attributes").Preload("Attributes.BeneficiaryParty").Find(&allPayments).Error
	return allPayments, err
}

// Create a new payment
func Create(DB *gorm.DB, payment *schema.Payment) (*schema.Payment, error) {
	validationErrors := schema.Validate(payment)
	if len(validationErrors) != 0 {
		return nil, consolidateValidationErrors(validationErrors, fmt.Sprintf("Validation errors whilst creating payment with id %s", payment.ID))
	}

	if createErr := DB.Create(&payment).Error; createErr != nil {
		if !isUniqueConstraintError(createErr) {
			return nil, createErr
		}
		payment = loadAssociations(DB, payment)
		if createErr = DB.Create(&payment).Error; createErr != nil {
			return nil, createErr
		}
		if saveErr := DB.Save(&payment).Error; saveErr != nil {
			return nil, saveErr
		}
	}

	return payment, nil
}

func loadAssociations(DB *gorm.DB, payment *schema.Payment) *schema.Payment {
	beneficiaryParty := schema.Party{}
	if err := DB.Where(&schema.Party{
		AccountNumber: payment.Attributes.BeneficiaryParty.AccountNumber,
		BankID:        payment.Attributes.BeneficiaryParty.BankID,
		BankIDCode:    payment.Attributes.BeneficiaryParty.BankIDCode,
	}).First(&beneficiaryParty).Error; err == nil {
		payment.Attributes.BeneficiaryParty.ID = beneficiaryParty.ID
		payment.Attributes.BeneficiaryPartyID = beneficiaryParty.ID
	}
	return payment
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

	if len(consolidatedValidation) != 0 {
		return nil, consolidateValidationErrors(consolidatedValidation, "Errors")
	}

	if err := resetPayments(DB, payments); err != nil {
		return nil, fmt.Errorf("Error batch inserting: %s", err.Error())
	}

	newPayments, err := All(DB)
	if err != nil {
		return nil, err
	}
	return newPayments, nil
}

func resetPayments(DB *gorm.DB, payments []schema.Payment) (result error) {
	if err := DeleteAll(DB); err != nil {
		return err
	}

	for _, payment := range payments {
		if _, err := Create(DB, &payment); err != nil {
			return err
		}
	}
	return nil
}

// Get a payment by ID
func Get(DB *gorm.DB, ID string) (*schema.Payment, error) {
	err := schema.ValidateID(ID)
	if err != nil {
		return nil, &ValidationError{err.Error()}
	}

	payment := schema.Payment{}
	if err = DB.Preload("Attributes").Preload("Attributes.BeneficiaryParty").Where(&schema.Payment{ID: ID}).First(&payment).Error; err != nil {
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
	if err := schema.ValidateID(ID); err != nil {
		return nil, &ValidationError{err.Error()}
	} else if errs := schema.Validate(payment); len(errs) != 0 {
		return nil, consolidateValidationErrors(errs, fmt.Sprintf("Validation errors whilst Updating payment %s", payment.ID))
	}

	existingPayment, err := Get(DB, ID)
	if err != nil {
		return nil, err
	}

	payment = syncAssociations(DB, existingPayment, payment)

	if *existingPayment == *payment {
		return nil, nil
	} else if err = DB.Save(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func syncAssociations(DB *gorm.DB, from, to *schema.Payment) *schema.Payment {
	if to.Attributes.ID == 0 {
		to.Attributes.ID = from.Attributes.ID
		to.Attributes.InternalPaymentID = from.Attributes.InternalPaymentID
	}

	if schema.IsSameParty(&to.Attributes.BeneficiaryParty, &from.Attributes.BeneficiaryParty) {
		to.Attributes.BeneficiaryParty.ID = from.Attributes.BeneficiaryParty.ID
		to.Attributes.BeneficiaryPartyID = from.Attributes.BeneficiaryPartyID
	} else {
		to = loadAssociations(DB, to)
	}
	return to
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

func isUniqueConstraintError(err error) bool {
	return strings.Contains(err.Error(), "unique constraint")
}
