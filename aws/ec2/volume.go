package ec2

import awsec2 "github.com/aws/aws-sdk-go/service/ec2"

type Volume struct {
	client     volumesClient
	id         *string
	identifier string
}

func NewVolume(client volumesClient, id *string) Volume {
	return Volume{
		client:     client,
		id:         id,
		identifier: *id,
	}
}

func (v Volume) Delete() error {
	_, err := v.client.DeleteVolume(&awsec2.DeleteVolumeInput{
		VolumeId: v.id,
	})
	return err
}
