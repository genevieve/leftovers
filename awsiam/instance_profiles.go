package awsiam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

type InstanceProfiles struct {
	client iamClient
	logger logger
}

func NewInstanceProfiles(client iamClient, logger logger) InstanceProfiles {
	return InstanceProfiles{
		client: client,
		logger: logger,
	}
}

func (i InstanceProfiles) Delete() error {
	profiles, err := i.client.ListInstanceProfiles(&iam.ListInstanceProfilesInput{})
	if err != nil {
		return fmt.Errorf("Listing instance profiles: %s", err)
	}

	for _, p := range profiles.InstanceProfiles {
		n := p.InstanceProfileName
		_, err := i.client.DeleteInstanceProfile(&iam.DeleteInstanceProfileInput{InstanceProfileName: n})
		if err == nil {
			i.logger.Printf("SUCCESS deleting instance profile %s\n", *n)
		} else {
			i.logger.Printf("ERROR deleting instance profile %s: %s\n", *n, err)
		}
	}

	return nil
}
