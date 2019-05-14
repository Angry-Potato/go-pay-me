// +build integration

package payments

import (
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/stretchr/testify/assert"
)

func Test_All_Returns_List_Of_Payments(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	allPayments, err := All(DB)
	assert.Nil(t, err)
	assert.IsType(t, []Payment{}, allPayments)
}

func Test_Create_Returns_Created_Payment_When_Payment_Valid(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	createdPayment, err := Create(DB, incomingPayment)
	assert.Nil(t, err)
	assert.Equal(t, incomingPayment, createdPayment)
}

func Test_Create_Returns_Error_When_Creating_Existing_Payment(t *testing.T) {
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
	DB := testhelpers.DBConnection(t, &Payment{})
	Create(DB, validPayment())
	Create(DB, validPayment())
	err := DeleteAll(DB)
	assert.Nil(t, err)
}

func Test_DeleteAll_Returns_No_Error_If_No_Prior_Payments_Exist(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	DeleteAll(DB)
	err := DeleteAll(DB)
	assert.Nil(t, err)
}

func Test_SetAll_Returns_Inserted_Payments(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	allPayments := []Payment{
		*validPayment(),
		*validPayment(),
		*validPayment(),
		*validPayment(),
	}
	newPayments, err := SetAll(DB, allPayments)
	assert.Nil(t, err)
	assert.ElementsMatch(t, allPayments, newPayments)
}

func Test_SetAll_Returns_ValidationError_If_Any_Payments_Are_Invalid(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	badEgg := validPayment()
	badEgg.ID = "not a real uuid"
	badEgg.Version = -1
	allPayments := []Payment{
		*validPayment(),
		*validPayment(),
		*validPayment(),
		*validPayment(),
		*badEgg,
	}
	newPayments, err := SetAll(DB, allPayments)
	assert.NotNil(t, err)
	assert.Empty(t, newPayments)
	_, ok := err.(*ValidationError)
	assert.True(t, ok)
	assert.Contains(t, err.Error(), badEgg.ID)
}

func Test_Get_Returns_Payment_When_Payment_Exists(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	Create(DB, incomingPayment)
	foundPayment, err := Get(DB, incomingPayment.ID)
	assert.Nil(t, err)
	assert.NotNil(t, foundPayment)
	assert.Equal(t, foundPayment, incomingPayment)
}

func Test_Get_Returns_Nil_When_Payment_Does_Not_Exist(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	foundPayment, err := Get(DB, "unused-id")
	assert.NotNil(t, err)
	assert.Nil(t, foundPayment)
}

func Test_Get_Returns_ValidationError_When_Payment_ID_Invalid(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	foundPayment, err := Get(DB, "not a uuid")
	assert.NotNil(t, err)
	assert.Nil(t, foundPayment)
	_, ok := err.(*ValidationError)
	assert.True(t, ok)
}

func Test_Delete_Returns_Nil_Error_When_Payment_Exists(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	Create(DB, incomingPayment)
	err := Delete(DB, incomingPayment.ID)
	assert.Nil(t, err)
}

func Test_Delete_Returns_Error_When_Payment_Does_Not_Exist(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	err := Delete(DB, "unused-id")
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err))
}

func Test_Delete_Returns_ValidationError_When_Payment_ID_Invalid(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	err := Delete(DB, "not a uuid")
	assert.NotNil(t, err)
	_, ok := err.(*ValidationError)
	assert.True(t, ok)
}

func Test_Update_Returns_Updated_Payment_When_Payment_Exists_And_Valid_Update(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	Create(DB, incomingPayment)
	incomingPayment.Version = 999
	newPayment, err := Update(DB, incomingPayment.ID, incomingPayment)
	assert.Nil(t, err)
	assert.Equal(t, incomingPayment, newPayment)
}

func Test_Update_Returns_Nil_When_Payment_Exists_Unchanged(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	incomingPayment := validPayment()
	Create(DB, incomingPayment)
	newPayment, err := Update(DB, incomingPayment.ID, incomingPayment)
	assert.Nil(t, err)
	assert.Nil(t, newPayment)
}

func Test_Update_Returns_Error_When_Payment_Does_Not_Exist(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	payment, err := Update(DB, "unused-id", validPayment())
	assert.NotNil(t, err)
	assert.Nil(t, payment)
	assert.True(t, gorm.IsRecordNotFoundError(err))
}

func Test_Update_Returns_ValidationError_When_Payment_ID_Invalid(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	_, err := Update(DB, "not a uuid", validPayment())
	assert.NotNil(t, err)
	_, ok := err.(*ValidationError)
	assert.True(t, ok)
}

func Test_Update_Returns_ValidationError_When_Payment_Update_Invalid(t *testing.T) {
	DB := testhelpers.DBConnection(t, &Payment{})
	payment := validPayment()
	payment.Version = -22
	newPayment, err := Update(DB, payment.ID, payment)
	assert.NotNil(t, err)
	assert.Nil(t, newPayment)
	_, ok := err.(*ValidationError)
	assert.True(t, ok)
}
