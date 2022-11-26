CREATE TABLE `example`.`m_users` (
	id INT(11) AUTO_INCREMENT NOT NULL, 
    name VARCHAR(30) NOT NULL,
    age INT(3) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE `example`.`m_items` (
	id INT(11) AUTO_INCREMENT NOT NULL, 
	name VARCHAR(10) NOT NULL,
    PRIMARY KEY (id)
);