package iam

import (
	"fmt"
	"strings"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type rolesClient interface {
	ListRoles(*awsiam.ListRolesInput) (*awsiam.ListRolesOutput, error)
	DeleteRole(*awsiam.DeleteRoleInput) (*awsiam.DeleteRoleOutput, error)
}

type Roles struct {
	client   rolesClient
	logger   logger
	policies rolePolicies
}

func NewRoles(client rolesClient, logger logger, policies rolePolicies) Roles {
	return Roles{
		client:   client,
		logger:   logger,
		policies: policies,
	}
}

func (r Roles) List(filter string) (map[string]string, error) {
	roles, err := r.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, role := range roles {
		delete[*role.name] = ""
	}

	return delete, nil
}

func (r Roles) list(filter string) ([]Role, error) {
	roles, err := r.client.ListRoles(&awsiam.ListRolesInput{})
	if err != nil {
		return nil, fmt.Errorf("Listing roles: %s", err)
	}

	var resources []Role
	for _, role := range roles.Roles {
		resource := NewRole(r.client, role.RoleName)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := r.logger.Prompt(fmt.Sprintf("Are you sure you want to delete role %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (r Roles) Delete(roles map[string]string) error {
	var resources []Role
	for name, _ := range roles {
		resources = append(resources, NewRole(r.client, &name))
	}

	return r.cleanup(resources)
}

func (r Roles) cleanup(resources []Role) error {
	for _, resource := range resources {
		err := r.policies.Delete(*resource.name)
		if err != nil {
			return fmt.Errorf("Deleting policies for %s: %s", resource.identifier, err)
		}

		err = resource.Delete()
		if err == nil {
			r.logger.Printf("SUCCESS deleting role %s\n", resource.identifier)
		} else {
			r.logger.Printf("ERROR deleting role %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}
