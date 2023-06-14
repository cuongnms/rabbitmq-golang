package lib

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"rabbitmq-golang/constant"
	"rabbitmq-golang/utils"
)

var RabbitChannel *amqp091.Channel
var rabbitConn *amqp091.Connection

func SetupRabbitMQConnectionChannel() (*amqp091.Connection, *amqp091.Channel) {

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", constant.USERNAME, constant.PASSWORD, constant.HOST, constant.PORT)

	conn, err := amqp091.Dial(url)

	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	utils.FailOnError(err, "Failed to open a channel")

	RabbitChannel = ch
	return rabbitConn, RabbitChannel

}
