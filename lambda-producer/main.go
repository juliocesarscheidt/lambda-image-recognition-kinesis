package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juliocesarscheidt/lambda-producer/infra/adapter"
	"log"
	"os"
	"time"
)

// define the clients outside the handler function
var rekognitionClient, _ = adapter.GetRekognitionClient()
var kinesisClient, _ = adapter.GetKinesisClient()

func HandleRequest(ctx context.Context, s3Event events.S3Event) (string, error) {
	// ctx := context.Background()
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
		messagesEncoded, err := adapter.DetectTextsFromImage(ctx, rekognitionClient, bucketName, imagePath)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		// publish to kinesis stream
		failedMessages, err := adapter.PublishToDataStream(ctx, kinesisClient, messagesEncoded, streamName, imagePath)
		if err != nil {
			log.Fatal(err)
			return "", err
		}

		fmt.Println(fmt.Sprintf("Published messages with the total of %d failed messages", failedMessages))
	}

	return "OK", nil
}

func main() {
	lambda.Start(HandleRequest)
}

// func main() {
// 	ctx := context.Background()

// 	s3Event := events.S3Event{
// 		Records: []events.S3EventRecord{
// 			{
// 				EventName:    "ObjectCreated:Post",
// 				EventSource:  "aws:s3",
// 				EventVersion: "2.1",
// 				S3: events.S3Entity{
// 					Bucket: events.S3Bucket{
// 						Arn:  "",
// 						Name: "rekognition-bucket-us-east-1-h4s7kfai",
// 					},
// 					Object: events.S3Object{
// 						ETag:      "",
// 						Key:       "test001.png",
// 						Sequencer: "",
// 						Size:      0,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	HandleRequest(ctx, s3Event)
// }
