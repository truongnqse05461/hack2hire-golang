package services

import (
	"context"
	"encoding/json"
	"errors"
	"hack2hire-2022/dtos"
	"hack2hire-2022/model"
	"hack2hire-2022/utils"

	"github.com/segmentio/kafka-go"
)

type BookingService interface {
	SayHello(id int64) (string, error)
	PublishReservation(ctx context.Context, req dtos.BookingReq) error
	SaveReservation(ctx context.Context, reservation model.Reservation) error
	GetShows(ctx context.Context) ([]model.Show, error)
	GetSeats(ctx context.Context, showId string) ([]model.Seat, error)
	GetReservations(ctx context.Context, showId, phoneNum string, seatCodes ...string) ([]model.Reservation, error)
	Notifications(ctx context.Context, message interface{}) error

	SaveShows(ctx context.Context, shows ...model.Show) error
	SaveSeats(ctx context.Context, seats ...model.Seat) error
}

type bookingService struct {
	db         *DB
	writer     *kafka.Writer
	kafkaTopic string
}

// PublishReservation implements BookingService
func (*bookingService) PublishReservation(ctx context.Context, req dtos.BookingReq) error {
	panic("unimplemented")
}

// Notifications implements BookingService
func (b *bookingService) Notifications(ctx context.Context, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return b.writer.WriteMessages(ctx, kafka.Message{
		Value: data,
	})
}

// SaveReservation implements BookingService
func (b *bookingService) SaveReservation(ctx context.Context, reservation model.Reservation) error {
	return b.db.SaveReservation(ctx, reservation)
}

// GetReservations implements BookingService
func (b *bookingService) GetReservations(ctx context.Context, showId string, phoneNum string, seatCodes ...string) ([]model.Reservation, error) {
	// check show exist
	if _, err := b.db.GetShowByID(ctx, showId); err != nil {
		return nil, errors.New(utils.ErrShowNotFound)
	}
	return b.db.GetReservations(ctx, showId, phoneNum, seatCodes...)
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
	// check show exist
	if _, err := b.db.GetShowByID(ctx, showId); err != nil {
		return nil, errors.New(utils.ErrShowNotFound)
	}

	return b.db.GetSeats(ctx, showId)
}

// GetShows implements BookingService
func (b *bookingService) GetShows(ctx context.Context) ([]model.Show, error) {
	return b.db.GetShows(ctx)
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
