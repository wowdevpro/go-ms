package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

var messageStorage []string

func main() {
	hub := newHub()
	go hub.run()

	go startConsumer(hub)

	r := mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	}).Methods("GET")
	r.HandleFunc("/messages/all", getAllStoredMessages).Methods("GET")

	headersOK := handlers.AllowedHeaders([]string{"*"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"*"})

	http.ListenAndServe(":8081", handlers.CORS(headersOK, originsOK, methodsOK)(r))
}