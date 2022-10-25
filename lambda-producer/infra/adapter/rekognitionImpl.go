package adapter

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/juliocesarscheidt/lambda-producer/application/adapter"
	"os"
)

func GetRekognitionClient() (*adapter.RekognitionClientAdapter, error) {
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
	client := rekognition.New(sess)
	rekognitionClientAdapter := &adapter.RekognitionClientAdapter{
		DetectTextWithContext: client.DetectTextWithContext,
	}
	return rekognitionClientAdapter, nil
}
