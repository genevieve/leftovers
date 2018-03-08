package vsphere

import (
	"context"
	"fmt"

	"github.com/genevieve/leftovers/aws/common"
	"github.com/vmware/govmomi/object"
)

type VirtualMachines struct {
	client client
	logger logger
}

func NewVirtualMachines(client client, logger logger) VirtualMachines {
	return VirtualMachines{
		client: client,
		logger: logger,
	}
}

func (v VirtualMachines) List(filter string) ([]common.Deletable, error) {
	folder, err := v.client.GetRootFolder(filter)
	if err != nil {
		return nil, fmt.Errorf("Getting root folder: %s", err)
	}

	var deletable []common.Deletable

	ctx := context.Background()

	children, err := folder.Children(ctx)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		childFolder, ok := child.(*object.Folder)
		if !ok {
			continue
		}

		grandchildren, err := childFolder.Children(ctx)
		if err != nil {
			return nil, err
		}

		for _, grandchild := range grandchildren {
			if vmGrandchild, ok := grandchild.(*object.VirtualMachine); ok {
				vm := NewVirtualMachine(vmGrandchild)

				proceed := v.logger.Prompt(fmt.Sprintf("Are you sure you want to delete virtual machine %s?", vm.Name()))
				if !proceed {
					continue
				}

				deletable = append(deletable, vm)
			}
		}
	}

	return deletable, nil
}
