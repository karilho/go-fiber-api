Rest API build using FIBER FRAMEWORK. 

This project was made using fiber framework from Golang, and it's a simple CRUD with DynamoDB, but
you can use any database you want, just need to change the connection and the methods.

In this project we have the following technologies:
- Fiber Framework
- DynamoDB
- Docker
- JWT Authentication
- Encryption to hash
- Unit tests 
- MVC - Model View Controller - Pattern
- Swagger - Not implemented yet

## Installation

```bash
Run - docker compose up -d 
Run - go run main.go
```

First we will setup dynamodb and then will start the server on port 3000 using a container from docker-compose.yml, and after
run the main.go file.

## ENDPOINTS to call the API:

### CREATE USER @POST -> localhost:3000/createUser

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

### LOGIN WITH CREATED USER @POST -> localhost:3000/login

```bash
curl --location 'localhost:3000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "joao@gmail.com",
    "password": "joao"
}'

When u login, u will receive a token, copy it and paste on the next endpoints, to test the authentication.
```

### FIND BY USERID @GET -> localhost:3000/getUserById/:userId

```bash
curl --location 'localhost:3000/getUserById/:userId' \
--header 'Authorization: :TOKENVALUE' \
--data ''
```

### FIND BY EMAIL @GET -> localhost:3000/getUserByEmail/:userEmail

```bash
curl --location 'localhost:3000/getUserByEmail/:userEmail' \
--header 'Authorization: :TOKENVALUE' \
--data ''
```

### UPDATE 1 or MORE FIELDS @PUT -> localhost:3000/updateUser/:userId

```bash
curl --location --request PUT 'localhost:3000/updateUser/:userId' \
--header 'Content-Type: application/json', 'Authorization: :TOKENVALUE' \
--data '{
    "age": 50
}
'
```

### DELETE @DELETE -> localhost:3000/:userId

```bash
curl --location --request DELETE 'localhost:3000/userId'
--header 'Authorization: :TOKENVALUE' \
--data ''
```

## Next Goals

- Implement a "base" user and password to login, and then, create a user and use it.
- Implement Swagger.
- Implement some Cloud Service, like a S3 using Localstack.
- Implement a CI/CD using Github Actions.
- Implement a MakeFile to run the project.
- Implement Integration Tests for S3 and Repository.
- Implement a FrontEnd using ReactJS.
- Make a deployment for all the enviromment.