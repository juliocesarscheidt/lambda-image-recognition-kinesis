package entrypoint

import (
	"context"
	"fmt"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/juliocesarscheidt/lambda-consumer/infra/adapter"
	"github.com/juliocesarscheidt/lambda-consumer/application/dto"
	"github.com/juliocesarscheidt/lambda-consumer/application/service"
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
			fmt.Println(fmt.Sprintf("Error: %s", err))
			return "", err
		}

		err := service.PersistItem(ctx, dynamoDbClient, tableName, messageDto)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err))
			return "", err
		}
	}

	return "OK", nil
}
