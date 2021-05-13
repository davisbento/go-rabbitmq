package main

import (
	"encoding/json"
	"log"

	"github.com/wagslane/go-rabbitmq"
)

type Message struct {
	Name string `json:"name"`
}

func main() {
	consumer, err := rabbitmq.NewConsumer("amqp://")
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.StartConsuming(
		func(d rabbitmq.Delivery) bool {
			message := string(d.Body)
			data := Message{}
			json.Unmarshal([]byte(message), &data)
			log.Printf("converted: %v", data.Name)
			// true to ACK, false to NACK
			return true
		},
		"my_queue",
		[]string{"routing_key"},
		rabbitmq.WithConsumeOptionsBindingExchange("my_exchange"))
	if err != nil {
		log.Fatal(err)
	}

	// block main thread so consumers run forever
	forever := make(chan struct{})
	<-forever
}
