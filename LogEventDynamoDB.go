package notifier

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kkesley/random"
)

//LogEventDynamoDB log an event to dynamodb table
func LogEventDynamoDB(tableName string, region string, event EventType) error {
	event.Principal = event.Principal + "[" + time.Now().Format("20060102150405.999999999Z07:00") + ":" + random.MakeID(20) + "]"
	ev, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		fmt.Println(err)
		return err
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return err
	}
	svc := dynamodb.New(sess)
	input := &dynamodb.PutItemInput{
		Item:      ev,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}
