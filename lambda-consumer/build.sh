#!/bin/bash

docker container run -v $PWD:/go/src/app -w /go/src/app -e GOOS=linux -e GOARCH=amd64 -e GO111MODULE=on -e CGO_ENABLED=0 golang:1.18-alpine go build -ldflags="-s -w" -o consumer

# GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w" -o consumer
# ./consumer

# go vet
# go fmt

# go test tests/**/*_test.go -v
go test -cover -coverpkg=github.com/juliocesarscheidt/lambda-consumer/application/usecase -coverprofile cover.out tests/**/*_test.go
# go tool cover -html=cover.out -o coverage.html

docker container run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:lts
docker container logs -f --tail 100 sonarqube

export SONARQUBE_URL='127.0.0.1:9000'
export TOKEN='TOKEN'
export PROJECT_KEY='lambda-consumer'

docker container run \
  --rm --name sonarscanner \
  --network host \
  -e SONAR_HOST_URL="http://${SONARQUBE_URL}" \
  -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=${PROJECT_KEY}" \
  -e SONAR_LOGIN="${TOKEN}" \
  -v "${PWD}:/usr/src" \
  -w /usr/src \
  sonarsource/sonar-scanner-cli:4

# go install github.com/jstemmer/go-junit-report/v2@latest
# go test ./... -v | go-junit-report > report.xml

# go mod download
# go mod tidy

# go run main.go
