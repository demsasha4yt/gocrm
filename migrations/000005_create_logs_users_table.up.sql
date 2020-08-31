CREATE TABLE logs_users (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR,
  operation_id BIGINT,
  CONSTRAINT fk_operation_id
    FOREIGN KEY(operation_id) 
      REFERENCES users_operations(id)
      ON DELETE CASCADE
);