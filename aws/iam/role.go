package iam

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
