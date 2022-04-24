CREATE TYPE sex AS enum ('male', 'female');

CREATE TYPE user_type AS enum ('client', 'employee', 'anonymous');

CREATE TYPE emp_position AS enum ('manager', 'employee','owner', 'admin');
 
CREATE TABLE IF NOT EXISTS USERS 
(   user_id              serial PRIMARY key, 
    first_name      CHARACTER VARYING(30) NOT NULL ,
    surname         CHARACTER VARYING(30) NOT NULL ,
    middle_name     CHARACTER VARYING(30) ,    
    email           CHARACTER VARYING(30) ,
    date_of_birth   DATE NOT NULL ,
    address         TEXT NOT NULL ,
    phone           CHARACTER VARYING(30) NOT NULL ,
    password        TEXT ,
    role            user_type NOT NULL ,
    verified        BOOLEAN NOT NULL ,
    sex             sex NOT NULL,          
    photo           TEXT
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

CREATE TABLE IF NOT EXISTS PET 
(   pet_id           serial PRIMARY key ,
    name         CHARACTER VARYING(30) NOT NULL ,
    type         CHARACTER VARYING(30) NOT NULL ,
    weight       SMALLINT NOT NULL ,
    diseases       TEXT,
    user_id      INTEGER REFERENCES USERS(user_id) ON DELETE CASCADE NOT NULL     
);

CREATE TABLE IF NOT EXISTS EMPLOYEE 
(   employee_id          serial PRIMARY key ,
    user_id     INTEGER REFERENCES USERS(user_id) ON DELETE CASCADE NOT NULL ,
    hotel_id    INTEGER REFERENCES HOTEL(hotel_id) ON DELETE CASCADE ,
    position    emp_position NOT NULL 
);

CREATE TABLE IF NOT EXISTS BOOKING 
(   booking_id               serial PRIMARY key,
    seat_id          INTEGER REFERENCES SEAT(seat_id) ON DELETE CASCADE NOT NULL ,
    pet_id           INTEGER REFERENCES PET(pet_id) ON DELETE CASCADE NOT NULL ,
    employee_id      INTEGER REFERENCES EMPLOYEE(employee_id) ON DELETE CASCADE NOT NULL  ,
    status           TEXT ,
    start_date       DATE NOT NULL ,
    end_date         DATE NOT NULL ,
    notes            TEXT
);  

CREATE TABLE IF NOT EXISTS IMAGES 
(   image_id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    ownerId         INTEGER 
);
