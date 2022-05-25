package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func StopInstance(client *ec2.EC2, instanceId string) error {
	_, err := client.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{&instanceId},
	})

	return err
}

func StoppingInstance() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	ec2Client := ec2.New(sess)

	instanceId := getInstances()
	err = StopInstance(ec2Client, instanceId)
	if err != nil {
		fmt.Printf("Couldn't stop instance: %v", err)
	}

	fmt.Println("Stopped instance with id: ", instanceId)
}
