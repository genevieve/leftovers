package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type rolePoliciesClient interface {
	ListRolePolicies(*awsiam.ListRolePoliciesInput) (*awsiam.ListRolePoliciesOutput, error)
	ListPolicies(*awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error)
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

		o.detach(n, roleName)

		_, err = o.client.DeleteRolePolicy(&awsiam.DeleteRolePolicyInput{
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

func (o RolePolicies) detach(n, roleName string) {
	policies, err := o.client.ListPolicies(&awsiam.ListPoliciesInput{
		Scope: aws.String("Local"),
	})
	if err == nil {
		for _, policy := range policies.Policies {
			if *policy.PolicyName == n {
				_, err = o.client.DetachRolePolicy(&awsiam.DetachRolePolicyInput{
					RoleName:  aws.String(roleName),
					PolicyArn: policy.Arn,
				})
				if err == nil {
					o.logger.Printf("SUCCESS detaching role policy %s\n", n)
				} else {
					o.logger.Printf("ERROR detaching role policy %s: %s\n", n, err)
				}
			}
		}
	} else {
		o.logger.Printf("ERROR getting role policy %s: %s\n", n, err)
	}
}
