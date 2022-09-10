package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hack2hire-2022/dtos"
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
		var message dtos.PublishReservationMessage
		if err := json.Unmarshal(m.Value, &message); err != nil {
			_ = fmt.Errorf("error %v", err)
		}
		err = r.bookingSvc.SaveReservation(ctx, message.Reservations...)
		if err != nil {
			_ = fmt.Errorf("error %v", err)
			//notification booking failed
			_ = r.bookingSvc.Notifications(ctx, nil)
		}
		//notification booking sucess
		_ = r.bookingSvc.Notifications(ctx, nil)
	}
}
