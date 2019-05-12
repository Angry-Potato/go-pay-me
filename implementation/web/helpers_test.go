package web

import (
	"fmt"
	"os"
	"testing"
)

func FullStackTest(t *testing.T) {
	t.Helper()

	if os.Getenv("FULL_STACK_TEST") != "true" {
		t.Skip("skipping test; $FULL_STACK_TEST not true")
	}
}

func APIAddress(t *testing.T) string {
	t.Helper()

	host, port := os.Getenv("API_HOST"), os.Getenv("API_PORT")
	return fmt.Sprintf("http://%s:%s", host, port)
}
