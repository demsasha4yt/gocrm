CREATE TABLE users_operations (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR NOT NULL,
  access_level BIGINT,
  CONSTRAINT fk_access_levels
    FOREIGN KEY(access_level) 
      REFERENCES access_levels(id)
      ON DELETE SET DEFAULT
);