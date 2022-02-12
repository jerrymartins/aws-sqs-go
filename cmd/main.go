package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})

	if err != nil {
		panic("erro")
	}

	svc := sqs.New(sess)
	result, err := svc.ListQueues(nil)

	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i, *url)
	}
}
