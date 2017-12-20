package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type policiesClient interface {
	ListPolicies(*awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error)
	DeletePolicy(*awsiam.DeletePolicyInput) (*awsiam.DeletePolicyOutput, error)
}

type Policies struct {
	client policiesClient
	logger logger
}

func NewPolicies(client policiesClient, logger logger) Policies {
	return Policies{
		client: client,
		logger: logger,
	}
}

func (o Policies) Delete() error {
	policies, err := o.client.ListPolicies(&awsiam.ListPoliciesInput{Scope: aws.String("Local")})
	if err != nil {
		return fmt.Errorf("Listing policies: %s", err)
	}

	for _, p := range policies.Policies {
		n := *p.PolicyName

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete policy %s?", n))
		if !proceed {
			continue
		}

		_, err = o.client.DeletePolicy(&awsiam.DeletePolicyInput{PolicyArn: p.Arn})
		if err == nil {
			o.logger.Printf("SUCCESS deleting policy %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting policy %s: %s\n", n, err)
		}
	}

	return nil
}
