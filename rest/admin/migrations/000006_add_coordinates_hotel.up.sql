ALTER TABLE HOTEL  
ADD COLUMN coordinates TEXT[];

UPDATE HOTEL SET coordinates = '{"53.89909164468815", "27.498996594142426"}' WHERE hotel_id = 1;
UPDATE HOTEL SET coordinates = '{"53.8945666689594", "27.544045611865155"}' WHERE hotel_id = 2;
UPDATE HOTEL SET coordinates = '{"53.889209895939615", "27.682389827072512"}' WHERE hotel_id = 3;
