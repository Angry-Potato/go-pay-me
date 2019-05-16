// +build acceptance

package web

import (
	"fmt"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/schema"
	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func validPayment() *schema.Payment {
	ID := uuid.New().String()
	return &schema.Payment{
		ID:             ID,
		Type:           "Payment",
		Version:        0,
		OrganisationID: uuid.New().String(),
		Attributes: schema.PaymentAttributes{
			Amount:               "30.20",
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

// The execution order of the following tests is significant
func Test_HTTP_Restful_Correctness(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))

	t.Run("When no payments exist", func(t *testing.T) {

		deleteAll(t, address)

		t.Run("GET payments returns empty list", func(t *testing.T) {
			allPayments := []schema.Payment{}
			resp, err := resty.R().SetResult(&allPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Empty(t, allPayments)
		})

		t.Run("DELETE payments is successful", func(t *testing.T) {
			resp, err := resty.R().Delete(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
		})

		t.Run("PUT payments is successful", func(t *testing.T) {
			allPayments := []schema.Payment{
				*validPayment(),
				*validPayment(),
				*validPayment(),
				*validPayment(),
			}
			allNewPayments := []schema.Payment{}
			resp, err := resty.R().SetResult(&allNewPayments).SetBody(allPayments).Put(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, len(allPayments), len(allNewPayments))
			actualPayments := []schema.Payment{}
			resp, err = resty.R().SetResult(&actualPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, actualPayments, allNewPayments)
		})
	})

	t.Run("When payments exist", func(t *testing.T) {

		futurePayment := *validPayment()

		deleteAll(t, address)
		initialPayments := createAll(t, address, []schema.Payment{
			*validPayment(),
			*validPayment(),
			*validPayment(),
			*validPayment(),
		})

		t.Run("GET payments returns all payments", func(t *testing.T) {
			allPayments := []schema.Payment{}
			resp, err := resty.R().SetResult(&allPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, initialPayments, allPayments)
		})

		t.Run("PUT same payments is successful", func(t *testing.T) {
			allNewPayments := []schema.Payment{}
			resp, err := resty.R().SetResult(&allNewPayments).SetBody(initialPayments).Put(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, initialPayments, allNewPayments)
		})

		t.Run("PUT new payments is successful", func(t *testing.T) {
			newPayments := []schema.Payment{
				*validPayment(),
				*validPayment(),
				*validPayment(),
				*validPayment(),
			}
			allNewPayments := []schema.Payment{}
			resp, err := resty.R().SetResult(&allNewPayments).SetBody(newPayments).Put(address)
			assert.Nil(t, err)
			assert.NotEqual(t, initialPayments, newPayments)
			assert.NotEqual(t, initialPayments, allNewPayments)
			actualPayments := []schema.Payment{}
			resp, err = resty.R().SetResult(&actualPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, actualPayments, allNewPayments)
		})

		t.Run("GET payment by ID is unsuccessful for non existent payment", func(t *testing.T) {
			resp, err := resty.R().Get(fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), futurePayment.ID))
			assert.Nil(t, err)
			assert.Equal(t, 404, resp.StatusCode())
		})

		t.Run("PUT payment by ID is unsuccessful for non existent payment", func(t *testing.T) {
			resp, err := resty.R().SetBody(validPayment()).Put(fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), futurePayment.ID))
			assert.Nil(t, err)
			assert.Equal(t, 404, resp.StatusCode())
		})

		t.Run("POST payment is successful", func(t *testing.T) {
			createdPayment := &schema.Payment{}
			resp, err := resty.R().SetBody(futurePayment).SetResult(&createdPayment).Post(address)
			assert.Nil(t, err)
			assert.Equal(t, 201, resp.StatusCode())
			futurePayment.Attributes.InternalPaymentID = futurePayment.ID
			futurePayment.Attributes.ID = createdPayment.Attributes.ID
			assert.Equal(t, futurePayment, *createdPayment)
			allPayments := []schema.Payment{}
			resp, err = resty.R().SetResult(&allPayments).Get(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Contains(t, allPayments, *createdPayment)
		})

		t.Run("GET payment by ID is successful for existing payment", func(t *testing.T) {
			foundPayment := schema.Payment{}
			resp, err := resty.R().SetResult(&foundPayment).Get(fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), futurePayment.ID))
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, futurePayment, foundPayment)
		})

		t.Run("PUT payment by ID is successful for existing payment", func(t *testing.T) {
			newVersion := int64(222)
			futurePayment.Version = newVersion
			updatedPayment := &schema.Payment{}
			resp, err := resty.R().SetBody(futurePayment).SetResult(updatedPayment).Put(fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), futurePayment.ID))
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, futurePayment, *updatedPayment)
			foundPayment := schema.Payment{}
			resp, err = resty.R().SetResult(&foundPayment).Get(fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), futurePayment.ID))
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			assert.Equal(t, newVersion, foundPayment.Version)
		})

		t.Run("DELETE payments is successful", func(t *testing.T) {
			resp, err := resty.R().Delete(address)
			assert.Nil(t, err)
			assert.Equal(t, 200, resp.StatusCode())
			allPayments := []schema.Payment{}
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

func createAll(t *testing.T, address string, allPayments []schema.Payment) []schema.Payment {
	t.Helper()
	createdPayments := []schema.Payment{}
	resp, err := resty.R().SetResult(&createdPayments).SetBody(allPayments).Put(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, len(allPayments), len(createdPayments))
	return createdPayments
}
