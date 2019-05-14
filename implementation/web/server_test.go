// +build unit

package web

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServerPort_Returns_Colon_Prefixed_Port_String(t *testing.T) {
	port := 1000
	actual := serverPort(port)
	expected := fmt.Sprintf(":%d", port)

	assert.Equal(t, expected, actual)
}
