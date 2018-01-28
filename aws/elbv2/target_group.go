package elbv2

type TargetGroup struct {
	client     targetGroupsClient
	name       *string
	arn        *string
	identifier string
}

func NewTargetGroup(client targetGroupsClient, name, arn *string) TargetGroup {
	return TargetGroup{
		client:     client,
		name:       name,
		arn:        arn,
		identifier: *name,
	}
}
