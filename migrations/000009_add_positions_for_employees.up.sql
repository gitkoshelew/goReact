CREATE TYPE employee_position AS enum ('manager', 'employee','owner', 'admin') ;

UPDATE employee SET position = 'employee' WHERE id = 1;
UPDATE employee SET position = 'manager' WHERE id = 2;
UPDATE employee SET position = 'admin' WHERE id = 3;