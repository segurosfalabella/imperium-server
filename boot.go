package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
	"github.com/segurosfalabella/imperium-server/manager"
)

// WsUpgrader ...
type WsUpgrader struct {
	websocket.Upgrader
}

// Upgrade ...
func (u *WsUpgrader) Upgrade(w http.ResponseWriter, r *http.Request) (connection.WsConn, error) {
	return u.Upgrader.Upgrade(w, r, nil)
}

func managerHandler(w http.ResponseWriter, r *http.Request) {
	wsUpgrader := new(WsUpgrader)
	conn, err := connection.Create(w, r, wsUpgrader)
	if err != nil {
		log.Print("HTTP request did't upgrade to WS because:", err)
		return
	}
	manager.Manage(conn)
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/manager", managerHandler)
	return r
}

func main() {
	log.Println("I'm listening on url ws://localhost:7700/manager")
	err := http.ListenAndServe(":7700", handler())
	if err != nil {
		log.Fatal(err)
	}
}
