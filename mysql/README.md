# MYSQL and Go

Adapted from: https://github.com/jboursiquot/workshop-go-database

## Setup
- [Install Go](https://github.com/WWG-Denver-Boulder/getting_started)
- [Install MySQL](https://dev.mysql.com/downloads/) or `brew install mysql`

## Start MySQL Server
- Will vary by install method

Ex:
- `brew tap homebrew/services`
- `brew services start mysql`
- `mysql -uroot -p` enter password

## Create Database
- `CREATE DATABSE job_board_db;`
- `USE job_board_db;`

## Create Tables
- `CREATE TABLE user_roles_t (login varchar(32) not null, type varchar(11) not null);`
- `show tables;`
- `desc user_roles_t;`

## Add Data
- `mysql -u root -p < load_data.sql`

## Delete Table or DB
- `DROP TABLE user_roles_t;`
- `DROP DATABASE job_board_db;`

## Resources
- [database/sql](https://golang.org/pkg/database/sql/)
- [mysql driver](https://github.com/go-sql-driver/mysql/)