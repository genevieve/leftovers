package compute

import (
	"fmt"

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

func (i Disks) Delete(filter string) error {
	var disks []*gcpcompute.Disk
	for _, zone := range i.zones {
		l, err := i.client.ListDisks(zone)
		if err != nil {
			return fmt.Errorf("Listing disks for zone %s: %s", zone, err)
		}
		disks = append(disks, l.Items...)
	}

	for _, d := range disks {
		if len(d.Users) > 0 {
			continue
		}

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete disk %s?", d.Name))
		if !proceed {
			continue
		}

		zoneName := i.zones[d.Zone]
		if err := i.client.DeleteDisk(zoneName, d.Name); err != nil {
			i.logger.Printf("ERROR deleting disk %s: %s\n", d.Name, err)
		} else {
			i.logger.Printf("SUCCESS deleting disk %s\n", d.Name)
		}
	}

	return nil
}
