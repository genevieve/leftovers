package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type IAMClient struct {
	ListInstanceProfilesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListInstanceProfilesInput
		}
		Returns struct {
			Output *iam.ListInstanceProfilesOutput
			Error  error
		}
	}

	DeleteInstanceProfileCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteInstanceProfileInput
		}
		Returns struct {
			Output *iam.DeleteInstanceProfileOutput
			Error  error
		}
	}

	ListServerCertificatesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListServerCertificatesInput
		}
		Returns struct {
			Output *iam.ListServerCertificatesOutput
			Error  error
		}
	}

	DeleteServerCertificateCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteServerCertificateInput
		}
		Returns struct {
			Output *iam.DeleteServerCertificateOutput
			Error  error
		}
	}
}

func (i *IAMClient) ListInstanceProfiles(input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {
	i.ListInstanceProfilesCall.CallCount++
	i.ListInstanceProfilesCall.Receives.Input = input

	return i.ListInstanceProfilesCall.Returns.Output, i.ListInstanceProfilesCall.Returns.Error
}

func (i *IAMClient) DeleteInstanceProfile(input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {
	i.DeleteInstanceProfileCall.CallCount++
	i.DeleteInstanceProfileCall.Receives.Input = input

	return i.DeleteInstanceProfileCall.Returns.Output, i.DeleteInstanceProfileCall.Returns.Error
}

func (i *IAMClient) ListServerCertificates(input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
	i.ListServerCertificatesCall.CallCount++
	i.ListServerCertificatesCall.Receives.Input = input

	return i.ListServerCertificatesCall.Returns.Output, i.ListServerCertificatesCall.Returns.Error
}

func (i *IAMClient) DeleteServerCertificate(input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {
	i.DeleteServerCertificateCall.CallCount++
	i.DeleteServerCertificateCall.Receives.Input = input

	return i.DeleteServerCertificateCall.Returns.Output, i.DeleteServerCertificateCall.Returns.Error
}
