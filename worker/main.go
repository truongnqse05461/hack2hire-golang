package main

import (
	"context"
	"hack2hire-2022/queue"
	"hack2hire-2022/worker/config"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	reader := queue.NewReader(config)

	queue.ConsumeMessages(context.Background(), reader)

	//sarama kafka
	// client, err := queue.NewKafkaClient(config)
	// if err != nil {
	// 	log.Fatal(nil, "cannot start kafka client", err)
	// }
	// defer func() { _ = client.Close() }()
	// queue.ConsumeMessage(client, config)
}
