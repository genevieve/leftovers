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
	tags, err := a.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, tag := range tags {
		delete[*tag.key] = *tag.resourceId
	}

	return delete, nil
}

func (a Tags) list(filter string) ([]Tag, error) {
	output, err := a.client.DescribeTags(&awsec2.DescribeTagsInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing tags: %s", err)
	}

	var resources []Tag
	for _, t := range output.Tags {
		resource := NewTag(a.client, t.Key, t.Value, t.ResourceId)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		//TODO: Prompt with key:value
		proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to delete tag %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (t Tags) Delete(tags map[string]string) error {
	for key, resourceId := range tags {
		_, err := t.client.DeleteTags(&awsec2.DeleteTagsInput{
			//TODO: Delete with key:value
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
