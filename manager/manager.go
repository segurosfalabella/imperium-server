package manager

import (
	"errors"
	"log"

	"github.com/segurosfalabella/imperium-server/dispatcher"

	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
)

var authToken = "alohomora"
var authTokenResponse = "imperio"

// Manage ...
func Manage(conn connection.WsConn) {
	if err := auth(conn); err != nil {
		log.Println(err)
	} else {
		dispatcher.Dispatch(conn, "hola")
	}
}

func auth(conn connection.WsConn) error {
	_, message, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	if _, error := validateCredentials(message); error != nil {
		return error
	}
	return conn.WriteMessage(websocket.TextMessage, []byte(authTokenResponse))
}

func validateCredentials(message []byte) (bool, error) {
	if string(message) == authToken {
		return true, nil
	}
	return false, errors.New("Invalid Credentials")
}
