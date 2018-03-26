package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type SecurityGroup struct {
	client     securityGroupsClient
	id         *string
	identifier string
	ingress    []*awsec2.IpPermission
	egress     []*awsec2.IpPermission
}

func NewSecurityGroup(client securityGroupsClient, id, groupName *string, tags []*awsec2.Tag, ingress []*awsec2.IpPermission, egress []*awsec2.IpPermission) SecurityGroup {
	identifier := *groupName

	var extra []string
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		identifier = fmt.Sprintf("%s (%s)", *groupName, strings.Join(extra, ", "))
	}

	return SecurityGroup{
		client:     client,
		id:         id,
		identifier: identifier,
		ingress:    ingress,
		egress:     egress,
	}
}

func (s SecurityGroup) Delete() error {
	err := s.revoke()
	if err != nil {
		return err
	}

	_, err = s.client.DeleteSecurityGroup(&awsec2.DeleteSecurityGroupInput{
		GroupId: s.id,
	})

	if err != nil {
		return fmt.Errorf("FAILED deleting security group %s: %s", s.identifier, err)
	}

	return nil
}

func (s SecurityGroup) Name() string {
	return s.identifier
}

func (s SecurityGroup) Type() string {
	return "security group"
}

func (s SecurityGroup) revoke() error {
	if len(s.ingress) > 0 {
		_, err := s.client.RevokeSecurityGroupIngress(&awsec2.RevokeSecurityGroupIngressInput{
			GroupId:       s.id,
			IpPermissions: s.ingress,
		})
		if err != nil {
			return fmt.Errorf("ERROR revoking security group ingress for %s: %s\n", s.identifier, err)
		}
	}

	if len(s.egress) > 0 {
		_, err := s.client.RevokeSecurityGroupEgress(&awsec2.RevokeSecurityGroupEgressInput{
			GroupId:       s.id,
			IpPermissions: s.egress,
		})
		if err != nil {
			return fmt.Errorf("ERROR revoking security group egress for %s: %s\n", s.identifier, err)
		}
	}

	return nil
}
