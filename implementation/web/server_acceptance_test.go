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

		t.Run("get payments returns empty list", func(t *testing.T) {
			allPayments := []payments.Payment{}
			resp, err := resty.R().SetResult(&allPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Empty(t, allPayments)
		})

		t.Run("delete payments returns success", func(t *testing.T) {
			resp, err := resty.R().Delete(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
		})

		t.Run("put payments returns success", func(t *testing.T) {
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

	t.Run("When payments exist", func(t *testing.T) {
		initialPayments := []payments.Payment{
			*validPayment(),
			*validPayment(),
			*validPayment(),
			*validPayment(),
		}

		deleteAll(t, address)
		createAll(t, address, initialPayments)

		t.Run("get payments returns all payments", func(t *testing.T) {
			allPayments := []payments.Payment{}
			resp, err := resty.R().SetResult(&allPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, initialPayments, allPayments)
		})

		t.Run("put same payments returns success", func(t *testing.T) {
			allNewPayments := []payments.Payment{}
			resp, err := resty.R().SetResult(&allNewPayments).SetBody(initialPayments).Put(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, initialPayments, allNewPayments)
		})

		t.Run("put new payments returns success", func(t *testing.T) {
			newPayments := []payments.Payment{
				*validPayment(),
				*validPayment(),
				*validPayment(),
				*validPayment(),
			}
			allNewPayments := []payments.Payment{}
			resp, err := resty.R().SetResult(&allNewPayments).SetBody(newPayments).Put(address)
			assert.Nil(t, err)
			assert.NotEqual(t, newPayments, initialPayments)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, newPayments, allNewPayments)
		})

		t.Run("delete payments returns success", func(t *testing.T) {
			resp, err := resty.R().Delete(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			allPayments := []payments.Payment{}
			resp, err = resty.R().SetResult(&allPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Empty(t, allPayments)
		})
	})
}

func deleteAll(t *testing.T, address string) {
	t.Helper()
	resp, err := resty.R().Delete(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func createAll(t *testing.T, address string, allPayments []payments.Payment) {
	t.Helper()
	resp, err := resty.R().SetBody(allPayments).Put(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
