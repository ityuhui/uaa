/*
- mysql -h host -uuser_name -ppassword < db.sql

- kubectl exec mysql-deployment-5f7776794b-6pqlj -- mysql -uuaa -ppassword

- use phpmyadmin
*/

USE uaa;

CREATE TABLE users 
(
    id INTEGER(32) auto_increment not null primary key,
    name VARCHAR(32) not null,
    password VARCHAR(32) not null
);