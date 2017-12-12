package ec2

import (
	"fmt"

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

func (a Tags) Delete() error {
	tags, err := a.client.DescribeTags(&awsec2.DescribeTagsInput{})
	if err != nil {
		return fmt.Errorf("Describing tags: %s", err)
	}

	for _, t := range tags.Tags {
		n := *t.Value
		proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to delete tag %s?", n))
		if !proceed {
			continue
		}

		_, err := a.client.DeleteTags(&awsec2.DeleteTagsInput{
			Tags:      []*awsec2.Tag{{Key: t.Key}},
			Resources: []*string{t.ResourceId},
		})
		if err == nil {
			a.logger.Printf("SUCCESS deleting tag %s\n", n)
		} else {
			a.logger.Printf("ERROR deleting tag %s: %s\n", n, err)
		}
	}

	return nil
}
