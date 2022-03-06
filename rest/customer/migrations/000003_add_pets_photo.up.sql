ALTER TABLE PET  
ADD COLUMN      photo     text;

UPDATE pet SET photo = '//photo1' WHERE id = 1;
UPDATE pet SET photo = '//photo2' WHERE id = 2;
UPDATE pet SET photo = '//photo3' WHERE id = 3;