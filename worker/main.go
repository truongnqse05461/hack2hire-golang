package main

import (
	"context"
	"hack2hire-2022/queue"
	"hack2hire-2022/service/config"
	"hack2hire-2022/services"
	workerCfg "hack2hire-2022/worker/config"
	workerSvc "hack2hire-2022/worker/service"
	"log"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Panic("Failed to get config", err)
	}

	kafkaCfg, err := workerCfg.Load()
	if err != nil {
		panic(err)
	}

	db, err := services.NewDB("mysql", config.MysqlDatabase, config.MysqlURL)
	if err != nil {
		log.Panic("Connect DB failed", err)
	}

	kafkaWriter := queue.NewWriter(kafkaCfg)
	defer func() { _ = kafkaWriter.Close() }()

	bookingService := services.NewService(db, kafkaWriter, kafkaCfg.KafkaTopic)

	reader := queue.NewReader(kafkaCfg)
	defer func() { _ = reader.Close() }()

	handler := workerSvc.NewHandler(bookingService, reader)

	handler.ConsumeMessages(context.Background())
}
