package usecase

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/juliocesarscheidt/lambda-producer/infra/adapter"
	"github.com/juliocesarscheidt/lambda-producer/application/usecase"
	"testing"
)

func TestPublishToDataStream(t *testing.T) {
	ctx := context.Background()

	kinesisClientMock := &adapter.KinesisClientAdapter{
		PutRecordsWithContext: func(ctx aws.Context, input *kinesis.PutRecordsInput, opts ...request.Option) (*kinesis.PutRecordsOutput, error) {
			return &kinesis.PutRecordsOutput{
				FailedRecordCount: aws.Int64(0),
			}, nil
		},
	}

	messagesEncoded := []byte{}
	streamName := "rekognition-stream"
	imagePath := "test001.png"

	failedMessages, err := usecase.PublishToDataStream(ctx, kinesisClientMock, messagesEncoded, streamName, imagePath)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if failedMessages != 0 {
		t.Errorf("Expected 0 failed messages, got %d", failedMessages)
	}
}
