package queue

import (
	"hack2hire-2022/worker/config"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewWriter(config config.Conf) *kafka.Writer {
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

	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(config.KafkaBrokers, ","),
		Balancer: &kafka.Hash{},
		Dialer:   dialer,
	})
}
