package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rabbitmq-golang/constant"
	"rabbitmq-golang/controller"
)

func SetupRoutes() {
	httpRouter := gin.Default()

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Golang rabbit mq API running"})
	})

	httpRouter.POST("task1", controller.Task1)
	httpRouter.POST("task2", controller.Task2)

	httpRouter.Run(constant.SERVERPORT)
}
