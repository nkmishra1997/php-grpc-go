# Readme
- A simple boilerplate to isllustrate client interaction with various microservices using a GRPC server.
- PHP Client/ Go Client interacting with a microservice (Go server) using GRPC.
- Simple CRUD app for restaurant listing.
## Setup

- Setup `grpc` and `protoc-gen-go`
- Run 
```bash
$ make compile_protos
```
- Start the go grpc server
```bash
$ cd server && go run main.go
```
- Start the go grpc client (If you wish to use Go client)
```bash
$ cd go_client && go run main.go
```
- Start the php grpc client (If you wish to use PHP client)
```bash
$ cd php_client && php -S localhost:3000 index.php
```

## Running with Docker

- Run
```bash
$ docker-compose up --build -d
```
