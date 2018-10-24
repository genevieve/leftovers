package groupingobjects

import (
	"context"
	"fmt"
)

type IPSet struct {
	client groupingObjectsAPI
	ctx    context.Context
	id     string
	name   string
}

func NewIPSet(client groupingObjectsAPI, ctx context.Context, name, id string) IPSet {
	return IPSet{
		client: client,
		ctx:    ctx,
		name:   name,
		id:     id,
	}
}

func (i IPSet) Delete() error {
	_, err := i.client.DeleteIPSet(i.ctx, i.id, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}
	return nil
}

func (i IPSet) Name() string {
	return i.name
}

func (i IPSet) Type() string {
	return "IP Set"
}
