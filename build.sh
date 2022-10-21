#!/bin/bash

# docker container run -e GOOS=linux -e GOARCH=amd64 -v $PWD:/app -w /app golang:1.18-alpine go build -ldflags="-s -w" -o main

# GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main

# zip main.zip main
# mv main.zip ../infrastrucuture/terraform/

# ./main

# go vet
# go fmt

# go mod download
# go mod tidy

# go run main.go
