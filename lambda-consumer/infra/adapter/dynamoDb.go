package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/request"
	"fmt"
	"os"
)

// client adapter
type DynamoDbClientAdapter struct {
	PutItemWithContext func(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error)
}

func GetDynamoDbClient() (*DynamoDbClientAdapter, error) {
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
	dynamoDbClientAdapter := &DynamoDbClientAdapter{
		PutItemWithContext: client.PutItemWithContext,
	}
	return dynamoDbClientAdapter, nil
}
