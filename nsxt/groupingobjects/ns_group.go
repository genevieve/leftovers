package groupingobjects

import (
	"context"
	"fmt"
)

type NSGroup struct {
	client groupingObjectsAPI
	ctx    context.Context
	id     string
	name   string
}

func NewNSGroup(client groupingObjectsAPI, ctx context.Context, name, id string) NSGroup {
	return NSGroup{
		client: client,
		ctx:    ctx,
		name:   name,
		id:     id,
	}
}

func (n NSGroup) Delete() error {
	_, err := n.client.DeleteNSGroup(n.ctx, n.id, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}
	return nil
}

func (n NSGroup) Name() string {
	return n.name
}

func (n NSGroup) Type() string {
	return "NS Group"
}
