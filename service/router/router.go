package router

import (
	"context"
	"fmt"
	"hack2hire-2022/model"
	"hack2hire-2022/service/config"
	"hack2hire-2022/service/handler"

	"hack2hire-2022/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Router struct {
	config     *config.Config
	writer     *kafka.Writer
	kafkaTopic string
}

func NewRouters(config *config.Config, writer *kafka.Writer, topic string) *Router {
	return &Router{
		config:     config,
		writer:     writer,
		kafkaTopic: topic,
	}
}

func (r *Router) InitGin() (*gin.Engine, error) {
	db, err := services.NewDB("mysql", r.config.MysqlURL)
	if err != nil {
		log.Panic("Connect DB failed", err)
	}
	bookingService := services.NewService(db, r.writer, r.kafkaTopic)

	// go func() {
	// 	err := generateSeats(bookingService)
	// 	if err != nil {
	// 		fmt.Printf("err generate seats: %s\n", err)
	// 		return
	// 	}
	// 	fmt.Printf("generate seats success\n")
	// }()

	newHandler := handler.NewHandler(bookingService)
	engine := gin.New()
	engine.GET("/health", newHandler.Health)
	bookingGroup := engine.Group("/booking")
	{
		bookingGroup.GET("/hello/:id", newHandler.GetMessage)
		//hack2hire
		bookingGroup.GET("/shows", newHandler.GetShows)
		bookingGroup.GET("/shows/:show_id/seats", newHandler.GetSeats)
		bookingGroup.GET("/shows/:show_id/reservations", newHandler.GetReservations)
		bookingGroup.POST("/shows/:show_id/reservations", newHandler.SaveReservation)

		bookingGroup.POST("/shows", newHandler.SaveShows)
	}
	return engine, nil
}

func generateSeats(svc services.BookingService) error {
	shows, err := svc.GetShows(context.Background())
	if err != nil {
		return err
	}
	id := 1
	var seats []model.Seat
	for _, s := range shows {
		for i := 0; i < 20; i++ {
			seats = append(seats, model.Seat{
				ID:     fmt.Sprintf("%d", id),
				Code:   fmt.Sprintf("S00%d", i+1),
				ShowId: s.ID,
				Status: "AVAILABLE",
			})
			id++
		}
	}
	return svc.SaveSeats(context.Background(), seats...)
}
