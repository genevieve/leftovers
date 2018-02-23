package iam

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevieve/leftovers/aws/common"
)

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

func (p Policies) List(filter string) ([]common.Deletable, error) {
	policies, err := p.client.ListPolicies(&awsiam.ListPoliciesInput{Scope: aws.String("Local")})
	if err != nil {
		return nil, fmt.Errorf("Listing policies: %s", err)
	}

	var resources []common.Deletable
	for _, o := range policies.Policies {
		resource := NewPolicy(p.client, o.PolicyName, o.Arn)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := p.logger.Prompt(fmt.Sprintf("Are you sure you want to delete policy %s?", resource.identifier))
		if !proceed {
			continue
		}

		versions, err := p.client.ListPolicyVersions(&awsiam.ListPolicyVersionsInput{PolicyArn: o.Arn})
		if err != nil {
			return nil, fmt.Errorf("Listing policy versions: %s", err)
		}

		for _, v := range versions.Versions {
			if !*v.IsDefaultVersion {
				versionResource := NewPolicyVersion(p.client, o.PolicyName, o.Arn, v.VersionId)
				resources = append(resources, versionResource)
			}
		}

		resources = append(resources, resource)
	}

	return resources, nil
}
