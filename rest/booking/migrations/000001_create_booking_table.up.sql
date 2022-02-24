CREATE TYPE BookingStatus AS enum ('pending', 'in-progress', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS BOOKING 
(   id               serial PRIMARY key,
    seat_id          INTEGER NOT NULL,
    pet_id           INTEGER NOT NULL,
    employee_id      INTEGER NOT NULL,
    status           BookingStatus NOT NULL,
    start_date       DATE NOT NULL,
    end_date         DATE NOT NULL,
    paid             BOOLEAN NOT NULL,
    notes            TEXT
);