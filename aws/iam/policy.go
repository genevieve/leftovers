package iam

import awsiam "github.com/aws/aws-sdk-go/service/iam"

type Policy struct {
	client     policiesClient
	name       *string
	arn        *string
	identifier string
}

func NewPolicy(client policiesClient, name, arn *string) Policy {
	return Policy{
		client:     client,
		name:       name,
		arn:        arn,
		identifier: *name,
	}
}

func (p Policy) Delete() error {
	_, err := p.client.DeletePolicy(&awsiam.DeletePolicyInput{
		PolicyArn: p.arn,
	})
	return err
}
