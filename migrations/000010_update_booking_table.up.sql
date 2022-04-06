ALTER TABLE BOOKING
ADD COLUMN IF NOT EXISTS transactionId INTEGER,
ADD COLUMN IF NOT EXISTS paid BOOLEAN;

UPDATE booking SET transactionId = 0, paid = false WHERE booking_id = 1;
UPDATE booking SET transactionId = 1, paid = true WHERE booking_id = 2;
UPDATE booking SET transactionId = 2, paid = true WHERE booking_id = 3;
