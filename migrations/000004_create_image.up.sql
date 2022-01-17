CREATE TABLE IF NOT EXISTS IMAGES 
(   id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    URL             TEXT NOT NULL ,
    ownerId         INTEGER 
);

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