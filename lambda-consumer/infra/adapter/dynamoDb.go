package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

// client adapter
type DynamoDbClientAdapter struct {
	PutItem func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
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
		log.Fatal(err)
		return nil, err
	}
	client := dynamodb.New(sess)
	dynamoDbClientAdapter := &DynamoDbClientAdapter{
		PutItem: client.PutItem,
	}
	return dynamoDbClientAdapter, nil
}

func PutItem(dynamoDbClient *DynamoDbClientAdapter, tableName string, item interface{}) error {
	dynamoDbItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatal(err)
		return err
	}
	inputPutItem := &dynamodb.PutItemInput{
		Item:      dynamoDbItem,
		TableName: aws.String(tableName),
	}
	_, err = dynamoDbClient.PutItem(inputPutItem)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
