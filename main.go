package main

import (
	"github.com/JIeeiroSst/go-app/messages/rabbitmq"
	"github.com/JIeeiroSst/go-app/repositories/mysql"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	config WebConfig
)

type WebConfig struct {
	Key string   					`envconfig:"WEB_KEY"`
	MysqlConfig mysql.Config 		`envconfig:"WEB_MYSQL"`
	RabbitConfig rabbitmq.Config	`envconfig:"WEB_Rabbit"`
}

func init() {
	err := envconfig.Process("", &config)
	if err!=nil {
		log.Error("Failed to open a envconfig")
	}
}

func main(){

}