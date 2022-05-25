package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetRunningInstances(client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {

	result, err := client.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Scheduler"),
				Values: []*string{aws.String("True")},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending"), aws.String("stopped")},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, err

}

func getInstances() string {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)

	}

	ec2Client := ec2.New(sess)
	var id string
	runningInstances, err := GetRunningInstances(ec2Client)
	if err != nil {
		fmt.Printf("Couldn't retrieve running instances: %v", err)

	}

	for _, reservation := range runningInstances.Reservations {
		for _, instance := range reservation.Instances {
			id = *instance.InstanceId
			return id
		}
		return id
	}
	return id
}
