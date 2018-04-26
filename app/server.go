package app

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

// Start Server
func Start() {
	err := http.ListenAndServe(":7700", handler())
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/manager", managerHandler)
	return r
}

func managerHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("HTTP request did't upgrade to WS because:", err)
		return
	}
	log.Println("Server is listening...")

	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("There was an error reading the message:", err)
			return
		}
		log.Printf("recv: %s", message)
		// err = conn.WriteMessage(mt, []byte(message))
		// if err != nil {
		// 	log.Println("There was an error writing the message::", err)
		// 	return
		// }
	}
}
