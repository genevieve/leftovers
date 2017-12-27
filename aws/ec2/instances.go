package ec2

import (
	"fmt"
	"strings"

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
			if a.alreadyShutdown(*i.State.Name) {
				continue
			}

			n := a.clearerName(*i.InstanceId, i.Tags)

			proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to terminate instance %s?", n))
			if !proceed {
				continue
			}

			_, err := a.client.TerminateInstances(&awsec2.TerminateInstancesInput{
				InstanceIds: []*string{i.InstanceId},
			})
			if err == nil {
				a.logger.Printf("SUCCESS terminating instance %s\n", n)
			} else {
				a.logger.Printf("ERROR terminating instance %s: %s\n", n, err)
			}
		}
	}

	return nil
}

func (a Instances) alreadyShutdown(state string) bool {
	return state == "shutting-down" || state == "terminated"
}

func (a Instances) clearerName(id string, tags []*awsec2.Tag) string {
	extra := []string{}
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		return fmt.Sprintf("%s (%s)", id, strings.Join(extra, ", "))
	}

	return id
}
