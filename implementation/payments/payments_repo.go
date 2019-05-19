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
	err := DB.Preload("Attributes").Preload("Attributes.BeneficiaryParty").Preload("Attributes.DebtorParty").Preload("Attributes.SponsorParty").Preload("Attributes.ForeignExchange").Preload("Attributes.ChargesInformation").Find(&allPayments).Error
	return allPayments, err
}

// Create a new payment
func Create(DB *gorm.DB, payment *schema.Payment) (*schema.Payment, error) {
	validationErrors := schema.Validate(payment)
	if len(validationErrors) != 0 {
		return nil, consolidateValidationErrors(validationErrors, fmt.Sprintf("Validation errors whilst creating payment with id %s", payment.ID))
	}

	if createErr := DB.Create(&payment).Error; createErr != nil {
		if !isUniqueConstraintError(createErr, "bankacc") {
			return nil, createErr
		}
		payment, err := loadAssociations(DB, payment)
		if err != nil {
			return nil, err
		}
		if createErr = DB.Set("gorm:association_autocreate", false).Create(&payment).Error; createErr != nil {
			return nil, createErr
		}
		if createErr = DB.Set("gorm:association_autocreate", false).Create(&payment.Attributes).Error; createErr != nil {
			return nil, createErr
		}
		if createErr = DB.Create(&payment.Attributes.ForeignExchange).Error; createErr != nil {
			return nil, createErr
		}
		if createErr = DB.Create(&payment.Attributes.ChargesInformation).Error; createErr != nil {
			return nil, createErr
		}
		if saveErr := DB.Set("gorm:association_autocreate", false).Save(&payment).Error; saveErr != nil {
			return nil, saveErr
		}
	}

	return payment, nil
}

func loadAssociations(DB *gorm.DB, payment *schema.Payment) (*schema.Payment, error) {
	payment.Attributes.InternalPaymentID = payment.ID
	beneficiaryParty := schema.Party{}
	if err := DB.Where(&schema.Party{
		AccountNumber: payment.Attributes.BeneficiaryParty.AccountNumber,
		BankID:        payment.Attributes.BeneficiaryParty.BankID,
		BankIDCode:    payment.Attributes.BeneficiaryParty.BankIDCode,
	}).First(&beneficiaryParty).Error; err == nil {
		payment.Attributes.BeneficiaryParty.ID = beneficiaryParty.ID
		payment.Attributes.BeneficiaryPartyID = beneficiaryParty.ID
	}
	if err := DB.Save(&payment.Attributes.BeneficiaryParty).Error; err != nil {
		return nil, err
	}

	debtorParty := schema.Party{}
	if err := DB.Where(&schema.Party{
		AccountNumber: payment.Attributes.DebtorParty.AccountNumber,
		BankID:        payment.Attributes.DebtorParty.BankID,
		BankIDCode:    payment.Attributes.DebtorParty.BankIDCode,
	}).First(&debtorParty).Error; err == nil {
		payment.Attributes.DebtorParty.ID = debtorParty.ID
		payment.Attributes.DebtorPartyID = debtorParty.ID
	}
	if err := DB.Save(&payment.Attributes.DebtorParty).Error; err != nil {
		return nil, err
	}

	sponsorParty := schema.Party{}
	if err := DB.Where(&schema.Party{
		AccountNumber: payment.Attributes.SponsorParty.AccountNumber,
		BankID:        payment.Attributes.SponsorParty.BankID,
		BankIDCode:    payment.Attributes.SponsorParty.BankIDCode,
	}).First(&sponsorParty).Error; err == nil {
		payment.Attributes.SponsorParty.ID = sponsorParty.ID
		payment.Attributes.SponsorPartyID = sponsorParty.ID
	}
	if err := DB.Save(&payment.Attributes.SponsorParty).Error; err != nil {
		return nil, err
	}

	return payment, nil
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
	if err = DB.Preload("Attributes").Preload("Attributes.BeneficiaryParty").Preload("Attributes.DebtorParty").Preload("Attributes.SponsorParty").Preload("Attributes.ForeignExchange").Preload("Attributes.ChargesInformation").Where(&schema.Payment{ID: ID}).First(&payment).Error; err != nil {
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

	if existingPayment == payment {
		return nil, nil
	} else if err = DB.Set("gorm:association_autocreate", false).Save(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func syncAssociations(DB *gorm.DB, from, to *schema.Payment) *schema.Payment {
	to.Attributes.ID = from.Attributes.ID
	to.Attributes.InternalPaymentID = from.Attributes.InternalPaymentID
	to.Attributes.ForeignExchange.ID = from.Attributes.ForeignExchange.ID
	to.Attributes.ForeignExchange.PaymentAttributesID = from.Attributes.ID
	to.Attributes.ChargesInformation.ID = from.Attributes.ChargesInformation.ID
	to.Attributes.ChargesInformation.PaymentAttributesID = from.Attributes.ID
	p, _ := loadAssociations(DB, to)
	return p
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

func isUniqueConstraintError(err error, constraintName string) bool {
	return strings.Contains(err.Error(), "unique constraint") && strings.Contains(err.Error(), constraintName)
}
