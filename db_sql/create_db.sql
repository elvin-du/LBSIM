create database LBSIM;
use LBSIM;
create table LBSIM.users(UserID INT PRIMARY KEY AUTO_INCREMENT,Name VARCHAR(50) NOT NULL UNIQUE, password VARCHAR(100) NOT NULL, groupid int);
create table LBSIM.users_relation(ID INT PRIMARY KEY AUTO_INCREMENT,UserAID INT NOT NULL,UserBID INT NOT NULL,Relation INT NOT NULL);

create table LBSIM.friend_requests(ID INT PRIMARY KEY AUTO_INCREMENT,UserAID INT NOT NULL,UserBID INT NOT NULL,remark varchar(200));
create table LBSIM.friend_responses(ID INT PRIMARY KEY AUTO_INCREMENT,UserAID INT NOT NULL,UserBID INT NOT NULL,remark varchar(200));
create table LBSIM.friend_offline_msgs(ID INT PRIMARY KEY AUTO_INCREMENT,UserAID INT NOT NULL,UserBID INT NOT NULL,msgs varchar(1000));
create table LBSIM.groups(groupid INT PRIMARY KEY AUTO_INCREMENT,introduction varchar(500));


