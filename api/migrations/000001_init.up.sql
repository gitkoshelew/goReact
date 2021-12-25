CREATE TABLE IF NOT EXISTS ACCOUNT 
(   id          serial PRIMARY key , 
    Login       TEXT NOT NULL ,
    Password    TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS USERS 
(   id              serial PRIMARY key,
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

CREATE TABLE IF NOT EXISTS CLIENT 
(   id          serial PRIMARY key , 
    user_id     INTEGER REFERENCES USERS(id) ON DELETE CASCADE NOT NULL 
);

CREATE TABLE IF NOT EXISTS PET 
(   id           serial PRIMARY key ,
    name         CHARACTER VARYING(30) NOT NULL ,
    type         CHARACTER VARYING(30) NOT NULL ,
    weignt       SMALLINT NOT NULL ,
    dieses       TEXT,
    client_id    INTEGER REFERENCES CLIENT(id) ON DELETE CASCADE NOT NULL 
);

CREATE TABLE IF NOT EXISTS EMPLOYEE 
(   id          serial PRIMARY key ,
    user_id     INTEGER REFERENCES USERS(id) ON DELETE CASCADE NOT NULL ,
    hotel_id    INTEGER REFERENCES HOTEL(id) ON DELETE CASCADE ,
    position    TEXT NOT NULL ,
    role        TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS BOOKING 
(   id               serial PRIMARY key,
    seat_id          INTEGER REFERENCES SEAT(id) ON DELETE CASCADE NOT NULL ,
    pet_id           INTEGER REFERENCES PET(id) ON DELETE CASCADE NOT NULL ,
    employee_id      INTEGER REFERENCES EMPLOYEE(id) ON DELETE CASCADE NOT NULL  ,
    status           TEXT ,
    start_date       DATE NOT NULL ,
    end_date         DATE NOT NULL ,
    client_notes     TEXT
);

CREATE TABLE IF NOT EXISTS IMAGES 
(   id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    URL             TEXT NOT NULL ,
    ownerId         INTEGER 
);

CREATE TABLE IF NOT EXISTS TOKENS 
(   uuid TEXT,
    userid INTEGER,
    expire CHARACTER VARYING(30)
);