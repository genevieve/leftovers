package iam

import (
	"fmt"
	"strings"

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

func (p Policies) List(filter string) (map[string]string, error) {
	policies, err := p.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, o := range policies {
		delete[*o.name] = *o.arn
	}

	return delete, nil
}

func (p Policies) list(filter string) ([]Policy, error) {
	policies, err := p.client.ListPolicies(&awsiam.ListPoliciesInput{Scope: aws.String("Local")})
	if err != nil {
		return nil, fmt.Errorf("Listing policies: %s", err)
	}

	var resources []Policy
	for _, o := range policies.Policies {
		resource := NewPolicy(p.client, o.PolicyName, o.Arn)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := p.logger.Prompt(fmt.Sprintf("Are you sure you want to delete policy %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (p Policies) Delete(policies map[string]string) error {
	for name, arn := range policies {
		_, err := p.client.DeletePolicy(&awsiam.DeletePolicyInput{PolicyArn: aws.String(arn)})

		if err == nil {
			p.logger.Printf("SUCCESS deleting policy %s\n", name)
		} else {
			p.logger.Printf("ERROR deleting policy %s: %s\n", name, err)
		}
	}

	return nil
}
