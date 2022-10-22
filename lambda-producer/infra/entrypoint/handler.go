package entrypoint

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/juliocesarscheidt/lambda-producer/application/service"
	"github.com/juliocesarscheidt/lambda-producer/infra/adapter"
	"log"
	"os"
	"time"
)

// define the clients outside the handler function
var rekognitionClient, _ = adapter.GetRekognitionClient()
var kinesisClient, _ = adapter.GetKinesisClient()

func HandleRequest(ctx context.Context, s3Event events.S3Event) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// get stream name from env
	streamName := os.Getenv("KINESIS_STREAM_NAME")

	// get information from the events
	for _, record := range s3Event.Records {
		s3 := record.S3
		bucketName := s3.Bucket.Name
		imagePath := s3.Object.Key
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, bucketName, imagePath)

		// retrieve the texts from the image
		messagesEncoded, err := service.DetectTextsFromImage(ctx, rekognitionClient, bucketName, imagePath)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err))
			return "", err
		}
		// publish to kinesis stream
		failedMessages, err := service.PublishToDataStream(ctx, kinesisClient, messagesEncoded, streamName, imagePath)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err))
			return "", err
		}

		fmt.Println(fmt.Sprintf("Published messages with the total of %d failed messages", failedMessages))
	}

	return "OK", nil
}
