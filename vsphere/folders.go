package vsphere

import (
	"context"
	"fmt"
	"time"

	"github.com/genevieve/leftovers/aws/common"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
)

type Folders struct {
	vimClient  *govmomi.Client
	logger     logger
	datacenter string
}

func NewFolders(logger logger, vimClient *govmomi.Client, datacenter string) Folders {
	return Folders{
		logger:     logger,
		vimClient:  vimClient,
		datacenter: datacenter,
	}
}

func (v Folders) List(filter string) ([]common.Deletable, error) {
	dc, err := DatacenterFromID(v.vimClient, v.datacenter)
	if err != nil {
		return nil, err
	}

	finder := find.NewFinder(v.vimClient.Client, true)
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
			name, _ := childFolder.Common.ObjectName(context.Background())
			folder := DeletableFolder{
				vimClient: v.vimClient,
				folder:    childFolder,
				name:      name,
			}
			proceed := v.logger.Prompt(fmt.Sprintf("Are you sure you want to delete empty folder %s?", name))
			if !proceed {
				continue
			}
			deletable = append(deletable, folder)
		}
	}

	return deletable, nil
}

type DeletableFolder struct {
	vimClient *govmomi.Client
	folder    *object.Folder
	name      string
}

func (f DeletableFolder) Delete() error {
	tctx, tcancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer tcancel()

	destroy, err := f.folder.Common.Destroy(tctx)
	if err != nil {
		return fmt.Errorf("error destroying folder: %s", err)
	}
	err = destroy.Wait(tctx)
	if err != nil {
		return fmt.Errorf("error waiting for folder to destroy: %s", err)
	}

	return nil
}

func (f DeletableFolder) Name() string {
	return f.name
}
