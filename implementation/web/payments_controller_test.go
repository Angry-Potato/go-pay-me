package web

import (
	"fmt"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/payments"
	"github.com/google/uuid"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
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

func Test_All_Payments_Returns_Successfully(t *testing.T) {
	testhelpers.FullStackTest(t)
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	allPayments := []payments.Payment{}
	resp, err := resty.R().SetResult(&allPayments).Get(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
	assert.IsType(t, []payments.Payment{}, allPayments)
}

func Test_Create_Payment_Returns_Successfully(t *testing.T) {
	testhelpers.FullStackTest(t)
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	paymentToCreate := validPayment()
	createdPayment := &payments.Payment{}
	resp, err := resty.R().SetBody(*paymentToCreate).SetResult(&createdPayment).Post(address)
	assert.Nil(t, err)
	assert.Equal(t, 201, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
	assert.Equal(t, paymentToCreate, createdPayment)
}

func Test_Create_Payment_Returns_Failure_For_Invalid_JSON(t *testing.T) {
	testhelpers.FullStackTest(t)
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	resp, err := resty.R().SetBody("really not json at all").SetHeader("Content-Type", "application/json").Post(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
}

func Test_Create_Payment_Returns_Failure_For_Invalid_Payment(t *testing.T) {
	testhelpers.FullStackTest(t)
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	paymentToCreate := validPayment()
	paymentToCreate.Version = -1
	resp, err := resty.R().SetBody(*paymentToCreate).Post(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
}

func Test_Create_Payment_Returns_Failure_For_Creating_Existing_Payment(t *testing.T) {
	testhelpers.FullStackTest(t)
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	paymentToCreate := validPayment()
	resty.R().SetBody(*paymentToCreate).Post(address)
	resp, err := resty.R().SetBody(*paymentToCreate).Post(address)
	assert.Nil(t, err)
	assert.Equal(t, 500, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
}

func Test_DeleteAll_Payments_Returns_Successfully(t *testing.T) {
	testhelpers.FullStackTest(t)
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	paymentToCreate := validPayment()
	resty.R().SetBody(*paymentToCreate).Post(address)
	resp, err := resty.R().Delete(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
