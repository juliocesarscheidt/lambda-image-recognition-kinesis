package adapter

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/juliocesarscheidt/lambda-producer/application/dto"
	"log"
	"os"
)

const confidenceThreshold = 90

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

func DetectTexts(ctx context.Context, rekognitionClient *RekognitionClientAdapter,
	detectTextInput *rekognition.DetectTextInput) ([]*rekognition.TextDetection, error) {
	detectTextOutput, err := rekognitionClient.DetectTextWithContext(ctx, detectTextInput)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return detectTextOutput.TextDetections, nil
}

func DetectTextsFromImage(ctx context.Context, rekognitionClient *RekognitionClientAdapter,
	bucketName string, imagePath string) ([]byte, error) {
	detectTextInput := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(bucketName),
				Name:   aws.String(imagePath),
			},
		},
	}
	textDetections, err := DetectTexts(ctx, rekognitionClient, detectTextInput)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// build messages dto from detected texts
	messageDto := BuildMessagesFromTexts(textDetections, imagePath)
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
