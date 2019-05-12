package testhelpers

import (
	"fmt"
	"os"
	"testing"
)

// FullStackTest skips the test if env var FULL_STACK_TEST != true
func FullStackTest(t *testing.T) {
	t.Helper()

	if os.Getenv("FULL_STACK_TEST") != "true" {
		t.Skip("skipping test; $FULL_STACK_TEST not true")
	}
}

// APIAddress returns http address of the API under test, populated by env vars API_HOST and API_PORT
func APIAddress(t *testing.T) string {
	t.Helper()

	host, port := os.Getenv("API_HOST"), os.Getenv("API_PORT")
	return fmt.Sprintf("http://%s:%s", host, port)
}
