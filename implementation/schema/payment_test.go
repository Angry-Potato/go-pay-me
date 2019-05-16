// +build unit

package schema

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func validPayment() *Payment {
	ID := uuid.New().String()
	return &Payment{
		ID:             ID,
		Type:           "Payment",
		Version:        0,
		OrganisationID: uuid.New().String(),
		Attributes: PaymentAttributes{
			Amount:               "some amount",
			Currency:             "great",
			EndToEndReference:    "here it is",
			NumericReference:     "1245",
			PaymentID:            "343535",
			PaymentPurpose:       "stuff",
			PaymentScheme:        "best",
			PaymentType:          "Credit",
			ProcessingDate:       "now",
			Reference:            "that guy",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "ImmediatePayment",
		},
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
