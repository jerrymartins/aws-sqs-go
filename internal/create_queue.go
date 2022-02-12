package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Create(s *session.Session, queueName string) *string {
	svc := sqs.New(s)

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: &queueName,
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})

	if err != nil {
		panic("error creating queue")
	}

	return result.QueueUrl
}
