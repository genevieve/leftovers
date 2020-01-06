package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type Instance struct {
	client      instancesClient
	name        string
	clearerName string
	zone        string
	diskNames   []string
}

func NewInstance(client instancesClient, name, zone string, tags *gcpcompute.Tags, networkInterfaces []*gcpcompute.NetworkInterface, disks []*gcpcompute.AttachedDisk) Instance {
	clearerName := name

	extra := []string{}
	for _, ni := range networkInterfaces {
		network := client.GetNetworkName(ni.Network)
		if len(network) > 0 {
			extra = append(extra, network)
		}
	}

	if tags != nil && len(tags.Items) > 0 {
		for _, tag := range tags.Items {
			extra = append(extra, tag)
		}
	}

	if len(extra) > 0 {
		clearerName = fmt.Sprintf("%s (%s)", name, strings.Join(extra, ", "))
	}

	diskNames := []string{}
	for _, d := range disks {
		diskNames = append(diskNames, d.DeviceName)
	}

	return Instance{
		client:      client,
		name:        name,
		clearerName: clearerName,
		zone:        zone,
		diskNames:   diskNames,
	}
}

func (i Instance) Delete() error {
	for _, d := range i.diskNames {
		err := i.client.SetDiskAutoDelete(i.zone, i.name, d)
		if err != nil {
			return fmt.Errorf("Set Disk Auto Delete: %s", err)
		}
	}

	err := i.client.DeleteInstance(i.zone, i.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (i Instance) Name() string {
	return i.clearerName
}

func (i Instance) Type() string {
	return "Compute Instance"
}
