package main

import (
	"BigDemo/rabbitMQ/conf"
	"BigDemo/rabbitMQ/helloworld/utils"
	"context"
	"flag"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

var (
	configFile = flag.String("f", "conf.yaml", "The Config File")
)

func main() {
	flag.Parse()

	c, err := conf.Init(*configFile)
	if err != nil {
		log.Panicln(err)
	}
	//连接到 RabbitMQ 服务器
	conn, err := amqp.Dial(c.Rabbit.Url)
	utils.FailOnError(err, "RabbitMQ连接失败")
	defer conn.Close()
	/*

	 */
	//创建信道
	ch, err := conn.Channel()
	utils.FailOnError(err, "打开通道失败")
	defer ch.Close()
	/*



	 */
	//声明一个队列发送消息
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "声明队列失败")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

}
