# Dockerfile for gRPC Go
FROM golang:1.8

LABEL maintainer="Nikhil Mishra <nikhil.mishra@zomato.com>"

WORKDIR $GOPATH/src/github.com/nkmishra1997/php-grpc-go

COPY . .

# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN make build

WORKDIR $GOPATH/src/github.com/nkmishra1997/php-grpc-go/server

EXPOSE 50051/tcp

# Run the executable
CMD ["./server"]