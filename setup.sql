-- create postgresql test user, createdb mean user can create database
create user biz_test with password 'pwd' createdb;

-- create original database for test user
create database biz_test with owner biz_test;

-- drop database biz_test;
-- drop user biz_test;