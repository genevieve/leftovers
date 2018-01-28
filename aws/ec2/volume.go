package ec2

type Volume struct {
	id     *string
	client volumesClient
}

func NewVolume(client volumesClient, id *string) Volume {
	return Volume{
		client: client,
		id:     id,
	}
}
