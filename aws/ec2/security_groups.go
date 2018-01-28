package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type securityGroupsClient interface {
	DescribeSecurityGroups(*awsec2.DescribeSecurityGroupsInput) (*awsec2.DescribeSecurityGroupsOutput, error)
	RevokeSecurityGroupIngress(*awsec2.RevokeSecurityGroupIngressInput) (*awsec2.RevokeSecurityGroupIngressOutput, error)
	RevokeSecurityGroupEgress(*awsec2.RevokeSecurityGroupEgressInput) (*awsec2.RevokeSecurityGroupEgressOutput, error)
	DeleteSecurityGroup(*awsec2.DeleteSecurityGroupInput) (*awsec2.DeleteSecurityGroupOutput, error)
}

type SecurityGroups struct {
	client securityGroupsClient
	logger logger
}

func NewSecurityGroups(client securityGroupsClient, logger logger) SecurityGroups {
	return SecurityGroups{
		client: client,
		logger: logger,
	}
}

func (e SecurityGroups) List(filter string) (map[string]string, error) {
	securityGroups, err := e.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, s := range securityGroups {
		delete[s.identifier] = *s.id
	}

	return delete, nil
}

func (e SecurityGroups) list(filter string) ([]SecurityGroup, error) {
	output, err := e.client.DescribeSecurityGroups(&awsec2.DescribeSecurityGroupsInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing security groups: %s", err)
	}

	var resources []SecurityGroup
	for _, sg := range output.SecurityGroups {
		resource := NewSecurityGroup(e.client, sg.GroupId, sg.GroupName, sg.Tags)

		if *sg.GroupName == "default" {
			continue
		}

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete security group %s?", resource.identifier))
		if !proceed {
			continue
		}

		e.revoke(sg)

		resources = append(resources, resource)
	}

	return resources, nil
}

func (s SecurityGroups) Delete(securityGroups map[string]string) error {
	var resources []SecurityGroup
	for name, id := range securityGroups {
		resources = append(resources, NewSecurityGroup(s.client, &id, &name, []*awsec2.Tag{}))
	}

	return s.cleanup(resources)
}

func (s SecurityGroups) cleanup(resources []SecurityGroup) error {
	for _, resource := range resources {
		err := resource.Delete()

		if err == nil {
			s.logger.Printf("SUCCESS deleting security group %s\n", resource.identifier)
		} else {
			s.logger.Printf("ERROR deleting security group %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}

func (e SecurityGroups) revoke(s *awsec2.SecurityGroup) {
	if len(s.IpPermissions) > 0 {
		_, err := e.client.RevokeSecurityGroupIngress(&awsec2.RevokeSecurityGroupIngressInput{
			GroupId:       s.GroupId,
			IpPermissions: s.IpPermissions,
		})
		if err != nil {
			e.logger.Printf("ERROR revoking security group ingress for %s: %s\n", *s.GroupName, err)
		}
	}

	if len(s.IpPermissionsEgress) > 0 {
		_, err := e.client.RevokeSecurityGroupEgress(&awsec2.RevokeSecurityGroupEgressInput{
			GroupId:       s.GroupId,
			IpPermissions: s.IpPermissionsEgress,
		})
		if err != nil {
			e.logger.Printf("ERROR revoking security group egress for %s: %s\n", *s.GroupName, err)
		}
	}
}
