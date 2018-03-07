package vsphere

import (
	"context"
	"fmt"

	"github.com/genevieve/leftovers/aws/common"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
)

type Folders struct {
	client     *govmomi.Client
	logger     logger
	datacenter string
}

func NewFolders(logger logger, client *govmomi.Client, datacenter string) Folders {
	return Folders{
		logger:     logger,
		client:     client,
		datacenter: datacenter,
	}
}

func (v Folders) List(filter string) ([]common.Deletable, error) {
	dc, err := DatacenterFromID(v.client, v.datacenter)
	if err != nil {
		return nil, err
	}

	finder := find.NewFinder(v.client.Client, true)
	finder.SetDatacenter(dc)
	ctx := context.Background()

	rootFolder, err := finder.Folder(ctx, filter)
	if err != nil {
		return nil, err
	}

	var deletable []common.Deletable

	children, err := rootFolder.Children(ctx)
	for _, child := range children {
		childFolder, ok := child.(*object.Folder)
		if !ok {
			continue
		}
		grandchildren, err := childFolder.Children(ctx)
		if err != nil {
			return nil, err
		}
		if len(grandchildren) == 0 {
			name, err := childFolder.Common.ObjectName(ctx)
			if err != nil {
				return nil, err
			}

			folder := NewFolder(childFolder, name)

			proceed := v.logger.Prompt(fmt.Sprintf("Are you sure you want to delete empty folder %s?", name))
			if !proceed {
				continue
			}

			deletable = append(deletable, folder)
		}
	}

	return deletable, nil
}
