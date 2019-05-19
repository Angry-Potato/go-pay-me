// +build unit

package payments

import (
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/schema"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_isSamePayment_Returns_True_When_Given_Two_Like_Payments(t *testing.T) {
	paymentA := schema.ValidPayment()
	assert.True(t, isSamePayment(*paymentA, *paymentA))
}

func Test_isSamePayment_Returns_True_When_Given_Two_Like_Payments_Bar_SenderCharges_IDs(t *testing.T) {
	ID := uuid.New().String()
	paymentA, paymentB := &schema.Payment{
		ID:             ID,
		Type:           "Payment",
		Version:        0,
		OrganisationID: ID,
		Attributes:     schema.ValidPaymentAttributes(),
	}, &schema.Payment{
		ID:             ID,
		Type:           "Payment",
		Version:        0,
		OrganisationID: ID,
		Attributes:     schema.ValidPaymentAttributes(),
	}

	paymentA.Attributes.ChargesInformation = schema.ValidCharges()
	for i, moneyA := range paymentA.Attributes.ChargesInformation.SenderCharges {
		moneyA.ID = uint(i)
	}
	assert.True(t, isSamePayment(*paymentA, *paymentB))
}

func Test_isSamePayment_Returns_False_When_Given_Two_Unlike_Payments(t *testing.T) {
	paymentA, paymentB := schema.ValidPayment(), schema.ValidPayment()
	assert.False(t, isSamePayment(*paymentA, *paymentB))
}
