package queue

import (
	"crypto/tls"
	"hack2hire-2022/worker/config"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewWriter(config config.Conf) *kafka.Writer {
	var tlsConfig *tls.Config
	var err error
	if config.KafkaTLSEnabled {
		tlsConfig, err = newTLSConfig(config.KafkaTLSClientCert, config.KafkaTLSClientKey, config.KafkaTLSCACertFile)
		if err != nil {
			log.Fatal(nil, "setup kafka TLS error", err)
		}
		tlsConfig.InsecureSkipVerify = true
	}

	return &kafka.Writer{
		Addr:                   kafka.TCP(strings.Split(config.KafkaBrokers, ",")...),
		Balancer:               &kafka.Hash{},
		BatchTimeout:           10 * time.Millisecond,
		Topic:                  config.KafkaTopic,
		AllowAutoTopicCreation: true,
		Transport: &kafka.Transport{
			TLS: tlsConfig,
		},
	}
}
