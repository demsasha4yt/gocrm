CREATE TABLE products_options_values (
	id BIGSERIAL PRIMARY KEY,
	product_id BIGINT,
	options_value_id BIGINT,
	CONSTRAINT fk_products
		FOREIGN KEY(product_id)
			REFERENCES products(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_options_values
		FOREIGN KEY(options_value_id)
			REFERENCES options_values(id)
				ON DELETE CASCADE
);
