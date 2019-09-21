package openstack

import "fmt"

type Image struct {
	name   string
	id     string
	client imageServiceClient
}

func NewImage(name string, id string, client imageServiceClient) Image {
	return Image{
		name:   fmt.Sprintf("%s %s", name, id),
		id:     id,
		client: client,
	}
}

func (i Image) Delete() error {
	return i.client.Delete(i.id)
}

func (i Image) Name() string {
	return i.name
}

func (i Image) Type() string {
	return "Image"
}
