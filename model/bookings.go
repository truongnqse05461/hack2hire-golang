package model

type User struct {
	Name        string `json:"name" db:"user_name"`
	PhoneNumber string `json:"phone_number" db:"user_phone"`
}

type Bookings struct {
	ID             string `json:"id" db:"id"`
	Code           string `json:"code" db:"code"`
	Status         string `json:"status" db:"status"`
	BookedDate     int64  `json:"booked_date" db:"booked_date"`
	CanceledDate   int64  `json:"canceled_date" db:"canceled_date"`
	CanceledReason string `json:"canceled_reason" db:"canceled_reason"`
	SeatId         string `json:"seat_id" db:"seat_id"`
	User
}

type Show struct {
	ID        string `json:"show_id" db:"id"`
	Name      string `json:"name" db:"name"`
	StartDate int64  `json:"start_date" db:"start_date"`
	ImageUrl  string `json:"image_url" db:"image_url"`
}

type Seat struct {
	ID         string `json:"seat_id" db:"id"`
	Code       string `json:"seat_code" db:"code"`
	ShowId     string `json:"show_id" db:"show_id"`
	Status     string `json:"status" db:"status"`
	BookedDate int64  `json:"booked_date,omitempty" db:"booked_date"`
}
