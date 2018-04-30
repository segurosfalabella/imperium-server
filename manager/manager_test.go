package manager_test

import (
	"errors"
	"testing"

	"github.com/gorilla/websocket"
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

func TestShouldReadAMessage(t *testing.T) {
	mockWsConn := new(MockWsConn)
	mockWsConn.On("ReadMessage").Return(0, []byte("alohomora"), nil)
	mockWsConn.On("WriteMessage").Return(errors.New("Error"))

	manager.Manage(mockWsConn)
	mockWsConn.AssertNumberOfCalls(t, "ReadMessage", 1)
	mockWsConn.AssertNumberOfCalls(t, "WriteMessage", 1)

}

func TestShouldFailReadingMessage(t *testing.T) {
	mockWsConn := new(MockWsConn)
	mockWsConn.On("ReadMessage").Return(0, []byte("aaaa"), errors.New("Error"))
	manager.Manage(mockWsConn)
	mockWsConn.AssertNumberOfCalls(t, "ReadMessage", 1)
}

func TestShouldPass(t *testing.T) {
	mockWsConn := new(MockWsConn)
	mockWsConn.On("ReadMessage").Return(0, []byte("alohomora"), nil)
	mockWsConn.On("WriteMessage").Return(nil)
	manager.Manage(mockWsConn)
	mockWsConn.AssertNumberOfCalls(t, "ReadMessage", 1)
	mockWsConn.AssertNumberOfCalls(t, "WriteMessage", 1)
}

func TestShouldExecuteWorker(t *testing.T) {
	mockWsConn := new(MockWsConn)
	mockWsConn.On("WriteMessage", websocket.TextMessage, mock.Anything).Return(nil)
	mockWsConn.On("ReadMessage").Return(websocket.TextMessage, []byte("imperio"), nil).Once()
	mockWsConn.On("WriteMessage").Return(websocket.TextMessage, []byte(`{"name":"dummy","description":"dummy description","command":"exit"}`), nil)

	manager.Manage(mockWsConn)

	mockWsConn.AssertNumberOfCalls(t, "ReadMessage", 1)
	mockWsConn.AssertNumberOfCalls(t, "WriteMessage", 2)
	mockWsConn.AssertCalled(t, "Execute")
}
