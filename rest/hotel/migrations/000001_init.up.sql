CREATE TYPE emp_position AS enum ('manager', 'employee','owner', 'admin');

CREATE TABLE IF NOT EXISTS HOTEL 
(   id              serial PRIMARY key, 
    name            CHARACTER VARYING(30) NOT NULL ,
    address         TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS ROOM 
(   id          serial PRIMARY key, 
    number      CHARACTER VARYING(30) NOT NULL ,
    pet_type    TEXT NOT NULL ,
    hotel_id    INTEGER REFERENCES HOTEL(id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS SEAT 
(   id           serial PRIMARY key, 
    room_id      INTEGER REFERENCES ROOM(id) ON DELETE CASCADE NOT NULL ,
    is_free      BOOLEAN NOT NULL,
    description  TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS EMPLOYEE 
(   id          serial PRIMARY key ,
    user_id     INTEGER NOT NULL UNIQUE,
    hotel_id    INTEGER REFERENCES HOTEL(id) ON DELETE CASCADE ,
    position    emp_position NOT NULL 
);