package main

import (
	"BigDemo/rabbitMQ"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

//Producer 生产者
func main() {

	//1、连接RabbitMQ
	mq, err := amqp.Dial("amqp://guest:guest@43.139.139.207:5672/")
	rabbitMQ.FailOnError(err, "连接失败")
	defer mq.Close()

	//2、连接信道
	channel, err := mq.Channel()
	rabbitMQ.FailOnError(err, "打开通道失败")
	defer channel.Close()

	//3、声明队列
	q, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	rabbitMQ.FailOnError(err, "声明队列失败")
	//4、发送消息,消息内容要转为二进制发送
	//context.WithTimeout()是设置了一个超时函数
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := "Hello World!"
	err = channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	rabbitMQ.FailOnError(err, "消息发送失败")
	log.Printf(" [x] Sent %s\n", body)

}
