package iam

import (
	"fmt"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
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
	profiles, err := i.client.ListInstanceProfiles(&awsiam.ListInstanceProfilesInput{})
	if err != nil {
		return fmt.Errorf("Listing instance profiles: %s", err)
	}

	for _, p := range profiles.InstanceProfiles {
		n := p.InstanceProfileName

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete instance profile %s?", *n))
		if !proceed {
			continue
		}

		_, err := i.client.DeleteInstanceProfile(&awsiam.DeleteInstanceProfileInput{InstanceProfileName: n})
		if err == nil {
			i.logger.Printf("SUCCESS deleting instance profile %s\n", *n)
		} else {
			i.logger.Printf("ERROR deleting instance profile %s: %s\n", *n, err)
		}
	}

	return nil
}
