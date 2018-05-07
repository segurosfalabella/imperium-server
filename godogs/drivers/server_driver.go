package drivers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
	"github.com/segurosfalabella/imperium-server/manager"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var server *http.Server

const addr = "127.0.0.1:7700"

// WsUpgrader structure
type WsUpgrader struct {
	websocket.Upgrader
}

// Upgrade method
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

// CloseServer function
func CloseServer() {
	server.Close()
}

// RunApp function
func RunApp() {
	server = &http.Server{
		Addr:    addr,
		Handler: handler(),
	}
	go server.ListenAndServe()
}
