package iam

import (
	"fmt"
	awsiam "github.com/aws/aws-sdk-go/service/iam"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface usersClient --output fakes/users_client.go
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

func (u Users) List(filter string, regex bool) ([]common.Deletable, error) {
	users, err := u.client.ListUsers(&awsiam.ListUsersInput{})
	if err != nil {
		return nil, fmt.Errorf("List IAM Users: %s", err)
	}

	var resources []common.Deletable
	for _, r := range users.Users {
		r := NewUser(u.client, u.policies, u.accessKeys, r.UserName)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := u.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (u Users) Type() string {
	return "iam-user"
}
