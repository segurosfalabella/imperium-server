package dispatcher_test

import (
	"errors"
	"testing"

	"github.com/segurosfalabella/imperium-server/dispatcher"
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
	returnArgs := conn.Called()
	return returnArgs.Error(0)
}
func TestDispatch(t *testing.T) {
	mockWsConn := new(MockWsConn)
	mockWsConn.On("WriteMessage", mock.Anything, mock.Anything).Return(errors.New("Error"))
	message := "hola"
	dispatcher.Dispatch(mockWsConn, message)
	mockWsConn.AssertNumberOfCalls(t, "WriteMessage", 1)
}

func TestDispatchAMessage(t *testing.T) {
	mockWsConn := new(MockWsConn)
	mockWsConn.On("WriteMessage", mock.Anything, mock.Anything).Return(nil)
	message := "hola"
	dispatcher.Dispatch(mockWsConn, message)
	mockWsConn.AssertNumberOfCalls(t, "WriteMessage", 1)
}
