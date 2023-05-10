package iam

import (
	"fmt"
	awsiam "github.com/aws/aws-sdk-go/service/iam"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface instanceProfilesClient --output fakes/instance_profiles_client.go
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

func (i InstanceProfiles) List(filter string, regex bool) ([]common.Deletable, error) {
	profiles, err := i.client.ListInstanceProfiles(&awsiam.ListInstanceProfilesInput{})
	if err != nil {
		return nil, fmt.Errorf("List IAM Instance Profiles: %s", err)
	}

	var resources []common.Deletable
	for _, p := range profiles.InstanceProfiles {
		r := NewInstanceProfile(i.client, p.InstanceProfileName, p.Roles, i.logger)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := i.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (i InstanceProfiles) Type() string {
	return "iam-instance-profile"
}
