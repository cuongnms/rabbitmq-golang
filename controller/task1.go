package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"net/http"
	"rabbitmq-golang/constant"
	"rabbitmq-golang/lib"
	"rabbitmq-golang/utils"
)

func Task1(c *gin.Context) {
	ch := lib.RabbitChannel
	err := ch.PublishWithContext(c, "", constant.QUEUE, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte{1, 2, 3, 4, 5},
		Type:        constant.TASK1,
	})

	utils.FailOnError(err, "Failed to publish a message")

	c.JSON(http.StatusOK, gin.H{
		"message": "Task-1 received successfully",
	})
}
