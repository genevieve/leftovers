package iam

type Policy struct {
	client     policiesClient
	name       *string
	arn        *string
	identifier string
}

func NewPolicy(client policiesClient, name, arn *string) Policy {
	return Policy{
		client:     client,
		name:       name,
		arn:        arn,
		identifier: *name,
	}
}
