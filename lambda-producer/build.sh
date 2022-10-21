#!/bin/bash

docker container run -v $PWD:/go/src/app -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:1.18-alpine go build -ldflags="-s -w" -o producer

# GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w" -o producer
# ./producer

# go vet
# go fmt

# go mod download
# go mod tidy

# go run main.go
