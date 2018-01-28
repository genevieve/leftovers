package iam

import awsiam "github.com/aws/aws-sdk-go/service/iam"

type Role struct {
	client     rolesClient
	name       *string
	identifier string
}

func NewRole(client rolesClient, name *string) Role {
	return Role{
		client:     client,
		name:       name,
		identifier: *name,
	}
}

func (r Role) Delete() error {
	_, err := r.client.DeleteRole(&awsiam.DeleteRoleInput{
		RoleName: r.name,
	})
	return err
}
