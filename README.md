Rest API build using FIBER FRAMEWORK. 

This project was made using fiber framework from Golang, and it's a simple CRUD with MongoDB, but
you can use any database you want, just need to change the connection and the methods.

In this project we have the following technologies:
- Fiber Framework
- MongoDB
- Docker
- JWT Authentication
- Encryption to hash
- Unit tests 
- MVC - Model View Controller - Pattern

This project was guided by a tutorial from @Hunconding: https://www.youtube.com/watch?v=vxDqv6BKZCw&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr

But have some differences:
- Use Fiber Instead GinGonic
- Use Docker
- Different organization from some fields and classes and tests.

## Installation

```bash
$ docker compose up -d ### This command will start the database on port 3000 using a container from docker-compose.yml
$ go run main.go ### This command will start the application from main.go
```

## ENDPOINTS to call the API:

### Create user @POST -> localhost:3000/createUser

```bash

curl --location 'localhost:3000/createUser' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "joao@gmail.com",
    "age": 21,
    "password": "joao",
    "name": "joao"
}'
```

### Login with created user @POST -> localhost:3000/login

```bash
curl --location 'localhost:3000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "joao@gmail.com",
    "password": "joao"
}'

```
When you do login, you will receive a token, copy it and paste on the next endpoints, to test the authentication.

### Find by userId @GET -> localhost:3000/getUserById/:userId

```bash
curl --location 'localhost:3000/getUserById/:userId' \
--header 'Authorization: :$token' \
--data ''
```

### Find by email @GET -> localhost:3000/getUserByEmail/:userEmail

```bash
curl --location 'localhost:3000/getUserByEmail/:userEmail' \
--header 'Authorization: :$token \
--data ''
```

### Update 1 or more fields @PUT -> localhost:3000/updateUser/:userId

```bash
curl --location --request PUT 'localhost:3000/updateUser/:userId' \
--header 'Content-Type: application/json', 'Authorization: :$token' \
--data '{
    "age": 50
}
'
```

### Delete some entry @DELETE -> localhost:3000/:userId

```bash
curl --location --request DELETE 'localhost:3000/userId'
--header 'Authorization: :$token' \
--data ''
```

## Next Goals

- Change from MVC to a Hexagonal Pattern
- Implement a "base" user and password to login, and then, create a user and use it.
- Implement Swagger.
- Implement some Cloud Service, like a S3 using Localstack.
- Implement a CI/CD using Github Actions.
- Implement a MakeFile to run the project.
- Implement Integration Tests for S3 and Repository.
- Implement a FrontEnd using ReactJS.
- Make a deployment for all the enviromment.
