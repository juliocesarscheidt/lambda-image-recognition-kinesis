package usecase

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/juliocesarscheidt/lambda-producer/application/adapter"
	"github.com/juliocesarscheidt/lambda-producer/application/usecase"
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
	messageDto := usecase.BuildMessagesFromTexts(textDetectionsMock, imagePath)

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
	messageDto := usecase.BuildMessagesFromTexts(textDetectionsMock, imagePath)

	// check if the detected texts returned will be only the ones with confidence above the threshold
	if len(messageDto.MessageTexts) != 0 {
		t.Errorf("Expected 0 message, got %d", len(messageDto.MessageTexts))
	}
}

func TestDetectTextsFromImage(t *testing.T) {
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
	messageDtoMock := usecase.BuildMessagesFromTexts(rekognitionMock.TextDetections, imagePath)
	messageEncodedMock, err := json.Marshal(&messageDtoMock)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	ctx := context.Background()
	rekognitionClientMock := &adapter.RekognitionClientAdapter{
		DetectTextWithContext: func(ctx context.Context, input *rekognition.DetectTextInput, opts ...request.Option) (*rekognition.DetectTextOutput, error) {
			return rekognitionMock, nil
		},
	}

	messagesEncoded, err := usecase.DetectTextsFromImage(ctx, rekognitionClientMock, bucketName, imagePath)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if len(messagesEncoded) != len(messageEncodedMock) {
		t.Errorf("Expected %v, got %v", messagesEncoded, messageEncodedMock)
	}
}
