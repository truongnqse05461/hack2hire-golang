package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hack2hire-2022/model"
	"hack2hire-2022/services"
	"log"

	"github.com/segmentio/kafka-go"
)

type ReservationHandler struct {
	bookingSvc services.BookingService
	reader     *kafka.Reader
}

func NewHandler(svc services.BookingService, reader *kafka.Reader) ReservationHandler {
	return ReservationHandler{
		bookingSvc: svc,
		reader:     reader,
	}
}

func (r *ReservationHandler) ConsumeMessages(ctx context.Context) {
	log.Println("start listening to events ...")
	for {
		m, err := r.reader.ReadMessage(ctx)
		if err != nil {
			_ = fmt.Errorf("error %v", err)
		}
		fmt.Printf("receive event on broker %s, topic %s, partition %d, offset %d, data %s\n", r.reader.Config().Brokers, m.Topic, m.Partition, m.Offset, string(m.Value))
		var reservation model.Reservation
		if err := json.Unmarshal(m.Value, &reservation); err != nil {
			_ = fmt.Errorf("error %v", err)
		}
		r.bookingSvc.SaveReservation(ctx, reservation)
	}
}
