package main

// import (
// 	"context"
// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/juliocesarscheidt/lambda-producer/infra/entrypoint"
// )

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
// 						Name: "rekognition-bucket-1yua2f4l",
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

// 	entrypoint.HandleRequest(ctx, s3Event)
// }
