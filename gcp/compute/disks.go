package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type disksClient interface {
	ListDisks(zone string) (*gcpcompute.DiskList, error)
	DeleteDisk(zone, disk string) error
}

type Disks struct {
	client disksClient
	logger logger
	zones  map[string]string
}

func NewDisks(client disksClient, logger logger, zones map[string]string) Disks {
	return Disks{
		client: client,
		logger: logger,
		zones:  zones,
	}
}

func (d Disks) List(filter string) (map[string]string, error) {
	disks := []*gcpcompute.Disk{}
	delete := map[string]string{}

	for _, zone := range d.zones {
		l, err := d.client.ListDisks(zone)
		if err != nil {
			return delete, fmt.Errorf("Listing disks for zone %s: %s", zone, err)
		}
		disks = append(disks, l.Items...)
	}

	for _, disk := range disks {
		if !strings.Contains(disk.Name, filter) {
			continue
		}

		if len(disk.Users) > 0 {
			continue
		}

		proceed := d.logger.Prompt(fmt.Sprintf("Are you sure you want to delete disk %s?", disk.Name))
		if !proceed {
			continue
		}

		delete[disk.Name] = d.zones[disk.Zone]
	}

	return delete, nil
}

func (d Disks) Delete(disks map[string]string) {
	for name, zone := range disks {
		err := d.client.DeleteDisk(zone, name)

		if err != nil {
			d.logger.Printf("ERROR deleting disk %s: %s\n", name, err)
		} else {
			d.logger.Printf("SUCCESS deleting disk %s\n", name)
		}
	}
}
