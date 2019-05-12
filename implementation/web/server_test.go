package web

import (
	"fmt"
	"testing"
)

func TestPort(t *testing.T) {
	port := 1000
	got := serverPort(port)
	want := fmt.Sprintf(":%d", port)

	if got != want {
		t.Errorf("serverPort(%d) = %s; want %s", port, got, want)
	}
}
