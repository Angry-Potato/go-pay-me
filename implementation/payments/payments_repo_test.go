package payments

import (
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/stretchr/testify/assert"
)

func Test_All_Returns_List_Of_Payments(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	allPayments, err := All(DB)
	assert.Nil(t, err)
	assert.IsType(t, []Payment{}, allPayments)
}

func Test_Create_Returns_Created_Payment_When_Payment_Valid(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	createdPayment, err := Create(DB, incomingPayment)
	assert.Nil(t, err)
	assert.Equal(t, incomingPayment, createdPayment)
}
