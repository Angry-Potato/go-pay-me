// +build unit

package schema

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_isUUID_Returns_False_For_Invalid_UUID(t *testing.T) {
	assert.False(t, isUUID("no uuid"))
	assert.False(t, isUUID(""))
}

func Test_isUUID_Returns_True_For_Valid_UUID(t *testing.T) {
	assert.True(t, isUUID(uuid.New().String()))
	assert.True(t, isUUID(uuid.New().String()))
}

func Test_contains_Returns_False_When_Needle_Not_In_Haystack(t *testing.T) {
	assert.False(t, contains([]string{}, ""))
	assert.False(t, contains([]string{}, "hey"))
	assert.False(t, contains([]string{"look", "out"}, "no"))
}

func Test_contains_Returns_True_When_Needle_In_Haystack(t *testing.T) {
	assert.True(t, contains([]string{"look", "look", "out"}, "look"))
	assert.True(t, contains([]string{"look"}, "look"))
	assert.True(t, contains([]string{"look", "at", "that"}, "at"))
}

func Test_isAmount_Returns_False_For_Invalid_Amount(t *testing.T) {
	assert.False(t, isAmount("no uuid"))
	assert.False(t, isAmount(""))
	assert.False(t, isAmount("1 1 1 "))
	assert.False(t, isAmount("99.0a"))
	assert.False(t, isAmount("99."))
	assert.False(t, isAmount(".99"))
	assert.False(t, isAmount("0.0"))
	assert.False(t, isAmount("1.011"))
}

func Test_isAmount_Returns_True_For_Valid_Amount(t *testing.T) {
	assert.True(t, isAmount("0"))
	assert.True(t, isAmount("0.00"))
	assert.True(t, isAmount("10.01"))
}
