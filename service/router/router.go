package router

import (
	"hack2hire-2022/service/config"
	"hack2hire-2022/service/handler"

	"hack2hire-2022/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Router struct {
	config     *config.Config
	writer     *kafka.Writer
	kafkaTopic string
}

func NewRouters(config *config.Config, writer *kafka.Writer, topic string) *Router {
	return &Router{
		config:     config,
		writer:     writer,
		kafkaTopic: topic,
	}
}

func (r *Router) InitGin() (*gin.Engine, error) {
	db, err := services.NewDB("mysql", r.config.MysqlURL)
	if err != nil {
		log.Panic("Connect DB failed", err)
	}
	sampleService := services.NewService(db, r.writer, r.kafkaTopic)
	newHandler := handler.NewHandler(sampleService)
	engine := gin.New()
	engine.GET("/health", newHandler.Health)
	sampleGroup := engine.Group("/sample")
	{
		sampleGroup.GET("/hello/:id", newHandler.GetMessage)
		sampleGroup.POST("/bookings", newHandler.SaveBookings)
	}
	return engine, nil
}
