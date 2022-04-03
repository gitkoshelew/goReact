ALTER TABLE SEAT
DROP COLUMN     is_free     CASCADE,
ADD COLUMN      rent_from   DATE[] DEFAULT '{}',   
ADD COLUMN      rent_to     DATE[] DEFAULT '{}';

UPDATE seat SET rent_from = '{2022-01-25}', rent_to = '{2022-01-30}' WHERE id = 1;
UPDATE seat SET rent_from = '{2022-02-10}', rent_to = '{2022-02-20}' WHERE id = 2;
UPDATE seat SET rent_from = '{2022-02-14}', rent_to = '{2022-03-08}' WHERE id = 3;