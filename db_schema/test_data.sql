/*
- mysql -h host -uuser_name -ppassword < db.sql

- kubectl exec mysql-deployment-5f7776794b-6pqlj -- mysql -uuaa -ppassword

- use phpmyadmin
*/

USE uaa;

insert into users values('','user1','password');
insert into users values('','user2','password');
insert into users values('','user3','password');