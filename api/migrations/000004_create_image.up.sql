CREATE TABLE IF NOT EXISTS IMAGES 
(   id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    URL             TEXT NOT NULL ,
    ownerId         INTEGER 
);