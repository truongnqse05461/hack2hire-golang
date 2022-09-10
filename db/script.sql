create table shows (
	id varchar(255) primary key,
    name varchar(255),
    start_date bigint,
    image_url text
);

create table reservations (
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

create table seats (
	id varchar(255) primary key,
    code varchar(255),
    show_id varchar(255),
    status varchar(255)
);
insert into seats values ('1', 'S001', '2', 'AVAILABLE');
insert into seats values ('2', 'S002', '2', 'AVAILABLE');
insert into seats values ('3', 'S003', '2', 'AVAILABLE');
insert into seats values ('4', 'S004', '2', 'AVAILABLE');
insert into seats values ('5', 'S005', '2', 'AVAILABLE');
insert into seats values ('6', 'S006', '2', 'AVAILABLE');
insert into seats values ('7', 'S007', '2', 'AVAILABLE');
insert into seats values ('8', 'S008', '2', 'AVAILABLE');
insert into seats values ('9', 'S009', '2', 'AVAILABLE');
insert into seats values ('10', 'S0010', '2', 'AVAILABLE');
insert into seats values ('11', 'S0011', '2', 'AVAILABLE');
insert into seats values ('12', 'S0012', '2', 'AVAILABLE');
insert into seats values ('13', 'S0013', '2', 'AVAILABLE');
insert into seats values ('14', 'S0014', '2', 'AVAILABLE');
insert into seats values ('15', 'S0015', '2', 'AVAILABLE');
insert into seats values ('16', 'S0016', '2', 'AVAILABLE');
insert into seats values ('17', 'S0017', '2', 'AVAILABLE');
insert into seats values ('18', 'S0018', '2', 'AVAILABLE');
insert into seats values ('19', 'S0019', '2', 'AVAILABLE');
insert into seats values ('20', 'S0020', '2', 'AVAILABLE');


select * from shows;

drop table shows;
drop table reservations;
drop table seats;