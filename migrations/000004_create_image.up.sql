CREATE TABLE IF NOT EXISTS IMAGES 
(   id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    ownerId         INTEGER 
);

INSERT INTO IMAGES (type, ownerId) VALUES 
('user', 1),
('user', 2),
('user', 3),
('user', 3)