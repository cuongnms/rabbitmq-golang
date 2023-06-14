package main

import (
	"fmt"
	"rabbitmq-golang/constant"
	"rabbitmq-golang/lib"
	"rabbitmq-golang/router"
	"rabbitmq-golang/utils"
)

func main() {
	connection, channel := lib.SetupRabbitMQConnectionChannel()
	defer connection.Close()
	defer channel.Close()

	requestQueue, err := channel.QueueDeclare(
		constant.QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)

	utils.FailOnError(err, "Failed to regist a queue")

	// Create consumer to consume message from queue
	request, err := channel.Consume(
		requestQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range request {
			switch d.Type {
			case constant.TASK1:
				fmt.Println("In progress task 1")
				fmt.Println(d.Body)
				d.Ack(false)
			case constant.TASK2:
				fmt.Println("In progress task 2")
				d.Ack(false)
			}

		}
	}()
	router.SetupRoutes()
}
