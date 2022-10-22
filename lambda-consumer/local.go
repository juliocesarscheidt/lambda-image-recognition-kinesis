package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/juliocesarscheidt/lambda-consumer/application/dto"
// 	"github.com/juliocesarscheidt/lambda-consumer/infra/entrypoint"
// )

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
// 		fmt.Println(fmt.Sprintf("Error: %s", err))
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

// 	entrypoint.HandleRequest(ctx, kinesisEvent)
// }
