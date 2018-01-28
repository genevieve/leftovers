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

func (v Volumes) List(filter string) (map[string]string, error) {
	volumes, err := v.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, volume := range volumes {
		delete[*volume.id] = ""
	}

	return delete, nil
}

func (v Volumes) list(filter string) ([]Volume, error) {
	output, err := v.client.DescribeVolumes(&awsec2.DescribeVolumesInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing volumes: %s", err)
	}

	var volumes []Volume
	for _, volume := range output.Volumes {
		if *volume.State != "available" {
			continue
		}

		proceed := v.logger.Prompt(fmt.Sprintf("Are you sure you want to delete volume %s?", *volume.VolumeId))
		if !proceed {
			continue
		}

		volumes = append(volumes, NewVolume(v.client, volume.VolumeId))
	}

	return volumes, nil
}

func (v Volumes) Delete(volumes map[string]string) error {
	var resources []Volume
	for id, _ := range volumes {
		resources = append(resources, NewVolume(v.client, &id))
	}

	return v.cleanup(resources)
}

func (v Volumes) cleanup(resources []Volume) error {
	for _, resource := range resources {
		err := resource.Delete()

		if err == nil {
			v.logger.Printf("SUCCESS deleting volume %s\n", resource.identifier)
		} else {
			v.logger.Printf("ERROR deleting volume %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}
