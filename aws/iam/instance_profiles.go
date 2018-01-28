package iam

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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

func (i InstanceProfiles) List(filter string) (map[string]string, error) {
	profiles, err := i.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, p := range profiles {
		delete[*p.name] = ""
	}

	return delete, nil
}

func (i InstanceProfiles) list(filter string) ([]InstanceProfile, error) {
	profiles, err := i.client.ListInstanceProfiles(&awsiam.ListInstanceProfilesInput{})
	if err != nil {
		return nil, fmt.Errorf("Listing instance profiles: %s", err)
	}

	var resources []InstanceProfile
	for _, p := range profiles.InstanceProfiles {
		resource := NewInstanceProfile(i.client, p.InstanceProfileName, p.Roles)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete instance profile %s?", resource.identifier))
		if !proceed {
			continue
		}

		for _, r := range p.Roles {
			role := *r.RoleName

			_, err := i.client.RemoveRoleFromInstanceProfile(&awsiam.RemoveRoleFromInstanceProfileInput{
				InstanceProfileName: resource.name,
				RoleName:            r.RoleName,
			})
			if err == nil {
				i.logger.Printf("SUCCESS removing role %s from instance profile %s\n", role, resource.identifier)
			} else {
				i.logger.Printf("ERROR removing role %s from instance profile %s: %s\n", role, resource.identifier, err)
			}
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (i InstanceProfiles) Delete(profiles map[string]string) error {
	for name, _ := range profiles {
		_, err := i.client.DeleteInstanceProfile(&awsiam.DeleteInstanceProfileInput{
			InstanceProfileName: aws.String(name),
		})

		if err == nil {
			i.logger.Printf("SUCCESS deleting instance profile %s\n", name)
		} else {
			i.logger.Printf("ERROR deleting instance profile %s: %s\n", name, err)
		}
	}

	return nil
}
