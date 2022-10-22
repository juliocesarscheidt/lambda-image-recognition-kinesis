package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juliocesarscheidt/lambda-producer/infra/entrypoint"
)

func main() {
	lambda.Start(entrypoint.HandleRequest)
}
