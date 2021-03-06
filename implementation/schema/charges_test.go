package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateCharges_Returns_Error_When_BearerCode_Is_Empty(t *testing.T) {
	invalidCharges := ValidCharges()
	invalidCharges.BearerCode = ""
	errs := validateCharges(&invalidCharges)
	assert.NotEmpty(t, errs)
}

func Test_validateCharges_Returns_Error_When_ReceiverChargesCurrency_Is_Empty(t *testing.T) {
	invalidCharges := ValidCharges()
	invalidCharges.ReceiverChargesCurrency = ""
	errs := validateCharges(&invalidCharges)
	assert.NotEmpty(t, errs)
}

func Test_validateCharges_Returns_Error_When_ReceiverChargesAmount_Is_Invalid(t *testing.T) {
	invalidCharges := ValidCharges()
	invalidCharges.ReceiverChargesAmount = "narp"
	errs := validateCharges(&invalidCharges)
	assert.NotEmpty(t, errs)
}

func Test_validateCharges_Returns_Error_When_SenderCharges_Has_Invalid_Money(t *testing.T) {
	invalidCharges, validMoney := ValidCharges(), ValidMoney()
	invalidCharges.SenderCharges = []*Money{&validMoney, &Money{Amount: "uh oh", Currency: "USD"}}
	errs := validateCharges(&invalidCharges)
	assert.NotEmpty(t, errs)
}

func Test_validateCharges_Returns_No_Error_When_Charges_Has_Empty_SenderCharges(t *testing.T) {
	validCharges := ValidCharges()
	validCharges.SenderCharges = []*Money{}
	errs := validateCharges(&validCharges)
	assert.Empty(t, errs)
}

func Test_validateCharges_Returns_No_Error_When_Charges_Is_Valid(t *testing.T) {
	validCharges := ValidCharges()
	errs := validateCharges(&validCharges)
	assert.Empty(t, errs)
}
