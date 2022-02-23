DROP DATABASE IF EXISTS Products;
CREATE DATABASE Products;
USE Products;
CREATE TABLE Products(
id int NOT NULL AUTO_INCREMENT,
name varchar(50),
type varchar(50),
PRIMARY KEY(id));
INSERT INTO Products VALUES(1,'Reebok','Bats');
INSERT INTO Products VALUES(2,'Mehfil','Biryani');