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
	GetShows(ctx context.Context) ([]model.Show, error)
	GetSeats(ctx context.Context, showId string) ([]model.Seat, error)
	SaveShows(ctx context.Context, shows ...model.Show) error
	SaveSeats(ctx context.Context, seats ...model.Seat) error
}

type bookingService struct {
	db         *DB
	writer     *kafka.Writer
	kafkaTopic string
}

// SaveSeats implements BookingService
func (b *bookingService) SaveSeats(ctx context.Context, seats ...model.Seat) error {
	return b.db.SaveSeats(ctx, seats...)
}

// SaveShows implements BookingService
func (b *bookingService) SaveShows(ctx context.Context, shows ...model.Show) error {
	return b.db.SaveShows(ctx, shows...)
}

// GetSeats implements BookingService
func (b *bookingService) GetSeats(ctx context.Context, showId string) ([]model.Seat, error) {
	return b.db.GetSeats(ctx, showId)
}

// GetShows implements BookingService
func (b *bookingService) GetShows(ctx context.Context) ([]model.Show, error) {
	return b.db.GetShows(ctx)
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
