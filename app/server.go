package app

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

const defaultPort = 8001

// Get port server, default is 8001
func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("IMPERIUM_SERVER_PORT")); err == nil {
		return port
	}
	return defaultPort
}

// Start Server
func Start() {
	log.Println("Starting Imperium Server at port", getPort())
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(getPort()), nil))
}
