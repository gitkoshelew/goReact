CREATE TYPE employee_position AS enum ('manager', 'employee','owner', 'admin') ;

UPDATE employee SET position = 'employee' WHERE employee_id = 1;
UPDATE employee SET position = 'manager' WHERE employee_id = 2;
UPDATE employee SET position = 'admin' WHERE employee_id = 3;