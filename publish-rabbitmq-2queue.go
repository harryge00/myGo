package main

import (
	"log"
	"os"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	password := os.Getenv("MQPASS")
	username := os.Getenv("MQUSER")
	host := os.Getenv("MQHOST")
	port := os.Getenv("MQPORT")
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)

	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queueName := os.Getenv("QUEUENAME")
	if queueName == "" {
		queueName = "hello-world"
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	for true {
		log.Printf(" [x] Sent 100 %s", body)
		time.Sleep(500 * time.Millisecond)
		for i := 0; i < 100; i++ {
			err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		}
	}
	
	failOnError(err, "Failed to publish a message")
}