ALTER TABLE SEAT
DROP COLUMN     is_free     CASCADE,
ADD COLUMN      rent_from   DATE,   
ADD COLUMN      rent_to     DATE;

UPDATE seat SET rent_from = '2022-01-25', rent_to = '2022-01-30' WHERE seat_id = 1;
UPDATE seat SET rent_from = '2022-02-10', rent_to = '2022-02-20' WHERE seat_id = 2;
UPDATE seat SET rent_from = '2022-02-14', rent_to = '2000-03-08' WHERE seat_id = 3;