package compute

import "fmt"

type FloatingIP struct {
	client floatingipsClient
	name   string
}

func NewFloatingIP(client floatingipsClient, name string) FloatingIP {
	return FloatingIP{
		client: client,
		name:   name,
	}
}

func (f FloatingIP) Delete() error {
	err := f.client.DeleteFloatingIP(f.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (f FloatingIP) Name() string {
	return f.name
}

func (FloatingIP) Type() string {
	return "Floating IP"
}
