package awsiam

import "github.com/aws/aws-sdk-go/service/iam"

type iamClient interface {
	ListInstanceProfiles(*iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)
	DeleteInstanceProfile(*iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)

	ListServerCertificates(*iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)
	DeleteServerCertificate(*iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)

	ListRoles(*iam.ListRolesInput) (*iam.ListRolesOutput, error)
	DeleteRole(*iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)
}
