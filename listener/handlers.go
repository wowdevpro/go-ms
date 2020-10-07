package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllStoredMessages(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	jsonMes, _ := json.Marshal(messageStorage)

	fmt.Fprint(writer, string(jsonMes))
}