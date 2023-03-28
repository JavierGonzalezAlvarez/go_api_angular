create table header (
    id_header serial,
    companyname varchar(50),
    address varchar(50),
    numberinvoice integer,
    date_time timestamp,
    created_at timestamp,
    PRIMARY KEY(id_header)
);

create table detail (
    id_detail serial primary key,
    id_header integer,
    description varchar(50),
    units  integer NOT NULL,
    price float,
    created_at timestamp,
    CONSTRAINT fk_header
        FOREIGN KEY(id_header)
            REFERENCES header(id_header)
);

INSERT INTO header(companyname, address, numberinvoice, date_time, created_at)
VALUES('ABC', 'uria', 12, '2023-02-12 15:00:00', now()),
      ('Dope', 'maissonave', 13, '2023-02-12 15:00:00', now());

INSERT INTO detail(id_header, description, units, price, created_at)
VALUES(1, 'product A', 2, 23.2, now()),
      (1, 'product B', 5, 10.0, now()),
      (2, 'product B', 19, 45.31, now());

create table usuario (
    id serial,
    username varchar(25),
    email varchar(50),
    created_at timestamp,
    PRIMARY KEY(id)
);

ALTER TABLE usuario ADD COLUMN token VARCHAR(255);