package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/structs"

	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

func ReceiveMessage() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	FailOnError(err, "Failed to connect AMQP")
	db, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	FailOnError(err, "Failed to connect DB mysql")
	defer db.Close()
	defer conn.Close()
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			msg := string(d.Body)
			news := structs.News{}
			json.Unmarshal([]byte(msg), &news)
			fmt.Println("Sending notification Add News")
			db.Create(news)
			fmt.Println("Author: " + news.Author + "\nBody: " + news.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
