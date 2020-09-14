CREATE TABLE customers (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	phone VARCHAR,
	address VARCHAR,
	created_by BIGINT,
	CONSTRAINT fk_users
		FOREIGN KEY(created_by)
			REFERENCES users(id)
				ON DELETE CASCADE
);
