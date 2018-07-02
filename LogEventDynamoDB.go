package notifier

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func LogEventDynamoDB(tableName string, region string, event EventType) error {
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
