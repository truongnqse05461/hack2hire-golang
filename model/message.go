package model

type Message struct {
	ID      int64  `gorm:"type: bigint; primary_key" json:"id"`
	Message string `gorm:"type: varchar(255)" json:"source"`
}

func (Message) TableName() string {
	return "message"
}
