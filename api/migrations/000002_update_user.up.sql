CREATE TYPE sex AS enum ('male', 'female');

CREATE TYPE user_type AS enum ('client', 'employee', 'anonymous');

ALTER TABLE USERS
ADD COLUMN      password        TEXT ,   
ADD COLUMN      role            user_type NOT NULL ,
ADD COLUMN      verified        BOOLEAN NOT NULL ,
ADD COLUMN      sex             INTEGER NOT NULL, /* 0 = male, 1 = female */
ADD COLUMN      photo           TEXT,
DROP COLUMN     account_id CASCADE;