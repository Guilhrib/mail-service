package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	sdkConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigFiles([]string{".aws/config"}),
		config.WithSharedCredentialsFiles([]string{".aws/credentials"}),
	)

	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	sqsClient := sqs.NewFromConfig(sdkConfig)
	fmt.Println("Let's list the queues for your account.")

	queueName := "mail-queue-test"
	queueUrl, err := sqsClient.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		fmt.Println("Couldn't get queue url. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	messageOutput, err := sqsClient.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl.QueueUrl,
		MaxNumberOfMessages: 1,
	})

	if err != nil {
		fmt.Println("Couldn't receive message. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	messageBytes := messageOutput.Messages[0].Body
	message := string(bytes.Join([][]byte{[]byte(*messageBytes)}, []byte{}))
	fmt.Println(message)
}
