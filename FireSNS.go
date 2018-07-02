package notifier

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

//FireSNS push notification to a topic
func FireSNS(topic string, message string) error {
	return FireSNSWithAttribute(topic, message, make(map[string]string))
}

//FireSNSWithAttribute push notification to a topic with an attribute
func FireSNSWithAttribute(topic string, message string, attributes map[string]string) error {
	svc := sns.New(session.New())
	// params will be sent to the publish call included here is the bare minimum params to send a message.
	snsAttributes := make(map[string]*sns.MessageAttributeValue)
	for key, val := range attributes {
		snsAttributes[key] = &sns.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(val),
		}
	}
	params := &sns.PublishInput{
		Message:           aws.String(message), // This is the message itself (can be XML / JSON / Text - anything you want)
		TopicArn:          aws.String(topic),   //Get this from the Topic in the AWS console.
		MessageAttributes: snsAttributes,
	}

	resp, err := svc.Publish(params) //Call to puclish the message

	if err != nil { //Check for errors
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	return nil
}
