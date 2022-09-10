package services

import (
	"context"
	"fmt"
	"hack2hire-2022/model"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

type DB struct {
	db *sqlx.DB
}

func migrateDB(db *sqlx.DB) error {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "bookings_1",
				Up: []string{
					`create table if not exists shows (
						id varchar(255) primary key,
						name varchar(255),
						start_date bigint,
						image_url text
					)`,
					`create table if not exists seats (
						id varchar(255) primary key,
						code varchar(255),
						show_id varchar(255),
						status varchar(255)
					)`,
					`create table if not exists reservations (
						id varchar(255) primary key,
						code varchar(255),
						status varchar(255),
						booked_date bigint,
						canceled_date bigint,
						canceled_reason text,
						user_name varchar(255),
						user_phone varchar(255),
						seat_id varchar(255)
					)`,
				},
				Down: []string{
					"DROP TABLE shows",
					"DROP TABLE seats",
					"DROP TABLE reservations",
				},
			},
			{
				Id: "bookings_2",
				Up: []string{
					`alter table reservations add constraint unq_seat_id_status unique (seat_id, status)`,
				},
			},
			{
				Id: "bookings_3",
				Up: []string{
					``,
				},
			},
			{
				Id: "bookings_4",
				Up: []string{
					``,
				},
			},
		},
	}

	_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	return err
}

func NewDB(dialect, dbName, urlConnection string) (*DB, error) {
	db, err := sqlx.Connect(dialect, urlConnection)
	if err != nil {
		return nil, errors.Wrap(err, "open mysql connection error")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping db error")
	}
	db.SetConnMaxLifetime(10 * time.Minute)
	if err := migrateDB(db); err != nil {
		return nil, err
	}
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

func (db *DB) SaveReservation(ctx context.Context, reservations ...model.Reservation) error {
	tx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	query := `INSERT INTO reservations (id, code, status, booked_date, canceled_date, canceled_reason, user_name, user_phone, seat_id)
			VALUES (:id, :code, :status, :booked_date, :canceled_date, :canceled_reason, :user_name, :user_phone, :seat_id)`
	for _, b := range reservations {
		b.Status = "ACCEPTED"
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

func (db *DB) GetShowByID(ctx context.Context, id string) (show model.Show, err error) {
	query := `SELECT * FROM shows WHERE id = ?`

	if err := db.db.QueryRowxContext(ctx, query, id).StructScan(&show); err != nil {
		return model.Show{}, err
	}
	return show, nil
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

func (db *DB) GetReservations(ctx context.Context, showId, phoneNum string, seatCodes ...string) ([]model.Reservation, error) {
	query := `SELECT r.*, s.code as seat_code, s.show_id as show_id FROM reservations r 
			JOIN seats s 
			ON r.seat_id = s.id 
			WHERE s.show_id = :show_id`

	if phoneNum != "" {
		query += ` AND r.user_phone = :user_phone `
	}
	if len(seatCodes) > 0 {
		query += fmt.Sprintf(` AND s.code in ('%s')`, strings.Join(seatCodes, "','"))
	}

	params := map[string]interface{}{
		"show_id":    showId,
		"user_phone": phoneNum,
	}

	rows, err := db.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var reservations []model.Reservation
	for rows.Next() {
		var reservation model.Reservation
		if err := rows.StructScan(&reservation); err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func (db *DB) GetSeatByCode(ctx context.Context, seatCode, showId string) (model.Seat, error) {
	query := `SELECT * FROM seats WHERE code = ? AND show_id = ?`

	var seat model.Seat
	if err := db.db.QueryRowxContext(ctx, query, seatCode, showId).StructScan(&seat); err != nil {
		return model.Seat{}, err
	}
	return seat, nil
}

func (db *DB) SaveShows(ctx context.Context, shows ...model.Show) error {
	tx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	query := `INSERT INTO shows (id, name, start_date, image_url)
			VALUES (:id, :name, :start_date, :image_url)`
	for _, s := range shows {

		if _, err := tx.NamedExecContext(ctx, query, s); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (db *DB) SaveSeats(ctx context.Context, seats ...model.Seat) error {
	tx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	query := `INSERT INTO seats (id, code, show_id, status)
			VALUES (:id, :code, :show_id, :status)`
	for _, s := range seats {

		if _, err := tx.NamedExecContext(ctx, query, s); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
