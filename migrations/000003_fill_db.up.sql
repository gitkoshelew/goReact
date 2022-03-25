INSERT INTO ACCOUNT ( Login, Password) VALUES 
('login1', 'password1'),
('login2', 'password2'),
('login3', 'password3'),
('login4', 'password4'),
('login5', 'password5'),
('login6', 'password6');

INSERT INTO USERS (email, password, role, verified, first_name , surname, middle_name, sex, date_of_birth, address, phone, photo) VALUES 
('ivan@mail.ru', 'password1', 'client', true, 'Ivan','Ivanov','Ivanovich', 'male', '2000-01-01', 'Minsk Pr. Nezavisimosti 22-222' ,'+375-29-154-89-33', 'PhotoURL...'),
('petr@mail.ru', 'password2', 'client', true, 'Petr','Petrov','Petrovich', 'male', '1999-03-13 ', 'Minsk Pr. Pobediteley 11-111' ,'+375-29-159-11-78', 'PhotoURL...'),
('liza@mail.ru', 'password3', 'client', true, 'Maria','Petrova','Evgenievna', 'female', '2001-11-11', 'Minsk Pr. Pobediteley 111-111' ,'+375-29-655-99-14', 'PhotoURL...'),
('olga@mail.ru', 'password4', 'employee', true, 'Olga','Oleeva','Vladimirovna', 'female', '2001-01-01', 'Minsk ul. Nesterova 1-32' ,'+375-29-675-00-00', 'PhotoURL...'),
('leha@mail.ru', 'password5', 'employee', true, 'Vladzimir','Sakhonchik','Alekseevish', 'male', '1998-02-11', 'Minsk Partizanskiy 124-1' ,'+375-29-111-22-33', 'PhotoURL...'),
('miha@mail.ru', 'password6', 'employee', true, 'Mikhail','Valevach','Dmitrievich', 'male', '1997-01-12', 'Minsk Pr. Dzerjinskogo 01-01' ,'+375-29-777-55-44', 'PhotoURL...');

INSERT INTO HOTEL (name, address) VALUES 
('PetsHotel1','ul. Pushkina 12'),
('PetsHotel2','ul. Sovetskaya 16'),
('PetsHotel3','ul. Ilimskaya 33');

INSERT INTO ROOM (pet_type, number, hotel_id) VALUES 
('cat', 101, 1),
('dog', 201, 2),
('cat', 301, 3),
('cat', 102, 1),
('dog', 202, 2),
('cat', 302, 3),
('cat', 103, 1),
('dog', 203, 2),
('cat', 303, 3),
('cat', 104, 1),
('dog', 204, 2),
('cat', 304, 3),
('cat', 105, 1),
('dog', 205, 2),
('cat', 305, 3),
('cat', 106, 1),
('dog', 206, 2),
('cat', 306, 3),
('cat', 107, 1),
('dog', 207, 2),
('cat', 307, 3),
('cat', 108, 1),
('dog', 208, 2),
('cat', 308, 3),
('cat', 109, 1),
('dog', 209, 2),
('cat', 309, 3),
('cat', 110, 1),
('dog', 210, 2),
('cat', 310, 3),
('cat', 111, 1),
('dog', 211, 2),
('cat', 311, 3),
('cat', 112, 1),
('dog', 212, 2),
('cat', 312, 3),
('cat', 113, 1),
('dog', 213, 2),
('cat', 313, 3),
('cat', 114, 1),
('dog', 214, 2),
('cat', 314, 3),
('cat', 115, 1),
('dog', 215, 2),
('cat', 315, 3),
('cat', 116, 1),
('dog', 216, 2),
('cat', 316, 3),
('cat', 117, 1),
('dog', 217, 2),
('cat', 317, 3),
('cat', 118, 1),
('dog', 218, 2),
('cat', 318, 3),
('cat', 119, 1),
('dog', 219, 2),
('cat', 319, 3),
('cat', 120, 1),
('dog', 220, 2),
('cat', 320, 3);

INSERT INTO SEAT (room_id, is_free, description) VALUES 
(1, false, 'VIP seat'),
(2, true, 'Seat for sick pets'),
(3, false, 'Regular seat');

INSERT INTO CLIENT (user_id) VALUES 
(1),
(2),
(3);

INSERT INTO PET (name , type, weight, dieses, client_id) VALUES 
('Murzik','cat',5 ,'no diseases' ,1),
('Barbos','dog',5 ,'1 diseas' ,2),
('Aliy','dog',5 ,'2 diseases' ,3);

INSERT INTO EMPLOYEE (user_id, hotel_id, position, role) VALUES 
(4 , 1, 'Position 1', 'role 1' ),
(5 , 2, 'Position 3', 'role 2' ),
(6 , 3, 'Position 3', 'role 3' );

INSERT INTO BOOKING (seat_id , pet_id, employee_id, status, start_date, end_date, client_notes) VALUES 
(1 ,1 ,1,'In processing', '2021-12-07', '2021-12-27','wash my pet pls twice a day'),
(2 ,2 ,2,'In work', '2021-12-01', '2021-12-15','feed my pet pls once a week'),
(3, 3 ,3,'Ended up', '2021-11-26', '2021-12-06','no comm' );
