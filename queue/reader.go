package queue

import (
	"context"
	"fmt"
	"hack2hire-2022/worker/config"
	"log"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func NewReader(config config.Conf) *kafka.Reader {
	var dialer *kafka.Dialer
	if config.KafkaTLSEnabled {
		tlsConfig, err := newTLSConfig(config.KafkaTLSClientCert, config.KafkaTLSClientKey, config.KafkaTLSCACertFile)
		if err != nil {
			log.Fatal(nil, "setup kafka TLS error", err)
		}
		tlsConfig.InsecureSkipVerify = true
		dialer = &kafka.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
			TLS:       tlsConfig,
		}
	}
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     strings.Split(config.KafkaBrokers, ","),
		GroupID:     config.KafkaConsumerGroup,
		Topic:       config.KafkaTopic,
		StartOffset: kafka.LastOffset,
		Dialer:      dialer,
	})
}

func ConsumeMessages(ctx context.Context, reader *kafka.Reader) {
	log.Println("start listening to events ...")
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			_ = fmt.Errorf("error %v", err)
		}
		fmt.Printf("receive event on broker %s, topic %s, partition %d, offset %d, data %s\n", reader.Config().Brokers, m.Topic, m.Partition, m.Offset, string(m.Value))
	}
}
