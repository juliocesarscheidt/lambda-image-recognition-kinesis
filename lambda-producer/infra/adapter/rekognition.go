package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"os"
)

// client adapter
type RekognitionClientAdapter struct {
	DetectTextWithContext func(ctx aws.Context, input *rekognition.DetectTextInput, opts ...request.Option) (*rekognition.DetectTextOutput, error)
}

func GetRekognitionClient() (*RekognitionClientAdapter, error) {
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		region = "us-east-1"
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}
	client := rekognition.New(sess)
	rekognitionClientAdapter := &RekognitionClientAdapter{
		DetectTextWithContext: client.DetectTextWithContext,
	}
	return rekognitionClientAdapter, nil
}
