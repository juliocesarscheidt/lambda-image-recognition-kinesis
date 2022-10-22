package service

import (
	"fmt"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/juliocesarscheidt/lambda-consumer/infra/adapter"
)

func PersistItem(ctx context.Context, dynamoDbClient *adapter.DynamoDbClientAdapter, tableName string, item interface{}) error {
	dynamoDbItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return err
	}
	inputPutItem := &dynamodb.PutItemInput{
		Item:      dynamoDbItem,
		TableName: aws.String(tableName),
	}
	_, err = dynamoDbClient.PutItemWithContext(ctx, inputPutItem)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return err
	}
	return nil
}
