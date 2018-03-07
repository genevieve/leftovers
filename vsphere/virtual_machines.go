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

type VirtualMachines struct {
	vimClient  *govmomi.Client
	logger     logger
	datacenter string
}

func NewVirtualMachines(logger logger, vimClient *govmomi.Client, datacenter string) VirtualMachines {
	return VirtualMachines{
		logger:     logger,
		vimClient:  vimClient,
		datacenter: datacenter,
	}
}

func (v VirtualMachines) List(filter string) ([]common.Deletable, error) {
	dc, err := DatacenterFromID(v.vimClient, v.datacenter)
	if err != nil {
		return nil, err
	}

	finder := find.NewFinder(v.vimClient.Client, true)
	finder.SetDatacenter(dc)
	ctx := context.Background()

	folder, err := finder.Folder(ctx, filter)
	if err != nil {
		return nil, err
	}

	var deletable []common.Deletable

	children, err := folder.Children(ctx)
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
				vm := DeletableVM{
					vimClient: v.vimClient,
					vm:        vmGrandchild,
				}
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

type DeletableVM struct {
	vimClient *govmomi.Client
	vm        *object.VirtualMachine
}

func (v DeletableVM) Delete() error {
	tctx, tcancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer tcancel()

	powerState, err := v.vm.PowerState(context.Background())
	if err != nil {
		return fmt.Errorf("error getting power state: %s", powerState)
	}
	if powerState == "poweredOn" || powerState == "suspended" {
		powerOff, err := v.vm.PowerOff(context.Background())
		if err != nil {
			return fmt.Errorf("error shutting down virtual machine: %s", err)
		}
		err = powerOff.Wait(tctx)
		if err != nil {
			return fmt.Errorf("error waiting for machine to shut down: %s", err)
		}
	}

	destroy, err := v.vm.Destroy(context.Background())
	if err != nil {
		return fmt.Errorf("error destroying virtual machine: %s", err)
	}
	err = destroy.Wait(tctx)
	if err != nil {
		return fmt.Errorf("error waiting for machine to destroy: %s", err)
	}
	return nil
}

func (v DeletableVM) Name() string {
	name, _ := v.vm.Common.ObjectName(context.Background())
	return name
}
