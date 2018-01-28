package iam

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type usersClient interface {
	ListUsers(*awsiam.ListUsersInput) (*awsiam.ListUsersOutput, error)
	DeleteUser(*awsiam.DeleteUserInput) (*awsiam.DeleteUserOutput, error)
}

type Users struct {
	client     usersClient
	logger     logger
	policies   userPolicies
	accessKeys accessKeys
}

func NewUsers(client usersClient, logger logger, policies userPolicies, accessKeys accessKeys) Users {
	return Users{
		client:     client,
		logger:     logger,
		policies:   policies,
		accessKeys: accessKeys,
	}
}

func (u Users) List(filter string) (map[string]string, error) {
	users, err := u.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, r := range users {
		delete[*r.name] = ""
	}

	return delete, nil
}

func (u Users) list(filter string) ([]User, error) {
	users, err := u.client.ListUsers(&awsiam.ListUsersInput{})
	if err != nil {
		return nil, fmt.Errorf("Listing users: %s", err)
	}

	var resources []User
	for _, r := range users.Users {
		resource := NewUser(u.client, r.UserName)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := u.logger.Prompt(fmt.Sprintf("Are you sure you want to delete user %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (u Users) Delete(users map[string]string) error {
	for name, _ := range users {
		err := u.accessKeys.Delete(name)
		if err != nil {
			return fmt.Errorf("Deleting access keys for %s: %s", name, err)
		}

		err = u.policies.Delete(name)
		if err != nil {
			return fmt.Errorf("Deleting policies for %s: %s", name, err)
		}

		_, err = u.client.DeleteUser(&awsiam.DeleteUserInput{
			UserName: aws.String(name),
		})
		if err == nil {
			u.logger.Printf("SUCCESS deleting user %s\n", name)
		} else {
			u.logger.Printf("ERROR deleting user %s: %s\n", name, err)
		}
	}

	return nil
}
