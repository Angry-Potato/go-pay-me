package web

import (
	"fmt"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/payments"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

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
