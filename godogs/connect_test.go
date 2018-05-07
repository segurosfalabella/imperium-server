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

func managerHandler(w http.ResponseWriter, r *http.Request) {}

func aServer() error {
	drivers.RunApp()
	return nil
}

func workerTryToConnectSendingAMessage() error {
	Worker = new(drivers.Worker)
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
	s.Step(`^a server$`, aServer)
	s.Step(`^worker try to connect sending "([^"]*)" message$`, workerTryToConnectSendingAMessage)
	s.Step(`^server should respond "([^"]*)" message$`, serverShouldRespondMessage)
}
