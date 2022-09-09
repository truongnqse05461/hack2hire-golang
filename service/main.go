package main

import (
	"hack2hire-2022/queue"
	"hack2hire-2022/service/config"
	"hack2hire-2022/service/router"
	workerCfg "hack2hire-2022/worker/config"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
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

	kafkaCfg, err := workerCfg.Load()
	if err != nil {
		panic(err)
	}

	//sarama kafka
	// client, err := queue.NewKafkaClient(kafkaCfg)
	// if err != nil {
	// 	log.Fatal(nil, "cannot start kafka client", err)
	// }
	// defer func() { _ = client.Close() }()
	// producer, err := sarama.NewSyncProducerFromClient(client)
	// if err != nil {
	// 	log.Fatal(nil, "cannot create kafka producer", err)
	// }

	kafkaWriter := queue.NewWriter(kafkaCfg)
	defer func() { _ = kafkaWriter.Close() }()

	router := router.NewRouters(config, kafkaWriter, kafkaCfg.KafkaTopic)

	engine, err := router.InitGin()
	if err != nil {
		log.Panic("Failed to init gin", err)
	}

	_ = engine.Run(":" + config.ListenPort)
}
