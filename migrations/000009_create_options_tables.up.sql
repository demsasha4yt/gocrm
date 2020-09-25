CREATE TABLE options_types (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	is_soft BOOLEAN
);

CREATE TABLE options (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	description VARCHAR,
	options_type_id BIGINT,
	CONSTRAINT fk_options_types
		FOREIGN KEY(options_type_id)
			REFERENCES options_types(id)
				ON DELETE CASCADE
);

CREATE TABLE options_values (
	id BIGSERIAL PRIMARY KEY,
	value VARCHAR,
	image VARCHAR,
	option_id BIGINT,
	options_type_id BIGINT,
	CONSTRAINT fk_options
		FOREIGN KEY(option_id)
			REFERENCES options(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_options_types
		FOREIGN KEY(options_type_id)
			REFERENCES options_types(id)
				ON DELETE CASCADE
);

CREATE TABLE softs_categories (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	value INT NOT NULL,
	CONSTRAINT uniq_value 
		UNIQUE(value)
);

CREATE TABLE options_softs (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR,
	image VARCHAR,
	options_value_id BIGINT,
	manufacturer_id BIGINT,
	soft_category_id BIGINT,
	CONSTRAINT fk_options_values
		FOREIGN KEY(options_value_id)
			REFERENCES options_values(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_manufacturers
		FOREIGN KEY(manufacturer_id)
			REFERENCES manufacturers(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_softs_categories
		FOREIGN KEY(soft_category_id)
			REFERENCES softs_categories(id)
				ON DELETE CASCADE
);
