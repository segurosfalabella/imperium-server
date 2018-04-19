package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetDefaultServerPort(t *testing.T) {
	assert.Equal(t, 8001, getPort())
}

func TestShouldGetEnvironmentServerPort(t *testing.T) {
	os.Setenv("IMPERIUM_SERVER_PORT", "8002")
	assert.Equal(t, 8002, getPort())
}

func TestShouldGetDefaultServerPortWhenInvalidPortValueIsSet(t *testing.T) {
	os.Setenv("IMPERIUM_SERVER_PORT", "abc")
	assert.Equal(t, 8001, getPort())
}
