package iam

import (
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type iamClient interface {
	ListInstanceProfiles(*awsiam.ListInstanceProfilesInput) (*awsiam.ListInstanceProfilesOutput, error)
	DeleteInstanceProfile(*awsiam.DeleteInstanceProfileInput) (*awsiam.DeleteInstanceProfileOutput, error)

	ListServerCertificates(*awsiam.ListServerCertificatesInput) (*awsiam.ListServerCertificatesOutput, error)
	DeleteServerCertificate(*awsiam.DeleteServerCertificateInput) (*awsiam.DeleteServerCertificateOutput, error)

	ListRoles(*awsiam.ListRolesInput) (*awsiam.ListRolesOutput, error)
	DeleteRole(*awsiam.DeleteRoleInput) (*awsiam.DeleteRoleOutput, error)

	ListRolePolicies(*awsiam.ListRolePoliciesInput) (*awsiam.ListRolePoliciesOutput, error)
	DeleteRolePolicy(*awsiam.DeleteRolePolicyInput) (*awsiam.DeleteRolePolicyOutput, error)
}
