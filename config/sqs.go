package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var (
	SqsClient *sqs.Client
	QueueURL  string
)

func InitSQS(queueUrl string) {
	cfg, err := LoadAWSConfig()
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	SqsClient = sqs.NewFromConfig(cfg)
	QueueURL = queueUrl
}

func SendMessageToSQS(ctx context.Context, messageBody string) error {
	input := &sqs.SendMessageInput{
		MessageBody: &messageBody,
		QueueUrl:    &QueueURL,
	}

	_, err := SqsClient.SendMessage(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func ReceiveMessageFromSQS(ctx context.Context) (*sqs.ReceiveMessageOutput, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            &QueueURL,
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     20,
	}

	return SqsClient.ReceiveMessage(ctx, input)
}

func DeleteMessageFromSQS(ctx context.Context, receiptHandle *string) error {
	input := &sqs.DeleteMessageInput{
		QueueUrl:      &QueueURL,
		ReceiptHandle: receiptHandle,
	}

	_, err := SqsClient.DeleteMessage(ctx, input)
	return err
}
