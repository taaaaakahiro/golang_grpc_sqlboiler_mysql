CREATE TABLE `example`.`m_user` (
	id INT(11) AUTO_INCREMENT NOT NULL, 
    name VARCHAR(30) NOT NULL,
    age INT(3) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE `example`.`m_item` (
	id INT(11) AUTO_INCREMENT NOT NULL, 
	name VARCHAR(10) NOT NULL,
    PRIMARY KEY (id)
);