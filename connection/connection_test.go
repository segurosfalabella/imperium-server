package connection_test

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/segurosfalabella/imperium-server/connection"
	"github.com/stretchr/testify/mock"
)

type MockWsConn struct {
	mock.Mock
}

func (conn *MockWsConn) Close() error {
	return nil
}

func (conn *MockWsConn) ReadMessage() (messageType int, p []byte, err error) {
	return 0, nil, nil
}

func (conn *MockWsConn) WriteMessage(messageType int, data []byte) error {
	return nil
}

type MockUpgrader struct {
	mock.Mock
}

func (u *MockUpgrader) Upgrade(w http.ResponseWriter, r *http.Request) (connection.WsConn, error) {
	returnArgs := u.Called(w, r)
	log.Println(returnArgs.Get(0).(connection.WsConn))
	return returnArgs.Get(0).(connection.WsConn), returnArgs.Error(1)
}

func TestConnection(t *testing.T) {
	tt := []struct {
		name string
		err  error
	}{
		{name: "Should create a valid connection", err: nil},
		{name: "Should fail when wrong address", err: errors.New("Bad connection params")},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			address := "http://127.0.0.1:7700/manager"
			responseWriter := httptest.NewRecorder()
			mockConn := new(MockWsConn)
			upgrader := new(MockUpgrader)

			request, _ := http.NewRequest("GET", address, nil)

			upgrader.On("Upgrade", responseWriter, request).Return(mockConn, tc.err)

			_, err := connection.Create(responseWriter, request, upgrader)

			if err != tc.err {
				t.Errorf("Can't create a connection: %v", tc.err)
			}
		})
	}
}
