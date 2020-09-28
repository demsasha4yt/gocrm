CREATE TABLE manufacturers (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	description VARCHAR,
	UNIQUE(name)
);

CREATE TABLE manufacturers_units (
	id BIGSERIAL PRIMARY KEY,
	manufacturer_id BIGINT,
	unit_id BIGINT,
	CONSTRAINT fk_inits
		FOREIGN KEY(unit_id)
				REFERENCES units(id)
					ON DELETE CASCADE,
	CONSTRAINT fk_manufacturers
		FOREIGN KEY(manufacturer_id)
				REFERENCES manufacturers(id)
					ON DELETE CASCADE,
	UNIQUE(manufacturer_id, unit_id)
);


CREATE TABLE categories (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	description VARCHAR,
	parent_id BIGINT CONSTRAINT not_id CHECK (id != parent_id),
	CONSTRAINT fk_categories
		FOREIGN KEY(parent_id)
				REFERENCES categories(id)
					ON DELETE CASCADE
);

CREATE TABLE categories_units (
		id BIGSERIAL PRIMARY KEY,
		category_id BIGINT,
		unit_id BIGINT,
		CONSTRAINT fk_categories
			FOREIGN KEY(category_id)
					REFERENCES categories(id)
						ON DELETE CASCADE,
		CONSTRAINT fk_units
			FOREIGN KEY(unit_id)
					REFERENCES units(id)
						ON DELETE CASCADE,
		UNIQUE(category_id, unit_id)
);

CREATE TABLE products (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	manufacturer_id BIGINT,
	CONSTRAINT fk_manufacturers
		FOREIGN KEY(manufacturer_id)
				REFERENCES manufacturers(id)
					ON DELETE CASCADE
);

CREATE TABLE products_units (
	id BIGSERIAL PRIMARY KEY,
	product_id BIGINT,
	unit_id BIGINT,
	CONSTRAINT fk_products
		FOREIGN KEY(product_id)
				REFERENCES products(id)
					ON DELETE CASCADE,
	CONSTRAINT fk_units
		FOREIGN KEY(unit_id)
				REFERENCES units(id)
					ON DELETE CASCADE,
		UNIQUE(product_id, unit_id)
);

CREATE TABLE products_categories (
	id BIGSERIAL PRIMARY KEY,
	product_id BIGINT,
	category_id BIGINT,
		CONSTRAINT fk_products
			FOREIGN KEY(product_id)
					REFERENCES products(id)
						ON DELETE CASCADE,
		CONSTRAINT fk_categories
			FOREIGN KEY(category_id)
					REFERENCES categories(id)
						ON DELETE CASCADE
)
