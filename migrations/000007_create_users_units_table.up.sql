CREATE TABLE users_units (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT,
	unit_id BIGINT,
	CONSTRAINT fk_user_id
		FOREIGN KEY(user_id)
			REFERENCES users(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_unit_id
		FOREIGN KEY(unit_id)
			REFERENCES units(id)
				ON DELETE CASCADE
);

INSERT INTO users_units(user_id, unit_id) VALUES(1, 2);
