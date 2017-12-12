package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type userPoliciesClient interface {
	ListUserPolicies(*awsiam.ListUserPoliciesInput) (*awsiam.ListUserPoliciesOutput, error)
	ListPolicies(*awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error)
	DetachUserPolicy(*awsiam.DetachUserPolicyInput) (*awsiam.DetachUserPolicyOutput, error)
	DeleteUserPolicy(*awsiam.DeleteUserPolicyInput) (*awsiam.DeleteUserPolicyOutput, error)
}

type userPolicies interface {
	Delete(userName string) error
}

type UserPolicies struct {
	client userPoliciesClient
	logger logger
}

func NewUserPolicies(client userPoliciesClient, logger logger) UserPolicies {
	return UserPolicies{
		client: client,
		logger: logger,
	}
}

func (o UserPolicies) Delete(userName string) error {
	policies, err := o.client.ListUserPolicies(&awsiam.ListUserPoliciesInput{UserName: aws.String(userName)})
	if err != nil {
		return fmt.Errorf("Listing user policies: %s", err)
	}

	for _, p := range policies.PolicyNames {
		n := *p

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete user policy %s?", n))
		if !proceed {
			continue
		}

		o.detach(n, userName)

		_, err = o.client.DeleteUserPolicy(&awsiam.DeleteUserPolicyInput{
			UserName:   aws.String(userName),
			PolicyName: p,
		})
		if err == nil {
			o.logger.Printf("SUCCESS deleting user policy %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting user policy %s: %s\n", n, err)
		}
	}

	return nil
}

func (o UserPolicies) detach(n, userName string) {
	policies, err := o.client.ListPolicies(&awsiam.ListPoliciesInput{
		Scope: aws.String("Local"),
	})
	if err == nil {
		for _, policy := range policies.Policies {
			if *policy.PolicyName == n {
				_, err = o.client.DetachUserPolicy(&awsiam.DetachUserPolicyInput{
					UserName:  aws.String(userName),
					PolicyArn: policy.Arn,
				})
				if err == nil {
					o.logger.Printf("SUCCESS detaching user policy %s\n", n)
				} else {
					o.logger.Printf("ERROR detaching user policy %s: %s\n", n, err)
				}
			}
		}
	} else {
		o.logger.Printf("ERROR getting user policy %s: %s\n", n, err)
	}
}
