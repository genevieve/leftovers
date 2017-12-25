package compute

import (
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	gcpcompute "google.golang.org/api/compute/v1"
)

type operationWaiter struct {
	Op      *gcpcompute.Operation
	Service *gcpcompute.Service
	Project string
}

func (c *operationWaiter) Wait() error {
	state := &resource.StateChangeConf{
		Delay:      10 * time.Second,
		Timeout:    10 * time.Minute,
		MinTimeout: 2 * time.Second,
		Pending:    []string{"PENDING", "RUNNING"},
		Target:     []string{"DONE"},
		Refresh:    c.refreshFunc(),
	}

	opRaw, err := state.WaitForState()
	if err != nil {
		return err
	}

	if resultOp := opRaw.(*gcpcompute.Operation); resultOp.Error != nil {
		return ComputeOperationError(*resultOp.Error)
	}

	return nil
}

func (c *operationWaiter) refreshFunc() resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var op *gcpcompute.Operation
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
