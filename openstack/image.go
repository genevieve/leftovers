package openstack

import "fmt"

type Image struct {
	name               string
	id                 string
	imageServiceClient ImageServiceClient
}

func NewImage(name string, id string, imageServiceClient ImageServiceClient) Image {
	return Image{
		name:               name,
		id:                 id,
		imageServiceClient: imageServiceClient,
	}
}

func (image Image) Delete() error {
	return image.imageServiceClient.Delete(image.id)
}
func (image Image) Name() string {
	return fmt.Sprintf("%s %s", image.name, image.id)
}

func (image Image) Type() string {
	return "Image"
}
