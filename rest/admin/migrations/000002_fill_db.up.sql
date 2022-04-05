

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
('dog', 202, 2),
('cat', 303, 3);

INSERT INTO SEAT (room_id, is_free, description) VALUES 
(1, false, 'VIP seat'),
(2, true, 'Seat for sick pets'),
(3, false, 'Regular seat');

INSERT INTO PET (name , type, weight, diseases, user_id) VALUES 
('Murzik','cat',5 ,'no dieses' ,1),
('Barbos','dog',5 ,'1 dieses' ,2),
('Aliy','dog',5 ,'2 dieses' ,3);

INSERT INTO EMPLOYEE (user_id, hotel_id, position ) VALUES 
(4 , 1, 'manager' ),
(5 , 2, 'employee' ),
(6 , 3, 'admin' );

INSERT INTO BOOKING (seat_id , pet_id, employee_id, status, start_date, end_date, notes) VALUES 
(1 ,1 ,1,'completed', '2021-12-07', '2021-12-27','wash my pet pls twice a day'),
(2 ,2 ,2,'cancelled', '2021-12-01', '2021-12-15','feed my pet pls once a week'),
(3, 3 ,3,'pending', '2021-11-26', '2021-12-06','no comm' );

INSERT INTO IMAGES (type, ownerId) VALUES 
('pet',  1),
('pet',  2),
('pet',  3),
('pet',  3),
('pet',  3),
('pet',  3),
('room', 1),
('room', 2),
('room', 3);