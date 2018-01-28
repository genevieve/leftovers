package iam

import awsiam "github.com/aws/aws-sdk-go/service/iam"

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

func (u User) Delete() error {
	_, err := u.client.DeleteUser(&awsiam.DeleteUserInput{
		UserName: u.name,
	})
	return err
}
