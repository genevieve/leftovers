package gcp

import (
	"fmt"

	compute "google.golang.org/api/compute/v1"
)

type disksClient interface {
	ListDisks() (*compute.DiskList, error)
	DeleteDisk(zone, disk string) (*compute.Operation, error)
}

type Disks struct {
	client disksClient
	logger logger
	zones  []string
}

func NewDisks(client disksClient, logger logger, zones []string) Disks {
	return Disks{
		client: client,
		logger: logger,
		zones:  zones,
	}
}

func (i Disks) Delete() error {
	disks, err := i.client.ListDisks()
	if err != nil {
		return fmt.Errorf("Listing disks: %s", err)
	}

	for _, d := range disks.Items {
		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete disk %s?", d.Name))
		if !proceed {
			continue
		}

		_, err = i.client.DeleteDisk(d.Zone, d.Name)
		if err == nil {
			i.logger.Printf("SUCCESS deleting disk %s\n", d.Name)
		} else {
			i.logger.Printf("ERROR deleting disk %s: %s\n", d.Name, err)
		}
	}

	return nil
}
