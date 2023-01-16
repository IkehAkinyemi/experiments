package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect AMQP", )
	defer conn.Close()

	chann, err := conn.Channel()
	failOnError(err, "failed to create a channel")
	defer chann.Close()

	q, err := chann.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare a queue")

	fmt.Println(q.Name)

	msgs, err := chann.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to consume messages")

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
		}
	}()
	log.Printf(" [*] Waiting for incoming message. To exit press Ctrl+C")
	<-forever
}