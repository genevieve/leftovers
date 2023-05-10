package ec2

import (
	"fmt"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	awssts "github.com/aws/aws-sdk-go/service/sts"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface imagesClient --output fakes/images_client.go
type imagesClient interface {
	DescribeImages(*awsec2.DescribeImagesInput) (*awsec2.DescribeImagesOutput, error)
	DeregisterImage(*awsec2.DeregisterImageInput) (*awsec2.DeregisterImageOutput, error)
}

//go:generate faux --interface stsClient --output fakes/sts_client.go
type stsClient interface {
	GetCallerIdentity(*awssts.GetCallerIdentityInput) (*awssts.GetCallerIdentityOutput, error)
}

type Images struct {
	client       imagesClient
	stsClient    stsClient
	logger       logger
	resourceTags resourceTags
}

func NewImages(client imagesClient, stsClient stsClient, logger logger, resourceTags resourceTags) Images {
	return Images{
		client:       client,
		stsClient:    stsClient,
		logger:       logger,
		resourceTags: resourceTags,
	}
}

func (i Images) List(filter string, regex bool) ([]common.Deletable, error) {
	caller, err := i.stsClient.GetCallerIdentity(&awssts.GetCallerIdentityInput{})
	if err != nil {
		return nil, fmt.Errorf("Get caller identity: %s", err)
	}

	images, err := i.client.DescribeImages(&awsec2.DescribeImagesInput{
		Owners: []*string{caller.Account},
	})
	if err != nil {
		return nil, fmt.Errorf("Describing EC2 Images: %s", err)
	}

	var resources []common.Deletable
	for _, image := range images.Images {
		r := NewImage(i.client, image.ImageId, i.resourceTags)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := i.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (i Images) Type() string {
	return "ec2-image"
}
