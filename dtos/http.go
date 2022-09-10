package dtos

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
