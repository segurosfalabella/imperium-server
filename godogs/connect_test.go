package main

import (
	"net/http"

	"github.com/DATA-DOG/godog"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var Server *http.Server

type message struct {
	value string
}

var receiveMessages []message

const addr = "127.0.0.1:7700"

func startServer(server *http.Server) {
	http.HandleFunc("/manager", managerHandler)
	go http.ListenAndServe(addr, nil)
}

func managerHandler(w http.ResponseWriter, r *http.Request) {

}

func aWorker() error {

	return godog.ErrPending
}

func workerTryToConnectToServer() error {
	return godog.ErrPending
}

func sendMessage(arg1 string) error {
	return godog.ErrPending
}

func serverShouldRespondMessage(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a worker$`, aWorker)
	s.Step(`^worker try to connect to server$`, workerTryToConnectToServer)
	s.Step(`^send "([^"]*)" message$`, sendMessage)
	s.Step(`^server should respond "([^"]*)" message$`, serverShouldRespondMessage)
}
