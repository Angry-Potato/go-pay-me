// +build unit

package db

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	url      = ""
	host     = "some-host"
	port     = "5432"
	user     = "some-user"
	password = "super-pass"
	database = "bucket"
)

func Test_ConnectionString_Returns_String_Containing_Host(t *testing.T) {
	actual := connectionString(url, host, port, user, password, database)

	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("(^| )host=%s($| )", host)), actual)
}

func Test_ConnectionString_Returns_String_Containing_Port(t *testing.T) {
	actual := connectionString(url, host, port, user, password, database)

	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("(^| )port=%s($| )", port)), actual)
}

func Test_ConnectionString_Returns_String_Containing_User(t *testing.T) {
	actual := connectionString(url, host, port, user, password, database)

	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("(^| )user=%s($| )", user)), actual)
}

func Test_ConnectionString_Returns_String_Containing_Password(t *testing.T) {
	actual := connectionString(url, host, port, user, password, database)

	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("(^| )password=%s($| )", password)), actual)
}

func Test_ConnectionString_Returns_String_Containing_Database(t *testing.T) {
	actual := connectionString(url, host, port, user, password, database)

	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("(^| )dbname=%s($| )", database)), actual)
}

func Test_ConnectionString_Returns_URL_If_Url_Not_Empty(t *testing.T) {
	nonEmptyURL := "some-db"
	actual := connectionString(nonEmptyURL, host, port, user, password, database)

	assert.Equal(t, nonEmptyURL, actual)
}
