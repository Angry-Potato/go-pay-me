// +build unit

package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validatePaymentAttributes_Returns_Error_When_Amount_Is_Invalid(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.Amount = ".0.0"
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_Currency_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.Currency = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_EndToEndReference_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.EndToEndReference = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_NumericReference_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.NumericReference = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentID_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.PaymentID = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentPurpose_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.PaymentPurpose = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentScheme_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.PaymentScheme = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_ProcessingDate_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.ProcessingDate = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_Reference_Is_Empty(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.Reference = ""
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentType_Is_Unknown(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.PaymentType = "oh dear, what's happened here?"
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_SchemePaymentSubType_Is_Unknown(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.SchemePaymentSubType = "oh dear, what's happened here?"
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_SchemePaymentType_Is_Unknown(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.SchemePaymentType = "oh dear, what's happened here?"
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_InternalPaymentID_Is_Not_A_UUID(t *testing.T) {
	attributes := ValidPaymentAttributes()
	attributes.InternalPaymentID = "nope, definitely not."
	errs := validatePaymentAttributes(&attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_No_Error_When_Attributes_Are_Valid(t *testing.T) {
	attributes := ValidPaymentAttributes()
	errs := validatePaymentAttributes(&attributes)
	assert.Empty(t, errs)
}
