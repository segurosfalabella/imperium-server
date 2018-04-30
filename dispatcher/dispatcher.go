package dispatcher

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
)

// Dispatch ..
func Dispatch(conn connection.WsConn, message string) {
	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("There was an error sending task:", err)
			return
		}
		// TODO: Salir de una manera elegante.
		return
	}
}
