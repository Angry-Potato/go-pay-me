package main

import (
	"strconv"
	"testing"
)

func TestDefaultPort(t *testing.T) {
	defaultPort := 1000
	got := port("", defaultPort)
	want := defaultPort

	if got != want {
		t.Errorf("port(\"\", %d) = %d; want %d", defaultPort, got, want)
	}
}

func TestEnvironmentPort(t *testing.T) {
	envPort, defaultPort := "8080", 1000
	got := port(envPort, defaultPort)
	want, _ := strconv.Atoi(envPort)

	if got != want {
		t.Errorf("port(%s, %d) = %d; want %d", envPort, defaultPort, got, want)
	}
}
