package internal

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Delete(s *session.Session, queueUrl string) {
	svc := sqs.New(s)
	_, err := svc.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: &queueUrl,
	})

	if err != nil {
		panic("error deleting queue")
	}
}
