CREATE TABLE manager (
     id INT(50) NOT NULL AUTO_INCREMENT,
     phone VARCHAR(200) NOT NULL,
     password VARCHAR(200) NOT NULL,
     PRIMARY KEY (id),
     UNIQUE (phone)
) default character set utf8 collate utf8_general_ci;

CREATE TABLE product (
     id INT(11) NOT NULL AUTO_INCREMENT,
     manager_id INT(50) DEFAULT NULL,
     category VARCHAR(200) NOT NULL,
     price VARCHAR(50) NOT NULL,
     name VARCHAR(200) NOT NULL,
     description VARCHAR(200) NOT NULL,
     size VARCHAR(10) NOT NULL,
     expired_date DATETIME,
     PRIMARY KEY(id),
     FOREIGN KEY (manager_id) REFERENCES manager(id)
) default character set utf8 collate utf8_general_ci;