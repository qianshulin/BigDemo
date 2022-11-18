package main

import (
	"BigDemo/rabbitMQ/helloworld/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	app := iris.New()
	app.Get("/", Consumer)
	app.Listen(":8080")
}

//Consumer 消费者
func Consumer(context.Context) {
	//连接rabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@43.139.139.207:5672/")
	utils.FailOnError(err, "RabbitMQ连接失败")
	defer conn.Close()
	//
	ch, err := conn.Channel()
	utils.FailOnError(err, "打开通道失败")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "消息队列声明失败")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
