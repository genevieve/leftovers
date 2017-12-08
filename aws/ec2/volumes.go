package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type Volumes struct {
	client ec2Client
	logger logger
}

func NewVolumes(client ec2Client, logger logger) Volumes {
	return Volumes{
		client: client,
		logger: logger,
	}
}

func (o Volumes) Delete() error {
	volumes, err := o.client.DescribeVolumes(&awsec2.DescribeVolumesInput{})
	if err != nil {
		return fmt.Errorf("Describing volumes: %s", err)
	}

	for _, v := range volumes.Volumes {
		state := *v.State
		if state != "available" {
			continue
		}

		n := *v.VolumeId
		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete volume %s?", n))
		if !proceed {
			continue
		}

		_, err := o.client.DeleteVolume(&awsec2.DeleteVolumeInput{VolumeId: aws.String(n)})
		if err == nil {
			o.logger.Printf("SUCCESS deleting volume %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting volume %s: %s\n", n, err)
		}
	}

	return nil
}
