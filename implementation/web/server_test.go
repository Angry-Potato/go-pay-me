package web

import (
	"fmt"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func Test_ServerPort_Returns_Colon_Prefixed_Port_String(t *testing.T) {
	port := 1000
	actual := serverPort(port)
	expected := fmt.Sprintf(":%d", port)

	assert.Equal(t, expected, actual)
}

func Test_Server_Status_Endpoint_Returns_Successfully(t *testing.T) {
	testhelpers.FullStackTest(t)
	statusAddress := fmt.Sprintf("%s/.status", testhelpers.APIAddress(t))
	resp, err := resty.R().Get(statusAddress)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
}
