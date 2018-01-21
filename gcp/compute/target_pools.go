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

func (a TargetPools) Delete(filter string) error {
	var pools []*gcpcompute.TargetPool
	for _, region := range a.regions {
		l, err := a.client.ListTargetPools(region)
		if err != nil {
			return fmt.Errorf("Listing target pools for region %s: %s", region, err)
		}
		pools = append(pools, l.Items...)
	}

	for _, p := range pools {
		n := p.Name

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to delete target pool %s?", n))
		if !proceed {
			continue
		}

		regionName := a.regions[p.Region]
		if err := a.client.DeleteTargetPool(regionName, n); err != nil {
			a.logger.Printf("ERROR deleting target pool %s: %s\n", n, err)
		} else {
			a.logger.Printf("SUCCESS deleting target pool %s\n", n)
		}
	}

	return nil
}
