package ec2

import (
	"fmt"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type volumesClient interface {
	DescribeVolumes(*awsec2.DescribeVolumesInput) (*awsec2.DescribeVolumesOutput, error)
	DeleteVolume(*awsec2.DeleteVolumeInput) (*awsec2.DeleteVolumeOutput, error)
}

type Volumes struct {
	client volumesClient
	logger logger
}

func NewVolumes(client volumesClient, logger logger) Volumes {
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

		_, err := o.client.DeleteVolume(&awsec2.DeleteVolumeInput{
			VolumeId: v.VolumeId,
		})
		if err == nil {
			o.logger.Printf("SUCCESS deleting volume %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting volume %s: %s\n", n, err)
		}
	}

	return nil
}
