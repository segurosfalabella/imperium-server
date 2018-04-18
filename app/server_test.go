package app_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func getRandomPort() string {
	return fmt.Sprintf("127.0.0.1:%d", rand.Intn(20000)+10000)
}

func TestRandomPort(t *testing.T) {
	addr := getRandomPort()

	if addr == "" {
		t.Error("Invalid port", addr)
	}
}
