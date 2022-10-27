package service

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/juliocesarscheidt/lambda-consumer/application/adapter"
	"github.com/juliocesarscheidt/lambda-consumer/application/dto"
	"github.com/juliocesarscheidt/lambda-consumer/application/usecase"
	"testing"
	"errors"
)

func TestPersistMessage(t *testing.T) {
	ctx := context.Background()

	dynamoDbClient := &adapter.DynamoDbClientAdapter{
		PutItemWithContext: func(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
			return &dynamodb.PutItemOutput{}, nil
		},
	}

	messageDtoMock := dto.MessageDto{
		Path: "test001.png",
		MessageTexts: []dto.MessageTextsDto{},
	}
	tableName := "rekognition-table"

	messageEncodedMock, _ := json.Marshal(&messageDtoMock)

	err := usecase.PersistMessage(ctx, dynamoDbClient, tableName, messageEncodedMock)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

func TestPersistMessageFailedWhenPersisting(t *testing.T) {
	ctx := context.Background()

	expectedErr := errors.New("Error while persisting data")

	dynamoDbClient := &adapter.DynamoDbClientAdapter{
		PutItemWithContext: func(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
			return nil, expectedErr
		},
	}

	messageDtoMock := dto.MessageDto{
		Path: "test001.png",
		MessageTexts: []dto.MessageTextsDto{},
	}
	tableName := "rekognition-table"
	messageEncodedMock, _ := json.Marshal(&messageDtoMock)

	err := usecase.PersistMessage(ctx, dynamoDbClient, tableName, messageEncodedMock)
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected err to be %v, got %v", expectedErr, err)
	}
}

func TestPersistMessageFailedWhenUnmarshaling(t *testing.T) {
	ctx := context.Background()

	expectedErr := errors.New("unexpected end of JSON input")

	dynamoDbClient := &adapter.DynamoDbClientAdapter{}

	tableName := "rekognition-table"
	messageEncodedMock := []byte{}

	err := usecase.PersistMessage(ctx, dynamoDbClient, tableName, messageEncodedMock)
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected err to be %v, got %v", expectedErr, err)
	}
}
