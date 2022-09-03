package main

import (
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"hack2hire-2022/service/config"
	"hack2hire-2022/service/router"
	"log"
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panic("Failed to init log zap", err)
	}

	logger = logger.WithOptions(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Panic("Failed to get config", err)
	}

	router := router.NewRouters(config)

	engine, err := router.InitGin()
	if err != nil {
		log.Panic("Failed to init gin", err)
	}

	_ = engine.Run(":" + config.ListenPort)
}
