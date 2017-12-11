package ec2

import (
	"fmt"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type SecurityGroups struct {
	client ec2Client
	logger logger
}

func NewSecurityGroups(client ec2Client, logger logger) SecurityGroups {
	return SecurityGroups{
		client: client,
		logger: logger,
	}
}

func (e SecurityGroups) Delete() error {
	groups, err := e.client.DescribeSecurityGroups(&awsec2.DescribeSecurityGroupsInput{})
	if err != nil {
		return fmt.Errorf("Describing security groups: %s", err)
	}

	for _, s := range groups.SecurityGroups {
		n := *s.GroupName

		if n == "default" {
			continue
		}

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete security group %s?", n))
		if !proceed {
			continue
		}

		e.revoke(s)

		_, err := e.client.DeleteSecurityGroup(&awsec2.DeleteSecurityGroupInput{
			GroupId: s.GroupId,
		})
		if err == nil {
			e.logger.Printf("SUCCESS deleting security group %s\n", n)
		} else {
			//List any security groups that mention this security group
			//Prompt if they are okay revoking rules from all these groups
			//Delete the one group
			e.logger.Printf("ERROR deleting security group %s: %s\n", n, err)
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
