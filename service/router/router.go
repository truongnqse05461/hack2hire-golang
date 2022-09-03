package router

import (
	"github.com/gin-gonic/gin"
	"hack2hire-2022/service/config"
	"hack2hire-2022/service/handler"
	"hack2hire-2022/services"
	"log"
)

type Router struct {
	config *config.Config
}

func NewRouters(config *config.Config) *Router {
	return &Router{
		config: config,
	}
}

func (r *Router) InitGin() (*gin.Engine, error) {
	db, err := services.NewDB("mysql", r.config.MysqlURL)
	if err != nil {
		log.Panic("Connect DB failed", err)
	}
	sampleService := services.NewService(db)
	newHandler := handler.NewHandler(sampleService)
	engine := gin.New()
	engine.GET("/health", newHandler.Health)
	sampleGroup := engine.Group("/sample")
	{
		sampleGroup.GET("/hello", newHandler.GetMessage)
	}
	return engine, nil
}
