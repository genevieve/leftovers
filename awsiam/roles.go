package awsiam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

type Roles struct {
	client iamClient
	logger logger
}

func NewRoles(client iamClient, logger logger) Roles {
	return Roles{
		client: client,
		logger: logger,
	}
}

func (o Roles) Delete() error {
	roles, err := o.client.ListRoles(&iam.ListRolesInput{})
	if err != nil {
		return fmt.Errorf("Listing roles: %s", err)
	}

	for _, r := range roles.Roles {
		n := r.RoleName
		_, err := o.client.DeleteRole(&iam.DeleteRoleInput{RoleName: n})
		if err == nil {
			o.logger.Printf("SUCCESS deleting role %s\n", *n)
		} else {
			o.logger.Printf("ERROR deleting role %s: %s\n", *n, err)
		}
	}

	return nil
}
