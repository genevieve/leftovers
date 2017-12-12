package iam

import (
	"fmt"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type instanceProfilesClient interface {
	ListInstanceProfiles(*awsiam.ListInstanceProfilesInput) (*awsiam.ListInstanceProfilesOutput, error)
	RemoveRoleFromInstanceProfile(*awsiam.RemoveRoleFromInstanceProfileInput) (*awsiam.RemoveRoleFromInstanceProfileOutput, error)
	DeleteInstanceProfile(*awsiam.DeleteInstanceProfileInput) (*awsiam.DeleteInstanceProfileOutput, error)
}

type InstanceProfiles struct {
	client instanceProfilesClient
	logger logger
}

func NewInstanceProfiles(client instanceProfilesClient, logger logger) InstanceProfiles {
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
		n := *p.InstanceProfileName

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete instance profile %s?", n))
		if !proceed {
			continue
		}

		for _, r := range p.Roles {
			role := *r.RoleName

			_, err := i.client.RemoveRoleFromInstanceProfile(&awsiam.RemoveRoleFromInstanceProfileInput{
				InstanceProfileName: p.InstanceProfileName,
				RoleName:            r.RoleName,
			})
			if err == nil {
				i.logger.Printf("SUCCESS removing role %s from instance profile %s\n", role, n)
			} else {
				i.logger.Printf("ERROR removing role %s from instance profile %s: %s\n", role, n, err)
			}
		}

		_, err := i.client.DeleteInstanceProfile(&awsiam.DeleteInstanceProfileInput{InstanceProfileName: p.InstanceProfileName})
		if err == nil {
			i.logger.Printf("SUCCESS deleting instance profile %s\n", n)
		} else {
			i.logger.Printf("ERROR deleting instance profile %s: %s\n", n, err)
		}
	}

	return nil
}
