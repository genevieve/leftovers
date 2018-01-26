package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type tagsClient interface {
	DescribeTags(*awsec2.DescribeTagsInput) (*awsec2.DescribeTagsOutput, error)
	DeleteTags(*awsec2.DeleteTagsInput) (*awsec2.DeleteTagsOutput, error)
}

type Tags struct {
	client tagsClient
	logger logger
}

func NewTags(client tagsClient, logger logger) Tags {
	return Tags{
		client: client,
		logger: logger,
	}
}

func (a Tags) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	tags, err := a.client.DescribeTags(&awsec2.DescribeTagsInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing tags: %s", err)
	}

	for _, t := range tags.Tags {
		n := *t.Value

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to delete tag %s?", n))
		if !proceed {
			continue
		}

		delete[*t.Key] = *t.ResourceId
	}

	return delete, nil
}

func (t Tags) Delete(tags map[string]string) error {
	for key, resourceId := range tags {
		_, err := t.client.DeleteTags(&awsec2.DeleteTagsInput{
			Tags:      []*awsec2.Tag{{Key: aws.String(key)}},
			Resources: []*string{aws.String(resourceId)},
		})

		if err == nil {
			t.logger.Printf("SUCCESS deleting tag %s\n", key)
		} else {
			t.logger.Printf("ERROR deleting tag %s: %s\n", key, err)
		}
	}

	return nil
}
