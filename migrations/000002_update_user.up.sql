CREATE TYPE sex AS enum ('male', 'female');

CREATE TYPE user_type AS enum ('client', 'employee', 'anonymous');

ALTER TABLE USERS
DROP COLUMN     account_id CASCADE,
ADD COLUMN      password        TEXT ,   
ADD COLUMN      role            user_type NOT NULL ,
ADD COLUMN      verified        BOOLEAN NOT NULL ,
ADD COLUMN      sex             sex NOT NULL,
ADD COLUMN      photo           TEXT;
;