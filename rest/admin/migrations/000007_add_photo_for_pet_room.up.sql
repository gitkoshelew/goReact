ALTER TABLE PET  
ADD COLUMN      photo     text;

ALTER TABLE ROOM  
ADD COLUMN      photo     text;

UPDATE pet SET photo = '//photo1' WHERE pet_id = 1;
UPDATE pet SET photo = '//photo2' WHERE pet_id = 2;
UPDATE pet SET photo = '//photo3' WHERE pet_id = 3;

UPDATE ROOM SET photo = '//photo1' WHERE room_id = 1;
UPDATE ROOM SET photo = '//photo2' WHERE room_id = 2;
UPDATE ROOM SET photo = '//photo3' WHERE room_id = 3;