package main

import (
	"aws-sqs-go/internal"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})
	if err != nil {
		panic("aws session")
	}
	internal.Create(sess, "temp-queue")
	internal.List(sess)
	queueUrl, _ := internal.GetUrl(sess, "temp-queue")
	internal.Delete(sess, queueUrl)
	internal.List(sess)
}
