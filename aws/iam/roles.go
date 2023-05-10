package iam

import (
	"fmt"
	awsiam "github.com/aws/aws-sdk-go/service/iam"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface rolesClient --output fakes/roles_client.go
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

func (o Roles) List(filter string, regex bool) ([]common.Deletable, error) {
	roles, err := o.client.ListRoles(&awsiam.ListRolesInput{})
	if err != nil {
		return nil, fmt.Errorf("List IAM Roles: %s", err)
	}

	var resources []common.Deletable
	for _, role := range roles.Roles {
		r := NewRole(o.client, o.policies, role.RoleName)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := o.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (o Roles) Type() string {
	return "iam-role"
}
