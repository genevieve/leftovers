package ec2

import (
	"fmt"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type instancesClient interface {
	DescribeInstances(*awsec2.DescribeInstancesInput) (*awsec2.DescribeInstancesOutput, error)
	TerminateInstances(*awsec2.TerminateInstancesInput) (*awsec2.TerminateInstancesOutput, error)
}

type Instances struct {
	client instancesClient
	logger logger
}

func NewInstances(client instancesClient, logger logger) Instances {
	return Instances{
		client: client,
		logger: logger,
	}
}

func (a Instances) Delete() error {
	instances, err := a.client.DescribeInstances(&awsec2.DescribeInstancesInput{})
	if err != nil {
		return fmt.Errorf("Describing instances: %s", err)
	}

	for _, r := range instances.Reservations {
		for _, i := range r.Instances {
			if alreadyShutdown(*i.State.Name) {
				continue
			}

			instanceId := *i.InstanceId
			n := name(i)
			proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to terminate instance %s%s?", instanceId, n))
			if !proceed {
				continue
			}

			_, err := a.client.TerminateInstances(&awsec2.TerminateInstancesInput{InstanceIds: []*string{i.InstanceId}})
			if err == nil {
				a.logger.Printf("SUCCESS terminating instance %s%s\n", instanceId, n)
			} else {
				a.logger.Printf("ERROR terminating instance %s%s: %s\n", instanceId, n, err)
			}
		}
	}

	return nil
}

func alreadyShutdown(state string) bool {
	return state == "shutting-down" || state == "terminated"
}

func name(i *awsec2.Instance) string {
	for _, t := range i.Tags {
		if *t.Key == "Name" {
			return fmt.Sprintf("/%s", *t.Value)
		}
	}
	return ""
}
