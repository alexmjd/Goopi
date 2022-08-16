DROP TABLE IF EXISTS product;

CREATE TABLE IF NOT EXISTS product (
	id int not null primary key auto_increment,
	name varchar(50) not null,
	price decimal(5,2) not null
)