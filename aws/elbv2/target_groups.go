package elbv2

import (
	"fmt"
	"strings"

	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
)

type targetGroupsClient interface {
	DescribeTargetGroups(*awselbv2.DescribeTargetGroupsInput) (*awselbv2.DescribeTargetGroupsOutput, error)
	DeleteTargetGroup(*awselbv2.DeleteTargetGroupInput) (*awselbv2.DeleteTargetGroupOutput, error)
}

type TargetGroups struct {
	client targetGroupsClient
	logger logger
}

func NewTargetGroups(client targetGroupsClient, logger logger) TargetGroups {
	return TargetGroups{
		client: client,
		logger: logger,
	}
}

func (t TargetGroups) List(filter string) (map[string]string, error) {
	targetGroups, err := t.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, g := range targetGroups {
		delete[g.identifier] = *g.arn
	}

	return delete, nil
}

func (t TargetGroups) list(filter string) ([]TargetGroup, error) {
	targetGroups, err := t.client.DescribeTargetGroups(&awselbv2.DescribeTargetGroupsInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing target groups: %s", err)
	}

	var resources []TargetGroup
	for _, g := range targetGroups.TargetGroups {
		resource := NewTargetGroup(t.client, g.TargetGroupName, g.TargetGroupArn)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := t.logger.Prompt(fmt.Sprintf("Are you sure you want to delete target group %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (t TargetGroups) Delete(targetGroups map[string]string) error {
	var resources []TargetGroup
	for name, arn := range targetGroups {
		resources = append(resources, NewTargetGroup(t.client, &name, &arn))
	}

	return t.cleanup(resources)
}

func (t TargetGroups) cleanup(resources []TargetGroup) error {
	for _, resource := range resources {
		err := resource.Delete()

		if err == nil {
			t.logger.Printf("SUCCESS deleting target group %s\n", resource.identifier)
		} else {
			t.logger.Printf("ERROR deleting target group %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}
