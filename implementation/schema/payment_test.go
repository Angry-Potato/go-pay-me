// +build unit

package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate_Returns_Error_When_ID_Is_Empty(t *testing.T) {
	invalidPayment := ValidPayment()
	invalidPayment.ID = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_ID_Is_Invalid(t *testing.T) {
	invalidPayment := ValidPayment()
	invalidPayment.ID = "I don't think UUIDs contain spaces, man"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_Type_Is_Empty(t *testing.T) {
	invalidPayment := ValidPayment()
	invalidPayment.Type = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_Type_Is_Invalid(t *testing.T) {
	invalidPayment := ValidPayment()
	invalidPayment.Type = "This will never be a type"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_Version_Is_Invalid(t *testing.T) {
	invalidPayment := ValidPayment()
	invalidPayment.Version = -1
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_OrganisationID_Is_Empty(t *testing.T) {
	invalidPayment := ValidPayment()
	invalidPayment.OrganisationID = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_OrganisationID_Is_Invalid(t *testing.T) {
	invalidPayment := ValidPayment()
	invalidPayment.OrganisationID = "I don't think UUIDs contain spaces, man"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_No_Errors_When_Payment_Is_Valid(t *testing.T) {
	invalidPayment := ValidPayment()
	errs := Validate(invalidPayment)
	assert.Empty(t, errs)
}
