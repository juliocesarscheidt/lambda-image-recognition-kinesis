package usecase

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/juliocesarscheidt/lambda-producer/application/adapter"
	"github.com/juliocesarscheidt/lambda-producer/application/dto"
)

const confidenceThreshold = 90

func BuildMessagesFromTexts(textDetections []*rekognition.TextDetection, imagePath string) dto.MessageDto {
	messageDto := dto.MessageDto{
		Path:         imagePath,
		MessageTexts: []dto.MessageTextsDto{},
	}
	for _, textDetection := range textDetections {
		text := *textDetection.DetectedText
		fmt.Println(fmt.Sprintf("Text: %s", text))

		textType := *textDetection.Type
		if textType != "LINE" {
			fmt.Println("Type is not LINE, skipping...")
			continue
		}
		confidence := *textDetection.Confidence
		fmt.Println(fmt.Sprintf("Confidence: %f", confidence))

		if confidence < confidenceThreshold {
			fmt.Println(fmt.Sprintf("Confidence is BELOW threshold (%d), skipping...", confidenceThreshold))
			continue
		}
		fmt.Println(fmt.Sprintf("Confidence is ABOVE or EQUAL threshold (%d), adding text to be published", confidenceThreshold))

		messageText := dto.MessageTextsDto{
			TextType:     textType,
			Confidence:   confidence,
			DetectedText: text,
		}
		messageDto.MessageTexts = append(messageDto.MessageTexts, messageText)
	}
	return messageDto
}

func DetectTextsFromImage(ctx context.Context, rekognitionClient *adapter.RekognitionClientAdapter,
	bucketName string, imagePath string) ([]byte, error) {
	detectTextInput := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(bucketName),
				Name:   aws.String(imagePath),
			},
		},
	}
	detectTextOutput, err := rekognitionClient.DetectTextWithContext(ctx, detectTextInput)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return nil, err
	}
	textDetections := detectTextOutput.TextDetections
	// build messages dto from detected texts
	messageDto := BuildMessagesFromTexts(textDetections, imagePath)
	messageEncoded, err := json.Marshal(&messageDto)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return nil, err
	}
	// messageEncoded is a array of bytes []byte{}
	fmt.Println(string(messageEncoded))

	messageSizeInBytes := binary.Size(messageEncoded)
	fmt.Println(fmt.Sprintf("Message size in KBs: %.2f", float64(messageSizeInBytes)/1024))

	return messageEncoded, nil
}
