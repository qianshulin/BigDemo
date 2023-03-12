package main

import (
	"BigDemo/rabbitMQ"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	mq := rabbitMQ.ConnectMQ()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer mq.Destory()
	body := bodyFrom(os.Args)
	err := mq.Channel.PublishWithContext(ctx,
		"",      // exchange
		"hello", // routing key
		false,   // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	rabbitMQ.FailOnError(err, "Failed to publish a message")
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
