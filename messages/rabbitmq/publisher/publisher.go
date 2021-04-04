package publisher

import (
	"github.com/JIeeiroSst/go-app/messages/rabbitmq"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Config struct {
	config rabbitmq.Config
}

func (c *Config)  Publisher(message string){
	conn,err:=amqp.Dial(c.config.DSN)
	if err!=nil {
		log.Error("Failed to connect to RabbitMQ :",err)
	}
	defer conn.Close()
	ch,err:=conn.Channel()
	if err!=nil{
		log.Error("Failed to open a channel :",err)
	}
	defer ch.Close()
	queue,err:=ch.QueueDeclare(
		"publisher",
		true,
		false,
		false,
		false,
		nil)
	if err!=nil{
		log.Error("Failed to declare a queue :",err)
	}
	err=ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err!=nil{
		log.Error("Failed to publish a message:",err)
	}
}

