package main

import (
	"fmt"
	"net/http"
)

func saveNewMessage(writer http.ResponseWriter, request *http.Request) {
	message := request.FormValue("message")
	token := request.FormValue("token")

	ok, login := checkToken(token)
	if !ok {
		http.Error(writer, "", http.StatusUnauthorized)
		return
	}

	if err := addMessageToNSQ(login, message); err != nil {
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(writer, "OK")
}
