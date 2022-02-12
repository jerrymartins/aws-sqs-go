package internal

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func List(s *session.Session) {
	svc := sqs.New(s)
	result, err := svc.ListQueues(nil)

	if err != nil {
		panic("error listing queues")
	}

	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i, *url)
	}
}
