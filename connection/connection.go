package connection

import (
	"net/http"
)

// WsConn interface
type WsConn interface {
	Close() error
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
}

// Upgrader...
type Upgrader interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (WsConn, error)
}

// Create method
func Create(w http.ResponseWriter, r *http.Request, upgrader Upgrader) (WsConn, error) {

	conn, err := upgrader.Upgrade(w, r)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
