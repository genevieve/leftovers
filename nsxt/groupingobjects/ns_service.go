package groupingobjects

import (
	"context"
	"fmt"
)

type NSService struct {
	client groupingObjectsAPI
	ctx    context.Context
	id     string
	name   string
}

func NewNSService(client groupingObjectsAPI, ctx context.Context, name, id string) NSService {
	return NSService{
		client: client,
		ctx:    ctx,
		name:   name,
		id:     id,
	}
}

func (n NSService) Delete() error {
	_, err := n.client.DeleteNSService(n.ctx, n.id, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}
	return nil
}

func (n NSService) Name() string {
	return n.name
}

func (n NSService) Type() string {
	return "NS Service"
}
