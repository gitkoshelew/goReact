CREATE TYPE sex AS enum ('male', 'female');

CREATE TYPE user_type AS enum ('client', 'employee', 'anonymous');

 
CREATE TABLE IF NOT EXISTS USERS 
(   id              serial PRIMARY key, 
    first_name      CHARACTER VARYING(30) NOT NULL ,
    surname         CHARACTER VARYING(30) NOT NULL ,
    middle_name     CHARACTER VARYING(30) ,    
    email           CHARACTER VARYING(30) ,
    date_of_birth   DATE NOT NULL ,
    address         TEXT NOT NULL ,
    phone           CHARACTER VARYING(30) NOT NULL ,
    password        TEXT ,
    role            user_type NOT NULL ,
    verified        BOOLEAN NOT NULL ,
    sex             sex NOT NULL,          
    photo           TEXT
);


CREATE TABLE IF NOT EXISTS PET 
(   id           serial PRIMARY key ,
    name         CHARACTER VARYING(30) NOT NULL ,
    type         CHARACTER VARYING(30) NOT NULL ,
    weight       SMALLINT NOT NULL ,
    diseases     TEXT,
    user_id      INTEGER REFERENCES USERS(id) ON DELETE CASCADE NOT NULL 
);


CREATE TABLE IF NOT EXISTS IMAGES 
(   id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    URL             TEXT NOT NULL ,
    ownerId         INTEGER 
);