package main

import (
	"github.com/bitly/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func startConsumer(hub *Hub) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("messages", "ch", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		hub.broadcast <-message.Body

		return nil
	}))


	var url string
	url = os.Getenv("NSQ_URL")
	if len(url) == 0 {
		url = ":4150"
	}

	err = consumer.ConnectToNSQD(url)
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumer.Stop()
}
