ALTER TABLE SEAT
DROP COLUMN     description     CASCADE,

ALTER TABLE ROOM  
ADD COLUMN description TEXT  NOT NULL DEFAULT '';  

UPDATE ROOM SET description = 'description' WHERE id = 1;
UPDATE ROOM SET description = 'description' WHERE id = 2;
UPDATE ROOM SET description = 'description' WHERE id = 3;
