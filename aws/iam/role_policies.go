package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type RolePolicies struct {
	client iamClient
	logger logger
}

type rolePolicies interface {
	Delete(roleName string) error
}

func NewRolePolicies(client iamClient, logger logger) RolePolicies {
	return RolePolicies{
		client: client,
		logger: logger,
	}
}

func (o RolePolicies) Delete(roleName string) error {
	policies, err := o.client.ListRolePolicies(&awsiam.ListRolePoliciesInput{RoleName: aws.String(roleName)})
	if err != nil {
		return fmt.Errorf("Listing role policies: %s", err)
	}

	for _, p := range policies.PolicyNames {
		n := *p

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete role policy %s?", n))
		if !proceed {
			continue
		}

		_, err := o.client.DeleteRolePolicy(&awsiam.DeleteRolePolicyInput{
			RoleName:   aws.String(roleName),
			PolicyName: p,
		})
		if err == nil {
			o.logger.Printf("SUCCESS deleting role policy %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting role policy %s: %s\n", n, err)
		}
	}

	return nil
}
