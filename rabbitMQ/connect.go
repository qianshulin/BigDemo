package rabbitMQ

import amqp "github.com/rabbitmq/amqp091-go"

//RabbitMQ 结构体
type RabbitMQ struct {
	//连接
	Conn    *amqp.Connection
	Channel *amqp.Channel
	////队列
	//QueueName string
	////交换机名称
	//ExChange string
	////绑定的key名称
	//Key string
	////连接的信息，上面已经定义好了
	//MqUrl string
}

//ConnectMQ 创建结构体实例，参数队列名称、交换机名称和bind的key（也就是几个大写的，除去定义好的常量信息）
func ConnectMQ() *RabbitMQ {
	mq, err := amqp.Dial("amqp://guest:guest@43.139.139.207:5672/")
	FailOnError(err, "连接失败")
	//2、连接信道
	channel, err := mq.Channel()
	FailOnError(err, "打开通道失败")

	return &RabbitMQ{Conn: mq, Channel: channel}
}

//Destory 关闭conn和chanel的方法
func (r *RabbitMQ) Destory() {
	r.Channel.Close()
	r.Conn.Close()
}
