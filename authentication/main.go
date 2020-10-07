package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type user struct {
	Login    string
	Password []byte
}

var users = map[string]user{}
var userTokens = map[string]string{}

func init() {
	pass1, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)
	pass2, _ := bcrypt.GenerateFromPassword([]byte("321"), bcrypt.MinCost)

	users["test"] = user{"test", pass1}
	users["test2"] = user{"test2", pass2}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/token/check", tokenCheck).Methods("GET")

	headersOK := handlers.AllowedHeaders([]string{"*"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"*"})

	http.ListenAndServe(":8080", handlers.CORS(headersOK, originsOK, methodsOK)(r))
}
