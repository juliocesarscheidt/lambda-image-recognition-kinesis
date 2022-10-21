package adapter

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"log"
	"os"
)

// client adapter
type KinesisClientAdapter struct {
	PutRecordsWithContext func(ctx aws.Context, input *kinesis.PutRecordsInput, opts ...request.Option) (*kinesis.PutRecordsOutput, error)
}

func GetKinesisClient() (*KinesisClientAdapter, error) {
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
	client := kinesis.New(sess)
	kinesisClientAdapter := &KinesisClientAdapter{
		PutRecordsWithContext: client.PutRecordsWithContext,
	}
	return kinesisClientAdapter, nil
}

func PublishToDataStream(ctx context.Context, kinesisClient *KinesisClientAdapter,
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
