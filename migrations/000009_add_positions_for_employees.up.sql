CREATE TYPE employee_position AS enum ('manager', 'employee','owner', 'admin') ;
ALTER TABLE EMPLOYEE
ALTER COLUMN position type employee_position USING 'manager';