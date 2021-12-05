

CREATE TABLE IF NOT EXISTS ACCOUNT 
(   id          serial PRIMARY key, 
    Lgint       TEXT,
    Password    TEXT 
);

CREATE TABLE IF NOT EXISTS USERS 
(   id            serial PRIMARY key,
    FirstName     CHARACTER VARYING(30),
    Surname       CHARACTER VARYING(30),
    MiddleName    CHARACTER VARYING(30),
    Email         CHARACTER VARYING(30),
    DateOfBirth   DATE,
    Adress        TEXT,
    Phone         CHARACTER VARYING(30),
    fk_account_ID INTEGER REFERENCES ACCOUNT(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS EMPLOYEE 
(   id              serial PRIMARY key,
    fk_user_ID      INTEGER REFERENCES USERS(id) ON DELETE CASCADE,
    Position        TEXT,
    Role            TEXT
);

CREATE TABLE IF NOT EXISTS CLIENT 
(   id              serial PRIMARY key, 
    fk_user_ID      INTEGER REFERENCES USERS(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS PET 
(   id           serial PRIMARY key,
    Name         CHARACTER VARYING(30),
    Type         CHARACTER VARYING(30),
    Weignt       SMALLINT,
    Email        CHARACTER VARYING(30),
    Dieses       TEXT,
    fk_client_ID INTEGER REFERENCES CLIENT(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS SEAT 
(   id          serial PRIMARY key, 
    RoomNumber  CHARACTER VARYING(30),
    isFree      BOOLEAN 
);

CREATE TABLE IF NOT EXISTS BOOKING 
(   id               serial PRIMARY key,
    fk_seat_ID       INTEGER REFERENCES SEAT(id) ON DELETE CASCADE,
    fk_pet_ID        INTEGER REFERENCES PET(id) ON DELETE CASCADE ,
    fk_employee_ID   INTEGER REFERENCES EMPLOYEE(id) ON DELETE CASCADE ,
    Status           TEXT,
    StartDate        DATE,
    EndDate          DATE,
    ClientNotes      TEXT
);


CREATE TABLE IF NOT EXISTS ROOM 
(   id           serial PRIMARY key, 
    RoomNumber   CHARACTER VARYING(30),
    petType      TEXT,
    fk_seat_ID   INTEGER REFERENCES SEAT(id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS HOTEL 
(   id              serial PRIMARY key, 
    Name            CHARACTER VARYING(30),
    Address         TEXT,
    fk_room_ID      INTEGER REFERENCES ROOM(id) ON DELETE CASCADE,
    fk_booking_ID   INTEGER REFERENCES BOOKING(id) ON DELETE CASCADE 
);



ALTER TABLE EMPLOYEE ADD COLUMN fk_hotel_ID INTEGER REFERENCES HOTEL(id) ON DELETE CASCADE;
ALTER TABLE CLIENT ADD COLUMN fk_booking_ID INTEGER REFERENCES BOOKING(id) ON DELETE CASCADE;
ALTER TABLE CLIENT ADD COLUMN fk_pet_ID INTEGER REFERENCES PET(id) ON DELETE CASCADE;