package manager_test

import (
	"errors"
	"testing"

	"github.com/segurosfalabella/imperium-server/manager"
	"github.com/stretchr/testify/mock"
)

type MockWsConn struct {
	mock.Mock
}

func (conn *MockWsConn) Close() error {
	return nil
}

func (conn *MockWsConn) ReadMessage() (messageType int, p []byte, err error) {
	returnArgs := conn.Called()
	return returnArgs.Int(0), returnArgs.Get(1).([]byte), returnArgs.Error(2)
}

func (conn *MockWsConn) WriteMessage(messageType int, data []byte) error {
	returnArgs := conn.Called()
	return returnArgs.Error(0)
}

func TestManage(t *testing.T) {
	tt := []struct {
		name       string
		message    []byte
		readCount  int
		readErr    error
		writeCount int
		writeErr   error
	}{
		{
			name:       "Should read an auth message",
			message:    []byte("alohomora"),
			readCount:  1,
			writeCount: 2,
		},
		{
			name:      "Should fail reading a message",
			message:   []byte("aaaa"),
			readCount: 1,
			readErr:   errors.New("fdasfsaddfafd"),
		},
		{
			name:      "Should fail auth",
			message:   []byte("aaaa"),
			readCount: 1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockWsConn := new(MockWsConn)
			mockWsConn.On("ReadMessage").Return(1, tc.message, tc.readErr)
			mockWsConn.On("WriteMessage").Return(tc.writeErr)

			manager.Manage(mockWsConn)

			mockWsConn.AssertNumberOfCalls(t, "ReadMessage", tc.readCount)
			mockWsConn.AssertNumberOfCalls(t, "WriteMessage", tc.writeCount)
		})
	}
}
