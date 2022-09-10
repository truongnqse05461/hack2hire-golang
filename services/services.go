package services

import (
	"context"
	"encoding/json"
	"hack2hire-2022/model"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type BookingService interface {
	SayHello(id int64) (string, error)
	Save(bookings model.Bookings) error
}

type bookingService struct {
	db         *DB
	writer     *kafka.Writer
	kafkaTopic string
}

// Save implements BookingService
func (s *bookingService) Save(bookings model.Bookings) error {
	data, err := json.Marshal(bookings)
	if err != nil {
		return err
	}
	err = s.writer.WriteMessages(context.Background(), kafka.Message{
		// Topic: s.kafkaTopic,
		Value: data,
	})
	if err != nil {
		zap.L().Error("send message failed", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (s *bookingService) SayHello(id int64) (string, error) {
	message, err := s.db.GetMessage(uint64(id))
	if err != nil {
		return "", err
	}
	return message, nil
}

var _ BookingService = (*bookingService)(nil)

func NewService(db *DB, writer *kafka.Writer, topic string) BookingService {
	return &bookingService{db: db, writer: writer, kafkaTopic: topic}
}
