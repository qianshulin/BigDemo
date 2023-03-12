package main

import (
	"BigDemo/rabbitMQ"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

var lock sync.Mutex

//Consumer 消费者
func main() {
	//1、连接rabbitMQ
	mq, err := amqp.Dial("amqp://guest:guest@43.139.139.207:5672/")
	rabbitMQ.FailOnError(err, "连接失败")
	defer mq.Close()

	//2、获取信道(通道)
	ch, err := mq.Channel()

	rabbitMQ.FailOnError(err, "打开通道失败")
	defer ch.Close()

	//3、(声明)生成一个队列
	//我们也在这里声明了队列。因为我们可能会在发布者之前启动消费者，所以我们希望在尝试使用队列中的消息之前确保队列存在。
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	rabbitMQ.FailOnError(err, "消息队列声明失败")
	//表明消费者身份接收数据
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	rabbitMQ.FailOnError(err, "Failed to register a consumer")

	//使用携程读取数据
	func() {
		for d := range msgs {
			log.Printf("收到一条消息: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

}
