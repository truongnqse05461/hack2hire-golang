package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"log"
	"strings"
)

type Event struct {
	Event string `json:"event"`
}

type consumerGroupHandler struct{ c chan []byte }

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.c <- msg.Value
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	client, err := newKafkaClient(config)
	if err != nil {
		log.Fatal(nil, "cannot start kafka client", err)
	}
	defer func() { _ = client.Close() }()

	consumeMessage(client, config)
}

func consumeMessage(client sarama.Client, config Conf) {
	group, err := sarama.NewConsumerGroupFromClient(config.KafkaConsumerGroup, client)
	if err != nil {
		log.Fatal(nil, "error on initialing kafka connection", err)
	}

	defer func() { _ = group.Close() }()

	go func() {
		for err := range group.Errors() {
			_ = fmt.Errorf("%v", err)
		}
	}()

	kafkaChan := make(chan []byte)
	eventChan := make(chan Event)

	go func() {
		for d := range kafkaChan {
			var eventData Event
			err := json.Unmarshal(d, &eventData)
			if err != nil {
				_ = fmt.Errorf("error %v", err)
			} else {
				eventChan <- eventData
			}
		}
	}()

	go func() {
		for e := range eventChan {
			log.Printf("Sample Event %v", e)
		}
	}()
	log.Println("start listening to events ...")
	ctx := context.Background()
	for {
		topics := strings.Split(config.KafkaTopic, ",")
		handler := consumerGroupHandler{kafkaChan}
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			_ = fmt.Errorf("error %v", err)
		}
	}
}

func newKafkaClient(config Conf) (sarama.Client, error) {
	configKafka := sarama.NewConfig()
	configKafka.Version = sarama.V1_0_0_0
	configKafka.Consumer.Return.Errors = true

	if config.KafkaTLSEnabled {
		tlsConfig, err := newTLSConfig(config.KafkaTLSClientCert, config.KafkaTLSClientKey, config.KafkaTLSCACertFile)
		if err != nil {
			log.Fatal(nil, "setup kafka TLS error", err)
		}
		tlsConfig.InsecureSkipVerify = true

		configKafka.Net.TLS.Enable = true
		configKafka.Net.TLS.Config = tlsConfig
	}

	client, err := sarama.NewClient(strings.Split(config.KafkaBrokers, ","), configKafka)
	if err != nil {
		log.Fatal(nil, "error on initialing kafka connection", err)
	}

	return client, err
}

func newTLSConfig(clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = caCertPool
	return &tlsConfig, err
}
