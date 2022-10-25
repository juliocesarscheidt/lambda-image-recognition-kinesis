package adapter

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/juliocesarscheidt/lambda-producer/application/adapter"
	"os"
)

func GetKinesisClient() (*adapter.KinesisClientAdapter, error) {
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
	kinesisClientAdapter := &adapter.KinesisClientAdapter{
		PutRecordsWithContext: client.PutRecordsWithContext,
	}
	return kinesisClientAdapter, nil
}
