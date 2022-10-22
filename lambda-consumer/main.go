package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juliocesarscheidt/lambda-consumer/infra/entrypoint"
)

func main() {
	lambda.Start(entrypoint.HandleRequest)
}
