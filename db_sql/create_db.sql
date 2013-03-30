create database LBSIM;
create table LBSIM.users(userid INT PRIMARY KEY AUTO_INCREMENT,name VARCHAR(50) NOT NULL UNIQUE, password VARCHAR(100) NOT NULL, groupid int);
create table LBSIM.users_relation(id INT PRIMARY KEY AUTO_INCREMENT,useraid INT NOT NULL,userbid INT NOT NULL,relation INT NOT NULL);

create table LBSIM.friend_requests(id INT PRIMARY KEY AUTO_INCREMENT,useraid INT NOT NULL,userbid INT NOT NULL,remark varchar(200));
create table LBSIM.friend_responses(id INT PRIMARY KEY AUTO_INCREMENT,useraid INT NOT NULL,userbid INT NOT NULL,remark varchar(200));
create table LBSIM.friend_offline_msgs(id INT PRIMARY KEY AUTO_INCREMENT,useraid INT NOT NULL,userbid INT NOT NULL,msgs varchar(1000));
create table LBSIM.groups(groupid INT PRIMARY KEY AUTO_INCREMENT,introduction varchar(500));


