 SET NAMES utf8mb4 ;
create user "suke"@"%%" identified by "";
grant all privileges on *.* to "suke"@"%%";
 DROP database IF EXISTS `test`;
create database `test` default character set utf8mb4 collate utf8mb4_unicode_ci; 
-- 切换到test_data数据库
use test; 