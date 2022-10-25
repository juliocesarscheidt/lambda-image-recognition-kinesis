package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/juliocesarscheidt/lambda-consumer/application/adapter"
	"github.com/juliocesarscheidt/lambda-consumer/application/dto"
)

func PersistMessage(ctx context.Context, dynamoDbClient *adapter.DynamoDbClientAdapter, tableName string, itemData []byte) error {
	var messageDto dto.MessageDto
	if err := json.Unmarshal(itemData, &messageDto); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return err
	}
	dynamoDbItem, err := dynamodbattribute.MarshalMap(messageDto)
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
