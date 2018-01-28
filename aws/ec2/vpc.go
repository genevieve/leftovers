package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type Vpc struct {
	id     *string
	name   string
	client vpcsClient
}

func NewVpc(client vpcsClient, id *string, tags []*awsec2.Tag) Vpc {
	name := *id

	var extra []string
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		name = fmt.Sprintf("%s (%s)", *id, strings.Join(extra, ","))
	}

	return Vpc{
		client: client,
		id:     id,
		name:   name,
	}
}
