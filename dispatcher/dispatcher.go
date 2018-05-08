package dispatcher

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
)

// Job struct
type Job struct {
	ID       string
	Command  string
	Response string
	ExitCode int
}

// ToJSON method
func (job *Job) ToJSON() string {
	binary, _ := json.Marshal(job)
	return string(binary)
}

// Dispatch ..
func Dispatch(conn connection.WsConn, message string) {
	for {
		job := Job{
			ID:      uuid.New().String(),
			Command: "health",
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(job.ToJSON()))
		if err != nil {
			log.Println("There was an error sending task:", err)
			return
		}
		// TODO: Salir de una manera elegante.
		return
	}
}
