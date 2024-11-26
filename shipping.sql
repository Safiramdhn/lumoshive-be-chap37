CREATE TABLE shipping_services (
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	cost_rate DECIMAL(10,2) DEFAULT 0
);

INSERT INTO shipping_services (name, cost_rate) VALUES
('JNE', 4000.00),
('TIKI', 5000.00),
('Pos Indonesia', 6500.00);