package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"rmq/controller"
	"rmq/controller/model"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type user struct {
	User string `json:"username"`
	Level int	`json:"level"`
	Point int	`json:"point"`
	CurrentLevel int `json:"current_level"`
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	fmt.Println("connected")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	defer conn.Close()
	mongoController := controller.UserController{}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	u := &model.Model{}

	go func() {
		for d := range msgs {

			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body,u)
			er := mongoController.SetCollection(context.Background(), "quiz", "userinfo")
			failOnError(er,"SetCollectionError")
			mongoController.Put(context.Background(),*u)
			failOnError(err,"unmarshal failed")
			fmt.Println(*u)
		}
	}()


	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
