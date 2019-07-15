package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/structs"

	"github.com/streadway/amqp"
)

func SendMessage(news structs.News) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	FailOnError(err, "Failed to connect AMQP")
	defer conn.Close()
	ch, err := conn.Channel()
	FailOnError(err, "Failed to connect Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"news", // queue name
		true,   // durable
		false,  // delete when used
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	FailOnError(err, "Failed to declare a queue")
	msg, err := json.Marshal(news)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(msg))

	err = ch.Publish(
		"notifExchange",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(string(msg)),
		})
	log.Printf(" [x] Sent %s", msg)
	FailOnError(err, "Failed to publish a message")
}
