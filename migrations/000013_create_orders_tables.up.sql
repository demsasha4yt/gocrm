CREATE TABLE orders (
	id BIGSERIAL PRIMARY KEY,
	customer_id BIGINT,
	unit_id BIGINT,
	created_by BIGINT,
	email VARCHAR,
	phone VARCHAR,
	shipping_date TIMESTAMP,
	CONSTRAINT fk_customers
		FOREIGN KEY(customer_id)
			REFERENCES customers(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_users
		FOREIGN KEY(created_by)
			REFERENCES users(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_unit_id
		FOREIGN KEY(unit_id)
			REFERENCES units(id)
				ON DELETE CASCADE
);

CREATE TABLE orders_variations (
	id BIGSERIAL PRIMARY KEY,
	order_id BIGINT,
	variation_id BIGINT,
	CONSTRAINT fk_orders
		FOREIGN KEY(order_id)
			REFERENCES orders(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_variations
		FOREIGN KEY(variation_id)
			REFERENCES variations(id)
				ON DELETE CASCADE
);
