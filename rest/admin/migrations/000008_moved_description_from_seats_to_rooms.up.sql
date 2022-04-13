ALTER TABLE SEAT
DROP COLUMN     description     CASCADE;

ALTER TABLE ROOM  
ADD COLUMN description TEXT  NOT NULL DEFAULT '';  

UPDATE ROOM SET description = 'description' WHERE room_id = 1;
UPDATE ROOM SET description = 'description' WHERE room_id = 2;
UPDATE ROOM SET description = 'description' WHERE room_id = 3;
