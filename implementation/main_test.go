package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Port_Returns_DefaultPort_When_EnvPort_Is_Empty(t *testing.T) {
	defaultPort := 1000
	actual := port("", defaultPort)
	expected := defaultPort

	assert.Equal(t, expected, actual)
}

func Test_Port_Returns_EnvPort_When_EnvPort_Is_Not_Empty(t *testing.T) {
	envPort, defaultPort := "8080", 1000
	actual := port(envPort, defaultPort)
	expected, _ := strconv.Atoi(envPort)

	assert.Equal(t, expected, actual)
}
