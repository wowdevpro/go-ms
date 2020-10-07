package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-nsq"
	"io/ioutil"
	"net/http"
	"os"
)

func addMessageToNSQ(login, message string) error {
	var url string
	url = os.Getenv("NSQ_URL")
	if len(url) == 0 {
		url = ":4150"
	}

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(url, config)
	defer w.Stop()

	params := []string{login, message}

	mes, err := json.Marshal(params)
	if err != nil {
		return err
	}

	err = w.Publish("messages", mes)
	if err != nil {
		return err
	}

	return nil
}

func checkToken(token string) (bool, string) {
	resp, err := http.Get("http://auth:8080/token/check?token=" + token)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	login, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err, resp.StatusCode)
		return false, ""
	}

	return true, string(login)
}
