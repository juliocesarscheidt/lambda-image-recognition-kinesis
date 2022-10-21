package adapter

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/juliocesarscheidt/lambda-producer/infra/adapter"
	"testing"
)

const confidenceThreshold = 90

func TestBuildMessagesFromTextsOnlyAboveThreshold(t *testing.T) {
	textDetectionsMock := []*rekognition.TextDetection{
		{
			DetectedText: aws.String("Text 1"),
			Type:         aws.String("LINE"),
			Confidence:   aws.Float64(confidenceThreshold + 1),
		},
		{
			DetectedText: aws.String("Text 2"),
			Type:         aws.String("LINE"),
			Confidence:   aws.Float64(confidenceThreshold),
		},
		{
			DetectedText: aws.String("Text 3"),
			Type:         aws.String("LINE"),
			Confidence:   aws.Float64(confidenceThreshold - 1),
		},
	}

	imagePath := "test001.png"
	messageDto := adapter.BuildMessagesFromTexts(textDetectionsMock, imagePath)

	// check if the detected texts returned will be only the ones with confidence above or equal threshold
	if len(messageDto.MessageTexts) != 2 {
		t.Errorf("Expected 2 message, got %d", len(messageDto.MessageTexts))
	}
}

func TestBuildMessagesFromTextsOnlyLineType(t *testing.T) {
	textDetectionsMock := []*rekognition.TextDetection{
		{
			DetectedText: aws.String("Text 1"),
			Type:         aws.String("TEXT"),
			Confidence:   aws.Float64(confidenceThreshold),
		},
		{
			DetectedText: aws.String("Text 2"),
			Type:         aws.String("TEXT"),
			Confidence:   aws.Float64(confidenceThreshold),
		},
	}

	imagePath := "test001.png"
	messageDto := adapter.BuildMessagesFromTexts(textDetectionsMock, imagePath)

	// check if the detected texts returned will be only the ones with confidence above the threshold
	if len(messageDto.MessageTexts) != 0 {
		t.Errorf("Expected 0 message, got %d", len(messageDto.MessageTexts))
	}
}

func TestDetectTexts(t *testing.T) {
	bucketName := "rekognition-bucket"
	imagePath := "test001.png"

	rekognitionMock := &rekognition.DetectTextOutput{
		TextDetections: []*rekognition.TextDetection{
			{
				DetectedText: aws.String("Text 1"),
				Type:         aws.String("LINE"),
				Confidence:   aws.Float64(confidenceThreshold),
			},
			{
				DetectedText: aws.String("Text 2"),
				Type:         aws.String("LINE"),
				Confidence:   aws.Float64(confidenceThreshold),
			},
		},
	}

	ctx := context.Background()

	rekognitionClientMock := &adapter.RekognitionClientAdapter{
		DetectTextWithContext: func(ctx context.Context, input *rekognition.DetectTextInput, opts ...request.Option) (*rekognition.DetectTextOutput, error) {
			return rekognitionMock, nil
		},
	}

	detectTextInput := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(bucketName),
				Name:   aws.String(imagePath),
			},
		},
	}
	textDetections, err := adapter.DetectTexts(ctx, rekognitionClientMock, detectTextInput)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	rekognitionMockTextDetections := rekognitionMock.TextDetections

	if len(rekognitionMockTextDetections) != len(textDetections) {
		t.Errorf("Expected %v, got %v", rekognitionMock.TextDetections, textDetections)
	}
}
