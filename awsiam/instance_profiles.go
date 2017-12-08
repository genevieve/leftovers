package awsiam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

type iamClient interface {
	ListInstanceProfiles(*iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)
	DeleteInstanceProfile(*iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)
}

type InstanceProfiles struct {
	client iamClient
}

func NewInstanceProfiles(client iamClient) InstanceProfiles {
	return InstanceProfiles{
		client: client,
	}
}

func (i InstanceProfiles) Delete() {
	profiles, err := i.client.ListInstanceProfiles(&iam.ListInstanceProfilesInput{})
	if err != nil {
		fmt.Printf("ERROR listing instance profiles: %s", err)
	}

	for _, p := range profiles.InstanceProfiles {
		n := p.InstanceProfileName
		_, err := i.client.DeleteInstanceProfile(&iam.DeleteInstanceProfileInput{InstanceProfileName: n})
		if err == nil {
			fmt.Printf("SUCCESS deleting instance profile %s\n", *n)
		} else {
			fmt.Printf("ERROR deleting instance profile %s: %s\n", *n, err)
		}
	}
}
