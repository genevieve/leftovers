package elbv2

import (
	"fmt"

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

func (t TargetGroups) Delete() error {
	targetGroups, err := t.client.DescribeTargetGroups(&awselbv2.DescribeTargetGroupsInput{})
	if err != nil {
		return fmt.Errorf("Describing target groups: %s", err)
	}

	for _, g := range targetGroups.TargetGroups {
		n := *g.TargetGroupName

		proceed := t.logger.Prompt(fmt.Sprintf("Are you sure you want to delete target group %s?", n))
		if !proceed {
			continue
		}

		_, err := t.client.DeleteTargetGroup(&awselbv2.DeleteTargetGroupInput{TargetGroupArn: g.TargetGroupArn})
		if err == nil {
			t.logger.Printf("SUCCESS deleting target group %s\n", n)
		} else {
			t.logger.Printf("ERROR deleting target group %s: %s\n", n, err)
		}
	}

	return nil
}
