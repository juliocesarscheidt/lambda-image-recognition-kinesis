package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// client adapter
type RekognitionClientAdapter struct {
	DetectTextWithContext func(ctx aws.Context, input *rekognition.DetectTextInput, opts ...request.Option) (*rekognition.DetectTextOutput, error)
}
