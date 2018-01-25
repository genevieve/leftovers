package iam

import (
	"fmt"
	"strings"

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

func (o Users) Delete(filter string) error {
	users, err := o.client.ListUsers(&awsiam.ListUsersInput{})
	if err != nil {
		return fmt.Errorf("Listing users: %s", err)
	}

	for _, r := range users.Users {
		n := *r.UserName

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete user %s?", n))
		if !proceed {
			continue
		}

		if err := o.accessKeys.Delete(n); err != nil {
			return fmt.Errorf("Deleting access keys for %s: %s", n, err)
		}

		if err := o.policies.Delete(n); err != nil {
			return fmt.Errorf("Deleting policies for %s: %s", n, err)
		}

		_, err = o.client.DeleteUser(&awsiam.DeleteUserInput{UserName: r.UserName})
		if err == nil {
			o.logger.Printf("SUCCESS deleting user %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting user %s: %s\n", n, err)
		}
	}

	return nil
}
