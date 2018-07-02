package notifier

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

//FireSNS push notification to a topic
func FireSNS(topic string, message string) error {
	svc := sns.New(session.New())
	// params will be sent to the publish call included here is the bare minimum params to send a message.
	params := &sns.PublishInput{
		Message:  aws.String(message), // This is the message itself (can be XML / JSON / Text - anything you want)
		TopicArn: aws.String(topic),   //Get this from the Topic in the AWS console.
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
