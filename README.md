# Readme

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
- Start the go grpc client
```bash
$ cd go_client && go run main.go
```
- Start the php grpc client
```bash
$ cd php_client && php -S localhost:3000 index.php
```

## Running with Docker

- Run
```bash
$ docker-compose up --build -d
```