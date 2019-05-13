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

func Test_Create_Returns_Error_When_Creating_Existing_Payment(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	Create(DB, incomingPayment)
	createdPayment, err := Create(DB, incomingPayment)
	assert.NotNil(t, err)
	assert.Nil(t, createdPayment)
	_, ok := err.(*ValidationError)
	assert.False(t, ok)
}

func Test_Create_Returns_ValidationError_When_Payment_Invalid(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	incomingPayment.ID = ""
	incomingPayment.Version = -1
	createdPayment, err := Create(DB, incomingPayment)
	assert.NotNil(t, err)
	assert.Nil(t, createdPayment)
	_, ok := err.(*ValidationError)
	assert.True(t, ok)
}

func Test_DeleteAll_Deletes_All_Payments(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	Create(DB, validPayment())
	Create(DB, validPayment())
	err := DeleteAll(DB)
	assert.Nil(t, err)
}

func Test_DeleteAll_Returns_No_Error_If_No_Prior_Payments_Exist(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	DeleteAll(DB)
	err := DeleteAll(DB)
	assert.Nil(t, err)
}

func Test_SetAll_Returns_Inserted_Payments(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	allPayments := []*Payment{
		validPayment(),
		validPayment(),
		validPayment(),
		validPayment(),
	}
	newPayments, err := SetAll(DB, allPayments)
	assert.Nil(t, err)
	dereferenced := []Payment{}
	for _, payment := range allPayments {
		dereferenced = append(dereferenced, *payment)
	}
	assert.ElementsMatch(t, dereferenced, newPayments)
}

func Test_SetAll_Returns_ValidationError_If_Any_Payments_Are_Invalid(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	badEgg := validPayment()
	badEgg.ID = "not a real uuid"
	badEgg.Version = -1
	allPayments := []*Payment{
		validPayment(),
		validPayment(),
		validPayment(),
		validPayment(),
		badEgg,
	}
	newPayments, err := SetAll(DB, allPayments)
	assert.NotNil(t, err)
	assert.Empty(t, newPayments)
	_, ok := err.(*ValidationError)
	assert.True(t, ok)
	assert.Contains(t, err.Error(), badEgg.ID)
}
