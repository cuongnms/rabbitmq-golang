package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"net/http"
	"rabbitmq-golang/constant"
	"rabbitmq-golang/lib"
	"rabbitmq-golang/utils"
)

func Task2(c *gin.Context) {
	ch := lib.RabbitChannel
	err := ch.PublishWithContext(
		c,
		"",             // exchange
		constant.QUEUE, // routing key
		false,          // mandatory
		false,          // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        nil,
			Type:        constant.TASK2,
		})

	utils.FailOnError(err, "Failed to publish a message")

	c.JSON(http.StatusOK, gin.H{
		"message": "Task-2 Received Successfully",
	})
}
