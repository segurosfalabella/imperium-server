package manager

import (
	"errors"

	"github.com/segurosfalabella/imperium-server/dispatcher"

	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
)

var authToken = "alohomora"
var authTokenResponse = "imperio"

// Manage ...
func Manage(conn connection.WsConn) {
	err := auth(conn)
	if err == nil {
		dispatcher.Dispatch(conn, "hola")
	}
}

func auth(conn connection.WsConn) error {
	_, message, _ := conn.ReadMessage()
	if string(message) != authToken {
		return errors.New("received a dark wizard")
	}
	err := conn.WriteMessage(websocket.TextMessage, []byte(authTokenResponse))
	if err != nil {
		return err
	}
	return nil
}
