// +build integration

package web

import (
	"fmt"
	"testing"

	"github.com/Angry-Potato/go-pay-me/implementation/testhelpers"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func Test_Server_Status_Endpoint_Returns_Successfully(t *testing.T) {
	statusAddress := fmt.Sprintf("%s/.status", testhelpers.APIAddress(t))
	resp, err := resty.R().Get(statusAddress)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, resp.String())
}
