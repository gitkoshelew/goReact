CREATE TABLE IF NOT EXISTS PERMISSIONS 
(   id              serial PRIMARY key, 
    name      CHARACTER VARYING(30) NOT NULL ,
    description         TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS permissions_employees
(   permissions_id            INTEGER REFERENCES PERMISSIONS(id) ON DELETE CASCADE NOT NULL  ,
    employee_id            INTEGER REFERENCES EMPLOYEE(id) ON DELETE CASCADE NOT NULL     
);

INSERT INTO PERMISSIONS (name, description) VALUES
('read_user','ability to read a user'), 
('creat_user','ability to create a user'),
('delete_user','ability to delete a user'),
('update_user','ability to update a user'),
('read_hotel','ability to read a hotel'), 
('creat_hotel','ability to create a hotel'),
('delete_hotel','ability to delete a hotel'),
('update_hotel','ability to update a hotel'),
('read_bookings','ability to read a booking'), 
('creat_bookings','ability to create a booking'),
('delete_bookings','ability to delete a booking'),
('update_bookings','ability to update a booking'),
('read_employee','ability to read a employee'), 
('creat_employee','ability to create a employee'),
('delete_employee','ability to delete a employee'),
('update_employee','ability to update a employee'),
('read_pet','ability to read a pet'), 
('creat_pet','ability to create a pet'),
('delete_pet','ability to delete a pet'),
('update_pet','ability to update a pet'),
('read_rooms','ability to read a room'), 
('creat_rooms','ability to create a room'),
('delete_rooms','ability to delete a room'),
('update_rooms','ability to update a room'),
('read_seats','ability to read a seat'), 
('creat_seats','ability to create a seat'),
('delete_seats','ability to delete a seat'),
('update_seats','ability to update a seat');

INSERT INTO permissions_employees (permissions_id, employee_id) VALUES 
(1,3),
(2,3),
(3,3),
(4,3),
(5,3),
(6,3),
(7,3),
(8,3),
(9,3),
(10,3),
(11,3),
(12,3),
(13,3),
(14,3),
(15,3),
(16,3),
(17,3),
(18,3),
(19,3),
(20,3),
(21,3),
(22,3),
(23,3),
(24,3),
(25,3),
(26,3),
(27,3),
(28,3),
(1,1),
(2,1),
(3,1),
(4,1),
(5,1),
(6,1),
(7,1),
(8,1),
(1,2),
(2,2),
(3,2),
(4,2),
(5,2),
(6,2),
(9,2),
(11,2);