INSERT INTO USERS (email, password, role, verified, first_name , surname, middle_name, sex, date_of_birth, address, phone, photo) VALUES 
('ivan@mail.ru', 'password1', 'client', true, 'Ivan','Ivanov','Ivanovich', 'male', '2000-01-01', 'Minsk Pr. Nezavisimosti 22-222' ,'+375-29-154-89-33', 'PhotoURL...'),
('petr@mail.ru', 'password2', 'client', true, 'Petr','Petrov','Petrovich', 'male', '1999-03-13 ', 'Minsk Pr. Pobediteley 11-111' ,'+375-29-159-11-78', 'PhotoURL...'),
('liza@mail.ru', 'password3', 'client', true, 'Maria','Petrova','Evgenievna', 'female', '2001-11-11', 'Minsk Pr. Pobediteley 111-111' ,'+375-29-655-99-14', 'PhotoURL...'),
('olga@mail.ru', 'password4', 'employee', true, 'Olga','Oleeva','Vladimirovna', 'female', '2001-01-01', 'Minsk ul. Nesterova 1-32' ,'+375-29-675-00-00', 'PhotoURL...'),
('leha@mail.ru', 'password5', 'employee', true, 'Vladzimir','Sakhonchik','Alekseevish', 'male', '1998-02-11', 'Minsk Partizanskiy 124-1' ,'+375-29-111-22-33', 'PhotoURL...'),
('miha@mail.ru', 'password6', 'employee', true, 'Mikhail','Valevach','Dmitrievich', 'male', '1997-01-12', 'Minsk Pr. Dzerjinskogo 01-01' ,'+375-29-777-55-44', 'PhotoURL...');

INSERT INTO PET (name , type, weight, diseases, user_id) VALUES 
('Murzik','cat',5 ,'no dieses' ,1),
('Barbos','dog',5 ,'1 dieses' ,2),
('Aliy','dog',5 ,'2 dieses' ,3);

INSERT INTO IMAGES (type, URL, ownerId) VALUES 
('pet', 'volume\images\pets\cat1.jpg', 1),
('pet', 'volume\images\pets\cat2.jpg', 2),
('pet', 'volume\images\pets\cat3.jpg', 3),
('pet', 'volume\images\pets\dog1.jpg', 3),
('pet', 'volume\images\pets\dog2.jpg', 3),
('pet', 'volume\images\pets\dog3.jpg', 3);