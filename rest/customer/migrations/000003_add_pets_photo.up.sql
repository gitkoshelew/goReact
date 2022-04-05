ALTER TABLE PET  
ADD COLUMN      photoURL     text;

UPDATE pet SET photoURL = '//photo1' WHERE id = 1;
UPDATE pet SET photoURL = '//photo2' WHERE id = 2;
UPDATE pet SET photoURL = '//photo3' WHERE id = 3;