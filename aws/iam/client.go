package iam

import (
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type iamClient interface {
	ListRolePolicies(*awsiam.ListRolePoliciesInput) (*awsiam.ListRolePoliciesOutput, error)
	ListPolicies(*awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error)
	DetachRolePolicy(*awsiam.DetachRolePolicyInput) (*awsiam.DetachRolePolicyOutput, error)
	DeleteRolePolicy(*awsiam.DeleteRolePolicyInput) (*awsiam.DeleteRolePolicyOutput, error)
}
