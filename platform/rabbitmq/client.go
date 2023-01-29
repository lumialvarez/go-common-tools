package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Client struct {
	Channel *amqp.Channel
}

func Init(url string) Client {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	log.Println("Connected to " + url)

	return Client{Channel: ch}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
