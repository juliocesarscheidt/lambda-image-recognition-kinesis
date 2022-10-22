package usecase

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/juliocesarscheidt/lambda-producer/infra/adapter"
)

func PublishToDataStream(ctx context.Context, kinesisClient *adapter.KinesisClientAdapter,
	messageEncoded []byte, streamName string, partitionKey string) (int64, error) {
	result, err := kinesisClient.PutRecordsWithContext(ctx, &kinesis.PutRecordsInput{
		Records: []*kinesis.PutRecordsRequestEntry{
			{
				Data:         messageEncoded,
				PartitionKey: aws.String(partitionKey),
			},
		},
		StreamName: aws.String(streamName),
	})
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return *result.FailedRecordCount, nil
}
