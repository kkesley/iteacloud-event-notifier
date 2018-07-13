package notifier

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//LogEventS3 use reusable s3 session
func LogEventS3(sess *session.Session, bucket string, key string, eventType EventType) error {
	if len(key) <= 0 {
		return errors.New("key must not be empty")
	}
	svc := s3.New(sess)
	strEmail, err := json.Marshal(eventType)
	if err != nil {
		return err
	}
	fmt.Println("sending to s3")
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(strings.Replace(eventType.Event, ":", "_", -1) + "/" + key + ".json"),
		Body:   bytes.NewReader(strEmail),
	})
	return err
}

//LogEventS3Default sends email without s3 session
func LogEventS3Default(region string, bucket string, key string, eventType EventType) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return LogEventS3(sess, bucket, time.Now().Format("20060102150405.999999999Z07:00")+"_"+eventType.ClientID+"_"+key, eventType)
}
