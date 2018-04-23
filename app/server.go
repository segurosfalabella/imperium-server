package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

const defaultPort = 7700

var upgrader = websocket.Upgrader{} // use default options

func managerHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, []byte("holi!! Amiguito"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

type handleler interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ListenAndServe(addr string, handler http.Handler) error
}

// ConfigHandler funcition
// var ConfigHandler = http.HandleFunc

// Start Server
func Start(handleler handleler) {
	handleler.HandleFunc("/", managerHandler)
	// http.HandleFunc("/", managerHandler)
	// err := http.ListenAndServe(":"+strconv.Itoa(defaultPort), nil)
	err := handleler.ListenAndServe(":"+strconv.Itoa(defaultPort), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening!!!")
}
