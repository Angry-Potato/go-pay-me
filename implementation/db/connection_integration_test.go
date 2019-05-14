// +build integration

package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Connect(t *testing.T) {
	url, host, port, user, password, database := "", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	resp, err := Connect(url, host, port, user, password, database)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
