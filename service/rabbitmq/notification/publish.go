package notification

import (
	"context"
	"encoding/json"
	"github.com/lumialvarez/go-common-tools/platform/rabbitmq"
	"github.com/lumialvarez/go-common-tools/service/rabbitmq/notification/dto"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

type Service struct {
	rabbitmq rabbitmq.Client
}

func Init() Service {
	rabbitUrl := os.Getenv("RABBITMQ_URL")
	rabbitmqClient := rabbitmq.Init(rabbitUrl)
	return Service{rabbitmq: rabbitmqClient}
}

func (service *Service) DefineNotificationQueue() (*amqp.Queue, context.Context, error) {
	q, err := service.rabbitmq.Channel.QueueDeclare(
		"NOTIFICATION_QUEUE", // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Print(err, "Failed to declare a queue")
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return &q, ctx, nil
}

func (service *Service) PublishNotification(title string, detail string) error {
	q, ctx, err := service.DefineNotificationQueue()
	if err != nil {
		return err
	}

	notification := dto.Notification{
		Title:  title,
		Detail: detail,
	}

	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		log.Print(err, "Failed to extract json from dto", err)
		return err
	}

	err = service.rabbitmq.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonNotification,
		})
	if err != nil {
		log.Print(err, "Failed to publish a message", err)
		return err
	}

	return nil
}
