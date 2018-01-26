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
	delete := map[string]string{}

	profiles, err := i.client.ListInstanceProfiles(&awsiam.ListInstanceProfilesInput{})
	if err != nil {
		return delete, fmt.Errorf("Listing instance profiles: %s", err)
	}

	for _, p := range profiles.InstanceProfiles {
		n := *p.InstanceProfileName

		clearerName := i.clearerName(n, p.Roles)

		if !strings.Contains(clearerName, filter) {
			continue
		}

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete instance profile %s?", clearerName))
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
				i.logger.Printf("SUCCESS removing role %s from instance profile %s\n", role, clearerName)
			} else {
				i.logger.Printf("ERROR removing role %s from instance profile %s: %s\n", role, clearerName, err)
			}
		}

		delete[n] = ""
	}

	return delete, nil
}

func (i InstanceProfiles) Delete(profiles map[string]string) error {
	for name, _ := range profiles {
		_, err := i.client.DeleteInstanceProfile(&awsiam.DeleteInstanceProfileInput{InstanceProfileName: aws.String(name)})

		if err == nil {
			i.logger.Printf("SUCCESS deleting instance profile %s\n", name)
		} else {
			i.logger.Printf("ERROR deleting instance profile %s: %s\n", name, err)
		}
	}

	return nil
}

func (i InstanceProfiles) clearerName(name string, roles []*awsiam.Role) string {
	extra := []string{}
	for _, r := range roles {
		extra = append(extra, fmt.Sprintf("Role:%s", *r.RoleName))
	}

	if len(extra) > 0 {
		return fmt.Sprintf("%s (%s)", name, strings.Join(extra, ", "))
	}

	return name
}
