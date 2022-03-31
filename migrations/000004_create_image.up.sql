CREATE TABLE IF NOT EXISTS IMAGES 
(   id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    URL             TEXT NOT NULL ,
    ownerId         INTEGER 
);

INSERT INTO IMAGES (type, URL, ownerId) VALUES 
('user', '2022-03-31-05-04-34', 1),
('user', '2022-03-31-05-04-34', 2),
('user', '2022-03-31-05-04-34', 3),
('user', '2022-03-31-05-04-34', 3)