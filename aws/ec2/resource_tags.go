package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type resourceTags interface {
	Delete(filterName, filterValue string) error
}

type ResourceTags struct {
	client tagsClient
}

func NewResourceTags(client tagsClient) ResourceTags {
	return ResourceTags{
		client: client,
	}
}

func (r ResourceTags) Delete(filterName, filterValue string) error {
	output, err := r.client.DescribeTags(&awsec2.DescribeTagsInput{
		Filters: []*awsec2.Filter{{
			Name:   aws.String(filterName),
			Values: []*string{aws.String(filterValue)},
		}},
	})
	if err != nil {
		return fmt.Errorf("Describe: %s", err)
	}

	for _, t := range output.Tags {
		_, err := r.client.DeleteTags(&awsec2.DeleteTagsInput{
			Tags:      []*awsec2.Tag{{Key: t.Key, Value: t.Value}},
			Resources: []*string{t.ResourceId},
		})

		if err != nil {
			return fmt.Errorf("Delete %s:%s: %s", *t.Key, *t.Value, err)
		}
	}

	return nil
}
