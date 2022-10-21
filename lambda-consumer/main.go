package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juliocesarscheidt/lambda-consumer/application/dto"
	"github.com/juliocesarscheidt/lambda-consumer/infra/adapter"
	"log"
	"os"
	"time"
)

// define the clients outside the handler function
var dynamoDbClient, _ = adapter.GetDynamoDbClient()

func HandleRequest(ctx context.Context, kinesisEvent events.KinesisEvent) (string, error) {
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// get table name from env
	tableName := os.Getenv("TABLE_NAME")

	// get information from the events
	for _, record := range kinesisEvent.Records {
		kinesis := record.Kinesis

		fmt.Println(fmt.Sprintf("Partition Key: %s", kinesis.PartitionKey))
		fmt.Println(fmt.Sprintf("Sequence Number: %s", kinesis.SequenceNumber))

		fmt.Println(string(kinesis.Data))

		var messageDto dto.MessageDto
		if err := json.Unmarshal(kinesis.Data, &messageDto); err != nil {
			log.Fatal(err)
			return "", err
		}

		err := adapter.PutItem(dynamoDbClient, tableName, messageDto)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
	}

	return "OK", nil
}

func main() {
	lambda.Start(HandleRequest)
}

// func main() {
// 	ctx := context.Background()

// 	messageTexts := []dto.MessageTextsDto{
// 		{
// 			TextType:     "LINE",
// 			Confidence:   95.0,
// 			DetectedText: "Blackdevs",
// 		},
// 		{
// 			TextType:     "LINE",
// 			Confidence:   95.0,
// 			DetectedText: "Software",
// 		},
// 	}
// 	message := &dto.MessageDto{
// 		Path:         "test001.png",
// 		MessageTexts: messageTexts,
// 	}

// 	messageEncoded, err := json.Marshal(&message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(messageEncoded)

// 	kinesisEvent := events.KinesisEvent{
// 		Records: []events.KinesisEventRecord{
// 			{
// 				Kinesis: events.KinesisRecord{
// 					PartitionKey:   "test001.png",
// 					SequenceNumber: "49545115243490985018280067714973144582180062593244200961",
// 					Data:           []byte(messageEncoded),
// 				},
// 				EventSource:  "aws:kinesis",
// 				EventVersion: "1.0",
// 				EventID:      "shardId-000000000000:49545115243490985018280067714973144582180062593244200961",
// 				EventName:    "aws:kinesis:record",
// 				AwsRegion:    "us-east-1",
// 			},
// 		},
// 	}

// 	HandleRequest(ctx, kinesisEvent)
// }
