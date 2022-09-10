package dtos

import "hack2hire-2022/model"

type PublishReservationMessage struct {
	Reservations []model.Reservation `json:"reservations"`
}
