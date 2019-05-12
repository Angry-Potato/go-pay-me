package payments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validPayment() *Payment {
	return &Payment{
		ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
		Type:           "Payment",
		Version:        0,
		OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	}
}

func Test_Validate_Returns_Error_When_ID_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.ID = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_ID_Is_Invalid(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.ID = "I don't think UUIDs contain spaces, man"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_Type_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Type = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_Type_Is_Invalid(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Type = "This will never be a type"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_Version_Is_Invalid(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.Version = -1
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_OrganisationID_Is_Empty(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.OrganisationID = ""
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_Error_When_OrganisationID_Is_Invalid(t *testing.T) {
	invalidPayment := validPayment()
	invalidPayment.OrganisationID = "I don't think UUIDs contain spaces, man"
	errs := Validate(invalidPayment)
	assert.NotEmpty(t, errs)
}

func Test_Validate_Returns_No_Errors_When_Payment_Is_Valid(t *testing.T) {
	invalidPayment := validPayment()
	errs := Validate(invalidPayment)
	assert.Empty(t, errs)
}
