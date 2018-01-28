package iam

type User struct {
	client     usersClient
	name       *string
	identifier string
}

func NewUser(client usersClient, name *string) User {
	return User{
		client:     client,
		name:       name,
		identifier: *name,
	}
}
