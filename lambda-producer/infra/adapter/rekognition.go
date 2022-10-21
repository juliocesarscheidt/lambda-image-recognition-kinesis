package adapter

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/juliocesarscheidt/lambda-producer/application/dto"
	"log"
	"os"
)

func GetRekognitionClient() (*rekognition.Rekognition, error) {
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
	return rekognition.New(sess), nil
}

func GetImageTexts(ctx context.Context, rekognitionClient *rekognition.Rekognition,
	bucketName string, imagePath string) ([]byte, error) {

	messageDto := dto.MessageDto{
		Path:         imagePath,
		MessageTexts: []dto.MessageTexts{},
	}
	detectTextOutput, err := rekognitionClient.DetectTextWithContext(ctx, &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(bucketName),
				Name:   aws.String(imagePath),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	textDetections := detectTextOutput.TextDetections
	for _, textDetection := range textDetections {
		textType := *textDetection.Type
		if textType != "LINE" {
			fmt.Println("Type is not LINE, skipping...")
			continue
		}
		confidence := *textDetection.Confidence
		fmt.Println(fmt.Sprintf("Confidence: %f", confidence))
		if confidence < 90 {
			fmt.Println("Confidence is less than 90, skipping...")
			continue
		}
		fmt.Println("Confidence is greater or equal 90, it will be published...")

		messageText := dto.MessageTexts{
			TextType:     textType,
			Confidence:   confidence,
			DetectedText: *textDetection.DetectedText,
		}
		messageDto.MessageTexts = append(messageDto.MessageTexts, messageText)
	}

	messageEncoded, err := json.Marshal(&messageDto)
	if err != nil {
		log.Fatal(err)
	}
	// messageEncoded is a array of bytes []byte{}
	fmt.Println(string(messageEncoded))

	messageSizeInBytes := binary.Size(messageEncoded)
	fmt.Println(fmt.Sprintf("Message size in KBs: %.2f", float64(messageSizeInBytes)/1024))

	return messageEncoded, nil
}
