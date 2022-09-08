package services

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DB struct {
	// db *gorm.DB
	db *sqlx.DB
}

func NewDB(dialect, urlConnection string) (*DB, error) {
	// db, err := gorm.Open(dialect, urlConnection)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "open mysql connection error")
	// }

	// if err = db.DB().Ping(); err != nil {
	// 	return nil, errors.Wrap(err, "ping db error")
	// }

	// db.DB().SetConnMaxLifetime(10 * time.Minute)
	// return &DB{
	// 	db: db,
	// }, nil
	db, err := sqlx.Connect(dialect, urlConnection)
	if err != nil {
		return nil, errors.Wrap(err, "open mysql connection error")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping db error")
	}
	db.SetConnMaxLifetime(10 * time.Minute)
	return &DB{
		db: db,
	}, nil
}

func (db *DB) GetMessage(id uint64) (string, error) {
	query := `SELECT message FROM bookings WHERE id = ?;`
	var message string

	err := db.db.QueryRowxContext(context.Background(), query, id).Scan(&message)
	if err != nil {
		return "", err
	}
	return message, nil
}
