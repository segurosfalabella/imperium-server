package app_test

import (
	"flag"
	"log"
	"net/http"
	"testing"

	"github.com/segurosfalabella/imperium-server/app"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var addr = flag.String("addr", "localhost:7700", "http service address")

type mockHandler struct {
	mock.Mock
}

func (h *mockHandler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	log.Println("************")
	h.Called(pattern, handler)
}

func (h *mockHandler) ListenAndServe(addr string, handler http.Handler) error {
	log.Println("************")
	args := h.Called(addr, handler)
	return args.Error(0)
}

func TestShouldListenAndServe(t *testing.T) {
	handler := new(mockHandler)
	handler.On("HandleFunc", "/", mock.Anything).Return()
	handler.On("ListenAndServe", mock.Anything, nil).Return(nil)

	app.Start(handler)

	assert.True(t, handler.AssertCalled(t, "HandleFunc", "/", mock.Anything))
	assert.True(t, handler.AssertCalled(t, "ListenAndServe", ":7700", nil))
}
