package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type targetPoolsClient interface {
	ListTargetPools(region string) (*gcpcompute.TargetPoolList, error)
	DeleteTargetPool(region string, targetPool string) error
}

type TargetPools struct {
	client  targetPoolsClient
	logger  logger
	regions map[string]string
}

func NewTargetPools(client targetPoolsClient, logger logger, regions map[string]string) TargetPools {
	return TargetPools{
		client:  client,
		logger:  logger,
		regions: regions,
	}
}

func (t TargetPools) List(filter string) (map[string]string, error) {
	pools := []*gcpcompute.TargetPool{}
	delete := map[string]string{}

	for _, region := range t.regions {
		l, err := t.client.ListTargetPools(region)
		if err != nil {
			return delete, fmt.Errorf("Listing target pools for region %s: %s", region, err)
		}

		pools = append(pools, l.Items...)
	}

	for _, pool := range pools {
		if !strings.Contains(pool.Name, filter) {
			continue
		}

		proceed := t.logger.Prompt(fmt.Sprintf("Are you sure you want to delete target pool %s?", pool.Name))
		if !proceed {
			continue
		}

		delete[pool.Name] = t.regions[pool.Region]
	}

	return delete, nil
}

func (t TargetPools) Delete(pools map[string]string) {
	for name, region := range pools {
		err := t.client.DeleteTargetPool(region, name)

		if err != nil {
			t.logger.Printf("ERROR deleting target pool %s: %s\n", name, err)
		} else {
			t.logger.Printf("SUCCESS deleting target pool %s\n", name)
		}
	}
}
