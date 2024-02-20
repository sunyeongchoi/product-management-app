CREATE TABLE manager (
     id INT(11) NOT NULL AUTO_INCREMENT,
     phone VARCHAR(20) NOT NULL,
     password VARCHAR(20) NOT NULL,
     PRIMARY KEY (id)
);

CREATE TABLE product (
     id INT(11) NOT NULL AUTO_INCREMENT,
     manager_id INT(11) DEFAULT NULL,
     category VARCHAR(20) NOT NULL,
     price VARCHAR(20) NOT NULL,
     name VARCHAR(20) NOT NULL,
     description VARCHAR(20) NOT NULL,
     size VARCHAR(20) NOT NULL,
     expired_date DATETIME,
     PRIMARY KEY(id),
     FOREIGN KEY (manager_id) REFERENCES manager(id)
);