// +build integration

package web

import (
	"fmt"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/schema"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func Test_Get_Payments_Returns_Successfully(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	allPayments := []schema.Payment{}
	resp, err := resty.R().SetResult(&allPayments).Get(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
	assert.IsType(t, []schema.Payment{}, allPayments)
}

func Test_Post_Payment_Returns_Successfully(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	paymentToCreate := schema.ValidPayment()
	createdPayment := &schema.Payment{}
	resp, err := resty.R().SetBody(*paymentToCreate).SetResult(&createdPayment).Post(address)
	assert.Nil(t, err)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, paymentToCreate, createdPayment)
}

func Test_Post_Payment_Returns_Failure_For_Invalid_JSON(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	resp, err := resty.R().SetBody("really not json at all").SetHeader("Content-Type", "application/json").Post(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
}

func Test_Post_Payment_Returns_Failure_For_Invalid_Payment(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	paymentToCreate := schema.ValidPayment()
	paymentToCreate.Version = -1
	resp, err := resty.R().SetBody(*paymentToCreate).Post(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
}

func Test_Put_Payments_Returns_Successfully(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	allPayments := []*schema.Payment{
		schema.ValidPayment(),
		schema.ValidPayment(),
		schema.ValidPayment(),
		schema.ValidPayment(),
	}
	allNewPayments := []schema.Payment{}
	resp, err := resty.R().SetResult(&allNewPayments).SetBody(allPayments).Put(address)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
	assert.IsType(t, []schema.Payment{}, allNewPayments)
}

func Test_Put_Payments_Returns_Failure_For_Single_Bad_Egg(t *testing.T) {
	address := fmt.Sprintf("%s/payments", testhelpers.APIAddress(t))
	badEgg := schema.ValidPayment()
	badEgg.ID = "oh no!"
	allPayments := []*schema.Payment{
		schema.ValidPayment(),
		schema.ValidPayment(),
		schema.ValidPayment(),
		schema.ValidPayment(),
		badEgg,
	}
	resp, err := resty.R().SetBody(allPayments).Put(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
}

func Test_Get_Payment_By_ID_Returns_Failure_For_Non_Existent_Payment(t *testing.T) {
	paymentToCreate := schema.ValidPayment()
	address := fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), paymentToCreate.ID)
	resp, err := resty.R().Get(address)
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode())
}

func Test_Get_Payment_By_ID_Returns_Failure_For_Bad_Request(t *testing.T) {
	address := fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), "ohhhhh sh*t")
	resp, err := resty.R().Get(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
}

func Test_Delete_Payment_By_ID_Returns_Failure_For_Non_Existent_Payment(t *testing.T) {
	paymentToCreate := schema.ValidPayment()
	address := fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), paymentToCreate.ID)
	resp, err := resty.R().Delete(address)
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode())
}

func Test_Delete_Payment_By_ID_Returns_Failure_For_Bad_Request(t *testing.T) {
	address := fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), "ohhhhh sh*t")
	resp, err := resty.R().Delete(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
}

func Test_Put_Payment_By_ID_Returns_Failure_For_Non_Existent_Payment(t *testing.T) {
	payment := schema.ValidPayment()
	address := fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), payment.ID)
	resp, err := resty.R().SetBody(payment).Put(address)
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode())
}

func Test_Put_Payment_By_ID_Returns_Failure_For_Bad_Request(t *testing.T) {
	address := fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), "ohhhhh sh*t")
	resp, err := resty.R().Put(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
	paymentToUpdate := schema.ValidPayment()
	address = fmt.Sprintf("%s/payments/%s", testhelpers.APIAddress(t), paymentToUpdate.ID)
	paymentToUpdate.OrganisationID = "oh no no no"
	resp, err = resty.R().SetBody(paymentToUpdate).Put(address)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode())
}
