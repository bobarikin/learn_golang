package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, server!\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Server has been started on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}