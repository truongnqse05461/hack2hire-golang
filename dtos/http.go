package dtos

import "hack2hire-2022/model"

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Meta struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type User struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type BookingReq struct {
	User      User     `json:"user" binding:"required"`
	SeatCodes []string `json:"seat_codes" binding:"required"`
}

type SaveShowsReq struct {
	Shows []model.Show `json:"shows"`
}

type ShowRes struct {
	Total int          `json:"total"`
	Shows []model.Show `json:"shows"`
}

type SeatRes struct {
	Total int          `json:"total"`
	Seats []model.Seat `json:"seat_list"`
}
