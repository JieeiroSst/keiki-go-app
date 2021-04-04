package consumer

import (
	"github.com/JIeeiroSst/go-app/messages/rabbitmq"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Config struct {
	config rabbitmq.Config

}

func (c *Config) Consumer(message string){
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

	msgs,err:=ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
		)
	if err!=nil {
		log.Error("Failed to register consumer:",err)
	}
	forever:=make(chan bool)
	go func() {
		for d:=range msgs {
			log.Println("Received a message:",d.Body)
			d.Ack(false)
		}
	}()
	<-forever
}