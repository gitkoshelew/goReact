CREATE TABLE IF NOT EXISTS CLIENT 
(   id          serial PRIMARY key , 
    user_id     INTEGER REFERENCES USERS(id) ON DELETE CASCADE NOT NULL 
);