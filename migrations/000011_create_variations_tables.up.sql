CREATE TABLE variations (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	description VARCHAR,
	price INT CONSTRAINT positive_price CHECK (price > 0)
);

CREATE TABLE products_variations (
	id BIGSERIAL PRIMARY KEY,
	product_id BIGINT,
	variation_id BIGINT,
	CONSTRAINT fk_products
		FOREIGN KEY(product_id)
			REFERENCES products(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_variations
		FOREIGN KEY(variation_id)
			REFERENCES variations(id)
				ON DELETE CASCADE
);

CREATE TABLE variations_options_values (
	id BIGSERIAL PRIMARY KEY,
	variation_id BIGINT,
	options_value_id BIGINT,
	CONSTRAINT fk_variations
		FOREIGN KEY(variation_id)
			REFERENCES variations(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_options_values
		FOREIGN KEY(options_value_id)
			REFERENCES options_values(id)
				ON DELETE CASCADE
);
