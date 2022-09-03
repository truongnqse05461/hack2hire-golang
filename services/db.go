package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

type DB struct {
	db *gorm.DB
}

func NewDB(dialect, urlConnection string) (*DB, error) {
	db, err := gorm.Open(dialect, urlConnection)
	if err != nil {
		return nil, errors.Wrap(err, "open mysql connection error")
	}

	if err = db.DB().Ping(); err != nil {
		return nil, errors.Wrap(err, "ping db error")
	}

	db.DB().SetConnMaxLifetime(10 * time.Minute)
	return &DB{
		db: db,
	}, nil
}

func (db *DB) GetMessage(id uint64) (string, error) {
	var message string
	err := db.db.Table("bookings").Select("message").Where("id = ?", id).Find(&message).Error
	if err != nil {
		return "", err
	}
	return message, nil
}
