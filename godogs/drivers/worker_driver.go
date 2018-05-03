package drivers

import (
	"net/url"

	"github.com/gorilla/websocket"
)

// Worker function
func Worker(messageType int, message string) (string, error) {
	wu := url.URL{Scheme: "ws", Host: addr, Path: "/manager"}
	wsconn, _, err := websocket.DefaultDialer.Dial(wu.String(), nil)
	if err != nil {
		return "", err
	}
	defer wsconn.Close()

	if err := wsconn.WriteMessage(1, []byte(message)); err != nil {
		log.Println(err)
		return "", err
	}
	_, responseMessage, err := wsconn.ReadMessage()
	if err != nil {
		return "", err
	}
	return string(responseMessage), nil
}
