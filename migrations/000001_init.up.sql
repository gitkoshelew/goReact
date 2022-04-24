CREATE TABLE IF NOT EXISTS ACCOUNT 
(   id          serial PRIMARY key , 
    Login       TEXT NOT NULL ,
    Password    TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS USERS 
(   user_id              serial PRIMARY key,
    first_name      CHARACTER VARYING(30) NOT NULL ,
    surname         CHARACTER VARYING(30) NOT NULL ,
    middle_name     CHARACTER VARYING(30) ,
    email           CHARACTER VARYING(30) ,
    date_of_birth   DATE NOT NULL ,
    address         TEXT NOT NULL ,
    phone           CHARACTER VARYING(30) NOT NULL ,
    account_id      INTEGER REFERENCES ACCOUNT(id) ON DELETE CASCADE NOT NULL 
);

CREATE TABLE IF NOT EXISTS HOTEL 
(   hotel_id              serial PRIMARY key, 
    name            CHARACTER VARYING(30) NOT NULL ,
    address         TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS ROOM 
(   room_id          serial PRIMARY key, 
    number      CHARACTER VARYING(30) NOT NULL ,
    pet_type    TEXT NOT NULL ,
    hotel_id    INTEGER REFERENCES HOTEL(hotel_id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS SEAT 
(   seat_id           serial PRIMARY key, 
    room_id      INTEGER REFERENCES ROOM(room_id) ON DELETE CASCADE NOT NULL ,
    is_free      BOOLEAN NOT NULL,
    description  TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS CLIENT 
(   client_id          serial PRIMARY key , 
    user_id     INTEGER REFERENCES USERS(user_id) ON DELETE CASCADE NOT NULL 
);

CREATE TABLE IF NOT EXISTS PET 
(   pet_id           serial PRIMARY key ,
    name         CHARACTER VARYING(30) NOT NULL ,
    type         CHARACTER VARYING(30) NOT NULL ,
    weight       NUMERIC NOT NULL ,
    dieses       TEXT,
    client_id    INTEGER REFERENCES CLIENT(client_id) ON DELETE CASCADE NOT NULL 
);

CREATE TABLE IF NOT EXISTS EMPLOYEE 
(   employee_id          serial PRIMARY key ,
    user_id     INTEGER REFERENCES USERS(user_id) ON DELETE CASCADE NOT NULL ,
    hotel_id    INTEGER REFERENCES HOTEL(hotel_id) ON DELETE CASCADE ,
    position    TEXT NOT NULL ,
    role        TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS BOOKING 
(   booking_id               serial PRIMARY key,
    seat_id          INTEGER REFERENCES SEAT(seat_id) ON DELETE CASCADE NOT NULL ,
    pet_id           INTEGER REFERENCES PET(pet_id) ON DELETE CASCADE NOT NULL ,
    employee_id      INTEGER REFERENCES EMPLOYEE(employee_id) ON DELETE CASCADE NOT NULL  ,
    status           TEXT ,
    start_date       DATE NOT NULL ,
    end_date         DATE NOT NULL ,
    client_notes     TEXT
);