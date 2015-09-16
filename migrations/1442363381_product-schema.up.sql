CREATE TABLE products (
	id BIGSERIAL PRIMARY KEY NOT NULL,
	name text NOT NULL,
	roast_date date NOT NULL,
	description text NOT NULL,
	price numeric(18,2) NOT NULL
);

CREATE UNIQUE INDEX idx_products_name on products (name);