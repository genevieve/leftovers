package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type SecurityGroup struct {
	client     securityGroupsClient
	logger     logger
	id         *string
	identifier string
	rtype      string
	ingress    []*awsec2.IpPermission
	egress     []*awsec2.IpPermission
}

func NewSecurityGroup(client securityGroupsClient, logger logger, id, groupName *string, tags []*awsec2.Tag, ingress []*awsec2.IpPermission, egress []*awsec2.IpPermission) SecurityGroup {
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
		logger:     logger,
		id:         id,
		identifier: identifier,
		rtype:      "EC2 Security Group",
		ingress:    ingress,
		egress:     egress,
	}
}

//TODO: Retryable error - DependencyViolation
func (s SecurityGroup) Delete() error {
	if len(s.ingress) > 0 {
		_, err := s.client.RevokeSecurityGroupIngress(&awsec2.RevokeSecurityGroupIngressInput{
			GroupId:       s.id,
			IpPermissions: s.ingress,
		})
		if err != nil {
			s.logger.Printf("[%s: %s] Revoke ingress: %s", s.Type(), s.Name(), err)
		} else {
			s.logger.Printf("[%s: %s] Revoked ingress", s.Type(), s.Name())
		}
	}

	if len(s.egress) > 0 {
		_, err := s.client.RevokeSecurityGroupEgress(&awsec2.RevokeSecurityGroupEgressInput{
			GroupId:       s.id,
			IpPermissions: s.egress,
		})
		if err != nil {
			s.logger.Printf("[%s: %s] Revoke egress: %s", s.Type(), s.Name(), err)
		} else {
			s.logger.Printf("[%s: %s] Revoked egress", s.Type(), s.Name())
		}
	}

	_, err := s.client.DeleteSecurityGroup(&awsec2.DeleteSecurityGroupInput{GroupId: s.id})
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (s SecurityGroup) Name() string {
	return s.identifier
}

func (s SecurityGroup) Type() string {
	return "EC2 Security Group"
}
