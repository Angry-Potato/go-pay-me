// +build unit

package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validatePaymentAttributes_Returns_Error_When_Amount_Is_Invalid(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.Amount = ".0.0"
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_Currency_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.Currency = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_EndToEndReference_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.EndToEndReference = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_NumericReference_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.NumericReference = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentID_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentID = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentPurpose_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentPurpose = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentScheme_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentScheme = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_ProcessingDate_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.ProcessingDate = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_Reference_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.Reference = ""
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_PaymentType_Is_Unknown(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentType = "oh dear, what's happened here?"
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_SchemePaymentSubType_Is_Unknown(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.SchemePaymentSubType = "oh dear, what's happened here?"
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_SchemePaymentType_Is_Unknown(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.SchemePaymentType = "oh dear, what's happened here?"
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_Error_When_InternalPaymentID_Is_Not_A_UUID(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.InternalPaymentID = "nope, definitely not."
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.NotEmpty(t, errs)
}

func Test_validatePaymentAttributes_Returns_No_Error_When_Are_Valid(t *testing.T) {
	invalidPayment := validPayment()
	errs := validatePaymentAttributes(&invalidPayment.Attributes)
	assert.Empty(t, errs)
}
