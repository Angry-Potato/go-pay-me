package payments

import (
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/stretchr/testify/assert"
)

func Test_All_Returns_Empty_List_When_No_Payments_Exist(t *testing.T) {
	testhelpers.FullStackTest(t)
	DB := testhelpers.DBConnection(t, &Payment{})
	allPayments, err := All(DB)
	assert.Nil(t, err)
	assert.Equal(t, []Payment{}, allPayments)
}
