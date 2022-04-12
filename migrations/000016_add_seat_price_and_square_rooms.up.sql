ALTER TABLE SEAT
ADD COLUMN price NUMERIC  NOT NULL DEFAULT 1;  

ALTER TABLE ROOM  
ADD COLUMN square NUMERIC  NOT NULL DEFAULT 1;  

UPDATE ROOM SET square = 3 WHERE room_id = 1;
UPDATE ROOM SET square = 3 WHERE room_id = 2;
UPDATE ROOM SET square = 3 WHERE room_id = 3;
UPDATE ROOM SET square = 3 WHERE room_id = 4;
UPDATE ROOM SET square = 3 WHERE room_id = 5;
UPDATE ROOM SET square = 3 WHERE room_id = 6;
UPDATE ROOM SET square = 3 WHERE room_id = 7;
UPDATE ROOM SET square = 3 WHERE room_id = 8;
UPDATE ROOM SET square = 3 WHERE room_id = 9;
UPDATE ROOM SET square = 3 WHERE room_id = 10;
UPDATE ROOM SET square = 3 WHERE room_id = 11;
UPDATE ROOM SET square = 3 WHERE room_id = 12;
UPDATE ROOM SET square = 3 WHERE room_id = 13;
UPDATE ROOM SET square = 3 WHERE room_id = 14;
UPDATE ROOM SET square = 3 WHERE room_id = 15;
UPDATE ROOM SET square = 3 WHERE room_id = 16;
UPDATE ROOM SET square = 3 WHERE room_id = 17;
UPDATE ROOM SET square = 3 WHERE room_id = 18;
UPDATE ROOM SET square = 3 WHERE room_id = 19;
UPDATE ROOM SET square = 3 WHERE room_id = 20;
UPDATE ROOM SET square = 3 WHERE room_id = 21;
UPDATE ROOM SET square = 3 WHERE room_id = 22;
UPDATE ROOM SET square = 3 WHERE room_id = 23;
UPDATE ROOM SET square = 3 WHERE room_id = 24;
UPDATE ROOM SET square = 3 WHERE room_id = 25;
UPDATE ROOM SET square = 3 WHERE room_id = 26;
UPDATE ROOM SET square = 3 WHERE room_id = 27;
UPDATE ROOM SET square = 3 WHERE room_id = 28;
UPDATE ROOM SET square = 3 WHERE room_id = 29;
UPDATE ROOM SET square = 3 WHERE room_id = 30;

UPDATE seat SET price = 1 WHERE room_id = 1;
UPDATE seat SET price = 12 WHERE room_id = 2;
UPDATE seat SET price = 13 WHERE room_id = 3;
