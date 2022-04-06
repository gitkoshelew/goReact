ALTER TABLE SEAT
ADD COLUMN price NUMERIC  NOT NULL DEFAULT 1;  

ALTER TABLE ROOM  
ADD COLUMN square NUMERIC  NOT NULL DEFAULT 1;  

UPDATE ROOM SET square = 3 WHERE room_id = 1;
UPDATE ROOM SET square = 23.5 WHERE room_id = 2;
UPDATE ROOM SET square = 33.5 WHERE room_id = 3;


UPDATE seat SET price = 1 WHERE seat_id = 1;
UPDATE seat SET price = 1.2 WHERE seat_id = 2;
UPDATE seat SET price = 13.5 WHERE seat_id = 3;