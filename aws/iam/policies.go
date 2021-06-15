package iam

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface policiesClient --output fakes/policies_client.go
type policiesClient interface {
	ListPolicies(*awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error)
	ListPolicyVersions(*awsiam.ListPolicyVersionsInput) (*awsiam.ListPolicyVersionsOutput, error)
	DeletePolicyVersion(*awsiam.DeletePolicyVersionInput) (*awsiam.DeletePolicyVersionOutput, error)
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

func (p Policies) List(filter string, regex bool) ([]common.Deletable, error) {
	policies, err := p.client.ListPolicies(&awsiam.ListPoliciesInput{Scope: aws.String("Local")})
	if err != nil {
		return nil, fmt.Errorf("List IAM Policies: %s", err)
	}

	var resources []common.Deletable
	for _, o := range policies.Policies {
		r := NewPolicy(p.client, p.logger, o.PolicyName, o.Arn)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := p.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (p Policies) Type() string {
	return "iam-policy"
}
