package vsphere

import (
	"context"
	"fmt"
	"time"

	"github.com/vmware/govmomi/object"
)

type Folder struct {
	folder *object.Folder
	name   string
}

func NewFolder(folder *object.Folder, name string) Folder {
	return Folder{
		folder: folder,
		name:   name,
	}
}

func (f Folder) Delete() error {
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

func (f Folder) Name() string {
	return f.name
}
