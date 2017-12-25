package compute

import (
	"strings"

	"github.com/hashicorp/terraform/helper/resource"
	compute "google.golang.org/api/compute/v1"
)

type operationWaiter struct {
	Op      *compute.Operation
	Service *compute.Service
	Project string
}

func (c *operationWaiter) refreshFunc() resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var op *compute.Operation
		var err error

		if c.Op.Zone != "" {
			zoneURLParts := strings.Split(c.Op.Zone, "/")
			zone := zoneURLParts[len(zoneURLParts)-1]
			op, err = c.Service.ZoneOperations.Get(c.Project, zone, c.Op.Name).Do()
		} else if c.Op.Region != "" {
			regionURLParts := strings.Split(c.Op.Region, "/")
			region := regionURLParts[len(regionURLParts)-1]
			op, err = c.Service.RegionOperations.Get(c.Project, region, c.Op.Name).Do()
		} else {
			op, err = c.Service.GlobalOperations.Get(c.Project, c.Op.Name).Do()
		}
		if err != nil {
			return nil, "", err
		}

		return op, op.Status, nil
	}
}
