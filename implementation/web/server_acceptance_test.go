// +build acceptance

package web

import (
	"fmt"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/payments"
	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func validPayment() *payments.Payment {
	return &payments.Payment{
		ID:             uuid.New().String(),
		Type:           "Payment",
		Version:        0,
		OrganisationID: uuid.New().String(),
	}
}
func Test_HTTP_Restful_Correctness(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))

	t.Run("When no payments exist", func(t *testing.T) {

		deleteAll(t, address)

		t.Run("Get Payments Returns Empty List", func(t *testing.T) {
			allPayments := []payments.Payment{}
			resp, err := resty.R().SetResult(&allPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Empty(t, allPayments)
		})

		t.Run("Delete Payments Returns Success", func(t *testing.T) {
			resp, err := resty.R().Delete(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
		})

		t.Run("Put Payments Returns Success", func(t *testing.T) {
			allPayments := []payments.Payment{
				*validPayment(),
				*validPayment(),
				*validPayment(),
				*validPayment(),
			}
			allNewPayments := []payments.Payment{}
			resp, err := resty.R().SetResult(&allNewPayments).SetBody(allPayments).Put(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, allPayments, allNewPayments)
		})
	})
}

func deleteAll(t *testing.T, address string) {
	resp, err := resty.R().Delete(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
