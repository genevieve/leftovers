package iam

import (
	"fmt"

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

func (o Roles) Delete() error {
	roles, err := o.client.ListRoles(&awsiam.ListRolesInput{})
	if err != nil {
		return fmt.Errorf("Listing roles: %s", err)
	}

	for _, r := range roles.Roles {
		n := *r.RoleName

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete role %s?", n))
		if !proceed {
			continue
		}

		if err := o.policies.Delete(n); err != nil {
			return fmt.Errorf("Deleting policies for %s: %s", n, err)
		}

		_, err = o.client.DeleteRole(&awsiam.DeleteRoleInput{RoleName: r.RoleName})
		if err == nil {
			o.logger.Printf("SUCCESS deleting role %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting role %s: %s\n", n, err)
		}
	}

	return nil
}
