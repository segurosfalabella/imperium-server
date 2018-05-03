package drivers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
	"github.com/segurosfalabella/imperium-worker/executer"
	"github.com/segurosfalabella/imperium-worker/receiver"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type WsUpgrader struct {
	websocket.Upgrader
}

// Upgrade ...
func (u *WsUpgrader) Upgrade(w http.ResponseWriter, r *http.Request) (connection.WsConn, error) {
	return u.Upgrader.Upgrade(w, r, nil)
}

const addr = "127.0.0.1:7700"

// RunApp function
func RunApp() {
}
