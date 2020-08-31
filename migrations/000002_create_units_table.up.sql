CREATE TABLE units (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR NOT NULL,
  address VARCHAR NOT NULL,
  task_plan BIGINT
);

INSERT INTO 
  units (name, address, task_plan)
VALUES 
  ('Нет', 'NONE', 0),
  ('Каширский двор', 'Каширское шоссе, дом.19, корпус. 2', 500000);