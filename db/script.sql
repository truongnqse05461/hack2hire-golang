create table if not exists shows (
	id varchar(255) primary key,
    name varchar(255),
    start_date bigint,
    image_url text
);

create table if not exists reservations (
	id varchar(255) primary key,
    code varchar(255),
    status varchar(255),
    booked_date bigint,
    canceled_date bigint,
    canceled_reason text,
    user_name varchar(255),
    user_phone varchar(255),
    seat_id varchar(255)
);

create table if not exists seats (
	id varchar(255) primary key,
    code varchar(255),
    show_id varchar(255),
    status varchar(255)
);

alter table reservations add constraint unq_seat_id_status unique (seat_id, status); 