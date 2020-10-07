package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func login(writer http.ResponseWriter, request *http.Request) {
	login := request.FormValue("login")
	password := request.FormValue("password")

	user, ok := users[login]
	if !ok {
		http.Error(writer, "Invalid user", http.StatusForbidden)
		return
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		http.Error(writer, "Login or password do not match", http.StatusForbidden)
		return
	}

	token, err := createToken(user.Login)
	if err != nil {
		http.Error(writer, "System error, can't create token", http.StatusInternalServerError)
		return
	}

	userTokens[token] = user.Login

	fmt.Fprint(writer, token)
}

func tokenCheck(writer http.ResponseWriter, request *http.Request) {
	token := request.URL.Query().Get("token")

	if len(token) == 0 {
		http.Error(writer, "Token param mismatch", http.StatusBadRequest)
		return
	}

	userLogin, ok := userTokens[token]

	if !ok {
		http.Error(writer, "Wrong token", http.StatusBadRequest)
		return
	}

	if !isValidToken(token) {
		http.Error(writer, "Token expired", http.StatusBadRequest)
		return
	}

	fmt.Fprint(writer, userLogin)
}
