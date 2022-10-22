package adapter

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
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
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return nil, err
	}
	client := kinesis.New(sess)
	kinesisClientAdapter := &KinesisClientAdapter{
		PutRecordsWithContext: client.PutRecordsWithContext,
	}
	return kinesisClientAdapter, nil
}
