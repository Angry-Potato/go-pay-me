// +build unit

package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate_Returns_Error_When_PaymentAttributes_Amount_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.Amount = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_Currency_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.Currency = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_EndToEndReference_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.EndToEndReference = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_NumericReference_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.NumericReference = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_PaymentID_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentID = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_PaymentPurpose_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentPurpose = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_PaymentScheme_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentScheme = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_ProcessingDate_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.ProcessingDate = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_Reference_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.Reference = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_PaymentType_Is_Unknown(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.PaymentType = "oh dear, what's happened here?"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_SchemePaymentSubType_Is_Unknown(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.SchemePaymentSubType = "oh dear, what's happened here?"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_SchemePaymentType_Is_Unknown(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.SchemePaymentType = "oh dear, what's happened here?"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_PaymentAttributes_InternalPaymentID_Is_Not_A_UUID(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Attributes.InternalPaymentID = "nope, definitely not."
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}
