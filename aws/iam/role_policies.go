package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type rolePoliciesClient interface {
	ListAttachedRolePolicies(*awsiam.ListAttachedRolePoliciesInput) (*awsiam.ListAttachedRolePoliciesOutput, error)
	DetachRolePolicy(*awsiam.DetachRolePolicyInput) (*awsiam.DetachRolePolicyOutput, error)
	DeleteRolePolicy(*awsiam.DeleteRolePolicyInput) (*awsiam.DeleteRolePolicyOutput, error)
}

type rolePolicies interface {
	Delete(roleName string) error
}

type RolePolicies struct {
	client rolePoliciesClient
	logger logger
}

func NewRolePolicies(client rolePoliciesClient, logger logger) RolePolicies {
	return RolePolicies{
		client: client,
		logger: logger,
	}
}

func (o RolePolicies) Delete(roleName string) error {
	policies, err := o.client.ListAttachedRolePolicies(&awsiam.ListAttachedRolePoliciesInput{RoleName: aws.String(roleName)})
	if err != nil {
		return fmt.Errorf("Listing role policies: %s", err)
	}

	for _, p := range policies.AttachedPolicies {
		n := *p.PolicyName

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete role policy %s?", n))
		if !proceed {
			continue
		}

		_, err := o.client.DetachRolePolicy(&awsiam.DetachRolePolicyInput{
			RoleName:  aws.String(roleName),
			PolicyArn: p.PolicyArn,
		})
		if err == nil {
			o.logger.Printf("SUCCESS detaching role policy %s\n", n)
		} else {
			o.logger.Printf("ERROR detaching role policy %s: %s\n", n, err)
		}

		_, err = o.client.DeleteRolePolicy(&awsiam.DeleteRolePolicyInput{
			RoleName:   aws.String(roleName),
			PolicyName: p.PolicyName,
		})
		if err == nil {
			o.logger.Printf("SUCCESS deleting role policy %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting role policy %s: %s\n", n, err)
		}
	}

	return nil
}
