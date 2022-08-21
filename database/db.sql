DROP TABLE IF EXISTS products;

CREATE TABLE IF NOT EXISTS products (
	id int not null primary key auto_increment,
	name varchar(50) not null,
	price decimal(5,2) not null
);

INSERT INTO product VALUES (1, "Table", 20.29);
INSERT INTO product VALUES (2, "Chaise", 11.1);
INSERT INTO product VALUES (3, "Lit", 499.49);
INSERT INTO product VALUES (4, "Bureau", 300);
INSERT INTO product VALUES (5, "Coffre", 999.99);