package adapter

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/juliocesarscheidt/lambda-consumer/application/adapter"
	"os"
)

func GetDynamoDbClient() (*adapter.DynamoDbClientAdapter, error) {
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		region = "us-east-1"
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return nil, err
	}
	client := dynamodb.New(sess)
	dynamoDbClientAdapter := &adapter.DynamoDbClientAdapter{
		PutItemWithContext: client.PutItemWithContext,
	}
	return dynamoDbClientAdapter, nil
}
