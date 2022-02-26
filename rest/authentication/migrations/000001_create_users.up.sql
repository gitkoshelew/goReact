CREATE TYPE user_type AS enum ('client', 'employee', 'anonymous');

CREATE TABLE IF NOT EXISTS USERS 
(   id              serial PRIMARY key,
    email           CHARACTER VARYING(50) NOT NULL,
    Password        TEXT NOT NULL,
    verified        BOOLEAN NOT NULL,
    role            user_type NOT NULL 
);