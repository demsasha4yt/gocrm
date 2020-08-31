CREATE TABLE users (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  login VARCHAR(128) NOT NULL,
  password VARCHAR(128) NOT NULL,
  email VARCHAR(255) NOT NULL,
  first_name VARCHAR(64),
  last_name VARCHAR(64),
  third_name VARCHAR(64),
  access_level BIGINT DEFAULT 0,
  unit_id BIGINT DEFAULT 1,
  last_login TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(login),
  UNIQUE(email),
  CONSTRAINT fk_access_levels
    FOREIGN KEY(access_level) 
      REFERENCES access_levels(id)
      ON DELETE SET DEFAULT,
  CONSTRAINT fk_units
    FOREIGN KEY(unit_id)
      REFERENCES units(id)
      ON DELETE CASCADE
);