package main

/*


*/
import (
	"log"
	"os"
	"fmt"
	"github.com/streadway/amqp"
	"strings"
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
	log.Printf("URL: %s", url)
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := os.Getenv("MQEXCHANGE")
	log.Println(exchangeName)
	err = ch.ExchangeDeclare(
		exchangeName,   // name
		"topic", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")


	body := bodyFrom(os.Args)
	routingKey := os.Getenv("MQKEY")
	log.Println(body, routingKey)
	err = ch.Publish(
		exchangeName, // exchange
		routingKey,     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}