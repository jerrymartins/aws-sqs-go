package internal

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetUrl(s *session.Session, queueName string) (string, error) {
	svc := sqs.New(s)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		return "", err
	}

	return *result.QueueUrl, nil
}
