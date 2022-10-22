package service

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/juliocesarscheidt/lambda-consumer/infra/adapter"
	"github.com/juliocesarscheidt/lambda-consumer/application/dto"
	"github.com/juliocesarscheidt/lambda-consumer/application/service"
	"testing"
	"errors"
)

func TestPersistItem(t *testing.T) {
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

	err := service.PersistItem(ctx, dynamoDbClient, tableName, messageDtoMock)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

func TestPersistItemFailed(t *testing.T) {
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

	err := service.PersistItem(ctx, dynamoDbClient, tableName, messageDtoMock)
	if err != expectedErr {
		t.Errorf("Expected err to be %v, got %v", expectedErr, err)
	}
}
