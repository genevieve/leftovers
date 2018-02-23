package iam

import (
	"fmt"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type PolicyVersion struct {
	client     policiesClient
	arn        *string
	version    *string
	identifier string
}

func NewPolicyVersion(client policiesClient, policyName, arn, version *string) PolicyVersion {
	return PolicyVersion{
		client:     client,
		arn:        arn,
		version:    version,
		identifier: fmt.Sprintf("%s-%s", *policyName, *version),
	}
}

func (p PolicyVersion) Delete() error {
	_, err := p.client.DeletePolicyVersion(&awsiam.DeletePolicyVersionInput{
		PolicyArn: p.arn,
		VersionId: p.version,
	})

	if err != nil {
		return fmt.Errorf("FAILED deleting policy version %s: %s", p.identifier, err)
	}

	return nil
}

func (p PolicyVersion) Name() string {
	return p.identifier
}
