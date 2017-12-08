package awsiam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

type Roles struct {
	client iamClient
}

func NewRoles(client iamClient) Roles {
	return Roles{
		client: client,
	}
}

func (o Roles) Delete() {
	roles, err := o.client.ListRoles(&iam.ListRolesInput{})
	if err != nil {
		fmt.Errorf("ERROR listing roles: %s", err)
	}

	for _, r := range roles.Roles {
		n := r.RoleName
		_, err := o.client.DeleteRole(&iam.DeleteRoleInput{RoleName: n})
		if err == nil {
			fmt.Printf("SUCCESS deleting role %s\n", *n)
		} else {
			fmt.Printf("ERROR deleting role %s: %s\n", *n, err)
		}
	}
}
