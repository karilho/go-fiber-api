Rest API build using FIBER FRAMEWORK. [IN PROGRESS - 40%]

This project was made using fiber framework from Golang, and it's a simple CRUD with DynamoDB, but
you can use any database you want, just need to change the connection and the methods.

In this project we have the following technologies:
- Fiber Framework
- DynamoDB
- Docker
- JWT - Not implemented yet
- Encryption to hash
- Unit tests - Not implemented yet
- MVC - Model View Controller - Pattern
- Swagger - Not implemented yet

## Installation

```bash
Run -docker run --name -d -p 27017:27017 mongo
Run - docker compose up -d [not working right now]
```
Run - go run main.go

First will setup dynamodb and then will start the server on port 3000 using a container from docker-compose.yml
After you will use the following endpoints:

localhost:3000/createUser

JSON EXAMPLE ->
{
    "email": "json@gmail.com",
    "name": "json",
    "password": "json",
    "age": 20
}

