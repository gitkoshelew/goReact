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

INSERT INTO ACCOUNT ( Login, Password) VALUES 
('login1', 'password1'),
('login2', 'password2'),
('login3', 'password3'),
('login4', 'password4'),
('login5', 'password5'),
('login6', 'password6');

INSERT INTO USERS (first_name , surname, middle_name, email, date_of_birth, address, phone, account_id) VALUES 
('Ivan','Ivanov','Ivanovich','ivan@mail.ru','2000-01-01' ,'Minsk Pr. Nezavisimosti 22-222' ,'+375-29-154-89-33', 1),
('Petr','Petrov','Petrovich','petr@mail.ru','1999-03-13 ','Minsk Pr. Pobediteley 11-111' ,'+375-29-159-11-78', 2),
('Maria','Petrova','Evgenievna','liza@mail.ru','2001-11-11' ,'Minsk Pr. Pobediteley 111-111' ,'+375-29-655-99-14', 3),
('Olga','Oleeva','Vladimirovna','olga@mail.ru','2001-01-01' ,'Minsk ul. Nesterova 1-32' ,'+375-29-675-00-00', 4),
('Vladzimir','Sakhonchik','Alekseevish','leha@mail.ru','1998-02-11' ,'Minsk Partizanskiy 124-1' ,'+375-29-111-22-33', 5),
('Mikhail','Valevach','Dmitrievich','miha@mail.ru','1997-01-12' ,'Minsk Pr. Dzerjinskogo 01-01' ,'+375-29-777-55-44', 6);

INSERT INTO HOTEL (name, address) VALUES 
('PetsHotel1','ul. Pushkina 12'),
('PetsHotel2','ul. Sovetskaya 16'),
('PetsHotel3','ul. Ilimskaya 33');

INSERT INTO ROOM (pet_type, number, hotel_id) VALUES 
('Cat', 101, 1),
('Dog', 202, 2),
('Cat', 303, 3);

INSERT INTO SEAT (room_id, is_free, description) VALUES 
(1, false, 'VIP seat'),
(2, true, 'Seat for sick pets'),
(3, false, 'Regular seat');

INSERT INTO CLIENT (user_id) VALUES 
(1),
(2),
(3);

INSERT INTO PET (name , type, weignt, dieses, client_id) VALUES 
('Murzik','cat',5 ,'no dieses' ,1),
('Barbos','dog',5 ,'1 dieses' ,2),
('Aliy','dog',5 ,'2 dieses' ,3);

INSERT INTO EMPLOYEE (user_id, hotel_id, position, role) VALUES 
(4 , 1, 'Position 1', 'role 1' ),
(5 , 2, 'Position 3', 'role 2' ),
(6 , 3, 'Position 3', 'role 3' );

INSERT INTO BOOKING (seat_id , pet_id, employee_id, status, start_date, end_date, client_notes) VALUES 
(1 ,1 ,1,'In processing', '2021-12-07', '2021-12-27','wash my pet pls twice a day'),
(2 ,2 ,2,'In work', '2021-12-01', '2021-12-15','feed my pet pls once a week'),
(3, 3 ,3,'Ended up', '2021-11-26', '2021-12-06','no comm' );

INSERT INTO IMAGES (type, URL, ownerId) VALUES 
('pet', 'volume\images\pets\cat1.jpg', 1),
('pet', 'volume\images\pets\cat2.jpg', 2),
('pet', 'volume\images\pets\cat3.jpg', 3),
('pet', 'volume\images\pets\dog1.jpg', 3),
('pet', 'volume\images\pets\dog2.jpg', 3),
('pet', 'volume\images\pets\dog3.jpg', 3),
('room', 'volume\images\rooms\room1.jpg', 1),
('room', 'volume\images\rooms\room2.jpg', 2),
('room', 'volume\images\rooms\room3.jpg', 3);
