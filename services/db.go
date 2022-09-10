package services

import (
	"context"
	"hack2hire-2022/model"
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

func (db *DB) SaveReservation(ctx context.Context, bookings ...model.Bookings) error {
	tx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	query := `INSERT INTO reservations (id, code, status, booked_date, user_name, user_phone, seat_id)
			VALUES (:id, :code, :status, :booked_date, :user_name, :user_phone, :seat_id)`
	for _, b := range bookings {

		if _, err := tx.NamedExecContext(ctx, query, b); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (db *DB) GetShows(ctx context.Context) ([]model.Show, error) {
	query := `SELECT * FROM shows`

	rows, err := db.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var shows []model.Show
	for rows.Next() {
		var show model.Show
		if err := rows.StructScan(&show); err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}
	return shows, nil
}

func (db *DB) GetSeats(ctx context.Context, showId string) ([]model.Seat, error) {
	query := `SELECT * FROM seats WHERE show_id = ?`

	rows, err := db.db.QueryxContext(ctx, query, showId)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var seats []model.Seat
	for rows.Next() {
		var seat model.Seat
		if err := rows.StructScan(&seat); err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}

func (db *DB) GetSeatByCode(ctx context.Context, seatCode string) (model.Seat, error) {
	query := `SELECT * FROM seats WHERE code = ?`

	var seat model.Seat
	if err := db.db.QueryRowxContext(ctx, query, seatCode).StructScan(&seat); err != nil {
		return model.Seat{}, err
	}
	return seat, nil
}
