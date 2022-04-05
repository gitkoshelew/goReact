ALTER TABLE SEAT
DROP COLUMN     is_free     CASCADE,
ADD COLUMN      rent_from   DATE[] DEFAULT '{}',   
ADD COLUMN      rent_to     DATE[] DEFAULT '{}';

UPDATE seat SET rent_from = '{2022-01-25, 2022-04-25}', rent_to = '{2022-01-30, 2022-04-30}' WHERE id = 1;
UPDATE seat SET rent_from = '{2022-02-10, 2022-05-25}', rent_to = '{2022-02-20, 2022-05-30}' WHERE id = 2;
UPDATE seat SET rent_from = '{2022-02-14, 2022-06-25}', rent_to = '{2022-03-08, 2022-06-30}' WHERE id = 3;