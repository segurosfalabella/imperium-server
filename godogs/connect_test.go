package main

import (
	"fmt"
	"net/http"

	"github.com/DATA-DOG/godog"
	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/godogs/drivers"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var Server *http.Server
var Worker *drivers.Worker
var workerMessage string

const addr = "127.0.0.1:7700"

func startServer(server *http.Server) {
	http.HandleFunc("/manager", managerHandler)
	go http.ListenAndServe(addr, nil)
}

func managerHandler(w http.ResponseWriter, r *http.Request) {}

func aWorker() error {
	Worker = new(drivers.Worker)
	return nil
}

func workerTryToConnectSendingAMessage() error {
	drivers.RunApp()
	workerMessage = "alohomora"
	_, err := Worker.Connect(websocket.TextMessage, workerMessage)
	return err
}

func serverShouldRespondMessage(serverResponse string) error {
	serverMessage, err := Worker.Connect(websocket.TextMessage, workerMessage)
	if err != nil {
		return err
	}
	if serverMessage != serverResponse {
		return fmt.Errorf("expected %s message but received %s", serverResponse, serverMessage)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a worker$`, aWorker)
	s.Step(`^worker try to connect sending "([^"]*)" message$`, workerTryToConnectSendingAMessage)
	s.Step(`^server should respond "([^"]*)" message$`, serverShouldRespondMessage)
}
