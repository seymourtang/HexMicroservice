# HexMicroservice Exercise
Short URL Generator System

# Getting Started
##  Setup Environment Variables
#### Use Redis
```
set URL_DB=redis
set REDIS_URL=redis://localhost:6379
```
#### Use MongoDB
```
set URL_DB=mongo
set MONGO_URL=mongo
set MONGO_DB=mongodb://localhost:27017
set MONGO_TIMEOUT=10
```
#### Use MySQL
> Execute the SQL script below to create your database.
```sql
create table ShortURL
(
    Id        int auto_increment primary key,
    Code      varchar(256) null,
    URL       varchar(256) null,
    CreatedAt mediumtext   null
);
```
Then setup environment variables as other ways do:
```
set URL_DB=mysql
set MYSQL_URL=root:123456@tcp(localhost:3306)/URLDB?charset=utf8
set MYSQL_DBNAME=ShortURL
```
## Run Application
> Select one of database above that you want.Then
```bash
go run main.go
```