package model

type Bookings struct {
	ID      int64  `gorm:"type: bigint; primary_key" json:"id"`
	Message string `gorm:"type: varchar(255)" json:"source"`
}

func (Bookings) TableName() string {
	return "bookings"
}
