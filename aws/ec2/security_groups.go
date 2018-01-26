package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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
	delete := map[string]string{}

	groups, err := e.client.DescribeSecurityGroups(&awsec2.DescribeSecurityGroupsInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing security groups: %s", err)
	}

	for _, s := range groups.SecurityGroups {
		if *s.GroupName == "default" {
			continue
		}

		n := e.clearerName(*s.GroupName, s.Tags)

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete security group %s?", n))
		if !proceed {
			continue
		}

		e.revoke(s)

		delete[n] = *s.GroupId
	}

	return delete, nil
}

func (s SecurityGroups) Delete(securityGroups map[string]string) error {
	for name, id := range securityGroups {
		_, err := s.client.DeleteSecurityGroup(&awsec2.DeleteSecurityGroupInput{
			GroupId: aws.String(id),
		})

		if err == nil {
			s.logger.Printf("SUCCESS deleting security group %s\n", name)
		} else {
			s.logger.Printf("ERROR deleting security group %s: %s\n", name, err)
		}
	}

	return nil
}

func (e SecurityGroups) clearerName(n string, tags []*awsec2.Tag) string {
	extra := []string{}
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		return fmt.Sprintf("%s (%s)", n, strings.Join(extra, ", "))
	}

	return n
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
