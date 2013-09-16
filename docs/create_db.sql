create database LBSIM;
create table LBSIM.users(id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(50) NOT NULL UNIQUE, password VARCHAR(100) NOT NULL);
create table LBSIM.users_location(id INT PRIMARY KEY AUTO_INCREMENT, longitude float NOT NULL, latitude float NOT NULL, time date NOT NULL);
create table LBSIM.users_relation(id INT PRIMARY KEY AUTO_INCREMENT, useraid INT NOT NULL, userbid INT NOT NULL, relation INT NOT NULL);
create table LBSIM.friend_requests(id INT PRIMARY KEY AUTO_INCREMENT, useraid INT NOT NULL, userbid INT NOT NULL, genus INT NOT NULL , remark varchar(200), read boolean NOT NULL, time date NOT NULL);
create table LBSIM.friend_responses(id INT PRIMARY KEY AUTO_INCREMENT, useraid INT NOT NULL, userbid INT NOT NULL, remark varchar(200), response boolean NOT NULL, time date NOT NULL);
create table LBSIM.friend_msgs(id INT PRIMARY KEY AUTO_INCREMENT, useraid INT NOT NULL, userbid INT NOT NULL, msgs varchar(1000), read boolean NOT NULL, time date NOT NULL);

