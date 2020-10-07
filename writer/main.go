package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/message/new", saveNewMessage).Methods("POST")

	headersOK := handlers.AllowedHeaders([]string{"*"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"*"})

	http.ListenAndServe(":8082", handlers.CORS(headersOK, originsOK, methodsOK)(r))
}