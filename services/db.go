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
					`INSERT INTO seats VALUES ('1','S001','1','AVAILABLE'),('10','S0010','1','AVAILABLE'),('100','S0020','2','AVAILABLE'),('101','S001','3','AVAILABLE'),('102','S002','3','AVAILABLE'),('103','S003','3','AVAILABLE'),('104','S004','3','AVAILABLE'),('105','S005','3','AVAILABLE'),('106','S006','3','AVAILABLE'),('107','S007','3','AVAILABLE'),('108','S008','3','AVAILABLE'),('109','S009','3','AVAILABLE'),('11','S0011','1','AVAILABLE'),('110','S0010','3','AVAILABLE'),('111','S0011','3','AVAILABLE'),('112','S0012','3','AVAILABLE'),('113','S0013','3','AVAILABLE'),('114','S0014','3','AVAILABLE'),('115','S0015','3','AVAILABLE'),('116','S0016','3','AVAILABLE'),('117','S0017','3','AVAILABLE'),('118','S0018','3','AVAILABLE'),('119','S0019','3','AVAILABLE'),('12','S0012','1','AVAILABLE'),('120','S0020','3','AVAILABLE'),('121','S001','4','AVAILABLE'),('122','S002','4','AVAILABLE'),('123','S003','4','AVAILABLE'),('124','S004','4','AVAILABLE'),('125','S005','4','AVAILABLE'),('126','S006','4','AVAILABLE'),('127','S007','4','AVAILABLE'),('128','S008','4','AVAILABLE'),('129','S009','4','AVAILABLE'),('13','S0013','1','AVAILABLE'),('130','S0010','4','AVAILABLE'),('131','S0011','4','AVAILABLE'),('132','S0012','4','AVAILABLE'),('133','S0013','4','AVAILABLE'),('134','S0014','4','AVAILABLE'),('135','S0015','4','AVAILABLE'),('136','S0016','4','AVAILABLE'),('137','S0017','4','AVAILABLE'),('138','S0018','4','AVAILABLE'),('139','S0019','4','AVAILABLE'),('14','S0014','1','AVAILABLE'),('140','S0020','4','AVAILABLE'),('141','S001','5','AVAILABLE'),('142','S002','5','AVAILABLE'),('143','S003','5','AVAILABLE'),('144','S004','5','AVAILABLE'),('145','S005','5','AVAILABLE'),('146','S006','5','AVAILABLE'),('147','S007','5','AVAILABLE'),('148','S008','5','AVAILABLE'),('149','S009','5','AVAILABLE'),('15','S0015','1','AVAILABLE'),('150','S0010','5','AVAILABLE'),('151','S0011','5','AVAILABLE'),('152','S0012','5','AVAILABLE'),('153','S0013','5','AVAILABLE'),('154','S0014','5','AVAILABLE'),('155','S0015','5','AVAILABLE'),('156','S0016','5','AVAILABLE'),('157','S0017','5','AVAILABLE'),('158','S0018','5','AVAILABLE'),('159','S0019','5','AVAILABLE'),('16','S0016','1','AVAILABLE'),('160','S0020','5','AVAILABLE'),('161','S001','6','AVAILABLE'),('162','S002','6','AVAILABLE'),('163','S003','6','AVAILABLE'),('164','S004','6','AVAILABLE'),('165','S005','6','AVAILABLE'),('166','S006','6','AVAILABLE'),('167','S007','6','AVAILABLE'),('168','S008','6','AVAILABLE'),('169','S009','6','AVAILABLE'),('17','S0017','1','AVAILABLE'),('170','S0010','6','AVAILABLE'),('171','S0011','6','AVAILABLE'),('172','S0012','6','AVAILABLE'),('173','S0013','6','AVAILABLE'),('174','S0014','6','AVAILABLE'),('175','S0015','6','AVAILABLE'),('176','S0016','6','AVAILABLE'),('177','S0017','6','AVAILABLE'),('178','S0018','6','AVAILABLE'),('179','S0019','6','AVAILABLE'),('18','S0018','1','AVAILABLE'),('180','S0020','6','AVAILABLE'),('181','S001','7','AVAILABLE'),('182','S002','7','AVAILABLE'),('183','S003','7','AVAILABLE'),('184','S004','7','AVAILABLE'),('185','S005','7','AVAILABLE'),('186','S006','7','AVAILABLE'),('187','S007','7','AVAILABLE'),('188','S008','7','AVAILABLE'),('189','S009','7','AVAILABLE'),('19','S0019','1','AVAILABLE'),('190','S0010','7','AVAILABLE'),('191','S0011','7','AVAILABLE'),('192','S0012','7','AVAILABLE'),('193','S0013','7','AVAILABLE'),('194','S0014','7','AVAILABLE'),('195','S0015','7','AVAILABLE'),('196','S0016','7','AVAILABLE'),('197','S0017','7','AVAILABLE'),('198','S0018','7','AVAILABLE'),('199','S0019','7','AVAILABLE'),('2','S002','1','AVAILABLE'),('20','S0020','1','AVAILABLE'),('200','S0020','7','AVAILABLE'),('201','S001','8','AVAILABLE'),('202','S002','8','AVAILABLE'),('203','S003','8','AVAILABLE'),('204','S004','8','AVAILABLE'),('205','S005','8','AVAILABLE'),('206','S006','8','AVAILABLE'),('207','S007','8','AVAILABLE'),('208','S008','8','AVAILABLE'),('209','S009','8','AVAILABLE'),('21','S001','10','AVAILABLE'),('210','S0010','8','AVAILABLE'),('211','S0011','8','AVAILABLE'),('212','S0012','8','AVAILABLE'),('213','S0013','8','AVAILABLE'),('214','S0014','8','AVAILABLE'),('215','S0015','8','AVAILABLE'),('216','S0016','8','AVAILABLE'),('217','S0017','8','AVAILABLE'),('218','S0018','8','AVAILABLE'),('219','S0019','8','AVAILABLE'),('22','S002','10','AVAILABLE'),('220','S0020','8','AVAILABLE'),('221','S001','9','AVAILABLE'),('222','S002','9','AVAILABLE'),('223','S003','9','AVAILABLE'),('224','S004','9','AVAILABLE'),('225','S005','9','AVAILABLE'),('226','S006','9','AVAILABLE'),('227','S007','9','AVAILABLE'),('228','S008','9','AVAILABLE'),('229','S009','9','AVAILABLE'),('23','S003','10','AVAILABLE'),('230','S0010','9','AVAILABLE'),('231','S0011','9','AVAILABLE'),('232','S0012','9','AVAILABLE'),('233','S0013','9','AVAILABLE'),('234','S0014','9','AVAILABLE'),('235','S0015','9','AVAILABLE'),('236','S0016','9','AVAILABLE'),('237','S0017','9','AVAILABLE'),('238','S0018','9','AVAILABLE'),('239','S0019','9','AVAILABLE'),('24','S004','10','AVAILABLE'),('240','S0020','9','AVAILABLE'),('25','S005','10','AVAILABLE'),('26','S006','10','AVAILABLE'),('27','S007','10','AVAILABLE'),('28','S008','10','AVAILABLE'),('29','S009','10','AVAILABLE'),('3','S003','1','AVAILABLE'),('30','S0010','10','AVAILABLE'),('31','S0011','10','AVAILABLE'),('32','S0012','10','AVAILABLE'),('33','S0013','10','AVAILABLE'),('34','S0014','10','AVAILABLE'),('35','S0015','10','AVAILABLE'),('36','S0016','10','AVAILABLE'),('37','S0017','10','AVAILABLE'),('38','S0018','10','AVAILABLE'),('39','S0019','10','AVAILABLE'),('4','S004','1','AVAILABLE'),('40','S0020','10','AVAILABLE'),('41','S001','11','AVAILABLE'),('42','S002','11','AVAILABLE'),('43','S003','11','AVAILABLE'),('44','S004','11','AVAILABLE'),('45','S005','11','AVAILABLE'),('46','S006','11','AVAILABLE'),('47','S007','11','AVAILABLE'),('48','S008','11','AVAILABLE'),('49','S009','11','AVAILABLE'),('5','S005','1','AVAILABLE'),('50','S0010','11','AVAILABLE'),('51','S0011','11','AVAILABLE'),('52','S0012','11','AVAILABLE'),('53','S0013','11','AVAILABLE'),('54','S0014','11','AVAILABLE'),('55','S0015','11','AVAILABLE'),('56','S0016','11','AVAILABLE'),('57','S0017','11','AVAILABLE'),('58','S0018','11','AVAILABLE'),('59','S0019','11','AVAILABLE'),('6','S006','1','AVAILABLE'),('60','S0020','11','AVAILABLE'),('61','S001','12','AVAILABLE'),('62','S002','12','AVAILABLE'),('63','S003','12','AVAILABLE'),('64','S004','12','AVAILABLE'),('65','S005','12','AVAILABLE'),('66','S006','12','AVAILABLE'),('67','S007','12','AVAILABLE'),('68','S008','12','AVAILABLE'),('69','S009','12','AVAILABLE'),('7','S007','1','AVAILABLE'),('70','S0010','12','AVAILABLE'),('71','S0011','12','AVAILABLE'),('72','S0012','12','AVAILABLE'),('73','S0013','12','AVAILABLE'),('74','S0014','12','AVAILABLE'),('75','S0015','12','AVAILABLE'),('76','S0016','12','AVAILABLE'),('77','S0017','12','AVAILABLE'),('78','S0018','12','AVAILABLE'),('79','S0019','12','AVAILABLE'),('8','S008','1','AVAILABLE'),('80','S0020','12','AVAILABLE'),('81','S001','2','AVAILABLE'),('82','S002','2','AVAILABLE'),('83','S003','2','AVAILABLE'),('84','S004','2','AVAILABLE'),('85','S005','2','AVAILABLE'),('86','S006','2','AVAILABLE'),('87','S007','2','AVAILABLE'),('88','S008','2','AVAILABLE'),('89','S009','2','AVAILABLE'),('9','S009','1','AVAILABLE'),('90','S0010','2','AVAILABLE'),('91','S0011','2','AVAILABLE'),('92','S0012','2','AVAILABLE'),('93','S0013','2','AVAILABLE'),('94','S0014','2','AVAILABLE'),('95','S0015','2','AVAILABLE'),('96','S0016','2','AVAILABLE'),('97','S0017','2','AVAILABLE'),('98','S0018','2','AVAILABLE'),('99','S0019','2','AVAILABLE');`,
				},
			},
			{
				Id: "bookings_4",
				Up: []string{
					`INSERT INTO shows VALUES ('1','Mặt trời trong suối lạnh',1662720744000,'https://cdn-www.vinid.net/3eb1bf15-banner-tin-tuc-1920x1080-1-1.jpg'),('10','Mây Lang Thang',1662720744000,'https://cdn-www.vinid.net/0368f61d-web-new-1920x1080-1.jpg'),('11','string',1662775772000,'https://cdn-www.vinid.net/0368f61d-web-new-1920x1080-1.jpg'),('12','test',1662776720000,'https://cdn-www.vinid.net/0368f61d-web-new-1920x1080-1.jpg'),('2','Muôn dặm đường xa',1662720744000,'https://cdn-www.vinid.net/3eb1bf15-banner-tin-tuc-1920x1080-1-1.jpg'),('3','VinID – 5 năm bên nhau',1662720744000,'https://cdn-www.vinid.net/2020/07/a29355db-1920x1080-top-banner-web.jpg'),('4','Uyên Linh & con đường âm nhạc',1662720744000,'https://cdn-www.vinid.net/b2877068-uyen-linh.jpg'),('5','Tạ Tình',1662720744000,'https://cdn-www.vinid.net/fa8741f7-web-news.png'),('6','Màu Nước Band',1662720744000,'https://cdn-www.vinid.net/4128e455-mau-nuoc-band.jpg'),('7','The Rap Show',1662720744000,'https://cdn-www.vinid.net/c3fd2f56-tin-tuc-app-1920x1080-1.jpg'),('8','Above World',1662720744000,'https://cdn-www.vinid.net/7609cf94-1920x1080-1.jpg'),('9','Đen Vâu',1662720744000,'https://cdn-www.vinid.net/2020/09/4760435c-rapper-vi%E1%BB%87t-nam.png');`,
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
