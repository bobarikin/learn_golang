package main

import (
	"lear/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Println("server is listening on 0.0.0.0:1234")
	log.Fatalln(server.Serve(listener))
}