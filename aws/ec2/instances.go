package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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

func (a Instances) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	instances, err := a.client.DescribeInstances(&awsec2.DescribeInstancesInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing instances: %s", err)
	}

	for _, r := range instances.Reservations {
		for _, i := range r.Instances {
			if a.alreadyShutdown(*i.State.Name) {
				continue
			}

			n := a.clearerName(*i.InstanceId, i.Tags, *i.KeyName)

			if !strings.Contains(n, filter) {
				continue
			}

			proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to terminate instance %s?", n))
			if !proceed {
				continue
			}

			delete[n] = *i.InstanceId
		}
	}

	return delete, nil
}

func (i Instances) Delete(instances map[string]string) error {
	for name, id := range instances {
		_, err := i.client.TerminateInstances(&awsec2.TerminateInstancesInput{
			InstanceIds: []*string{aws.String(id)},
		})

		if err == nil {
			i.logger.Printf("SUCCESS terminating instance %s\n", name)
		} else {
			i.logger.Printf("ERROR terminating instance %s: %s\n", name, err)
		}
	}

	return nil
}

func (a Instances) alreadyShutdown(state string) bool {
	return state == "shutting-down" || state == "terminated"
}

func (a Instances) clearerName(id string, tags []*awsec2.Tag, keyName string) string {
	extra := []string{}
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if keyName != "" {
		extra = append(extra, fmt.Sprintf("KeyPairName:%s", keyName))
	}

	if len(extra) > 0 {
		return fmt.Sprintf("%s (%s)", id, strings.Join(extra, ", "))
	}

	return id
}
