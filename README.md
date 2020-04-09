# HexMicroservice Exercise
Short URL Generator System

# How to run it?
##  Setup Environment Variables
### Use Redis
```
set URL_DB=redis
set REDIS_URL=redis://localhost:6379
```
### Use MongoDB
```
set MONGO_URL=mongo
set MONGO_DB=mongodb://localhost:27017
set MONGO_TIMEOUT=10
```
## Run Application
> Select one of database above that you want.Then
```bash
go run main.go
```