package artifacts

import (
	"fmt"

	"github.com/genevieve/leftovers/gcp/common"

	gcpartifact "google.golang.org/api/artifactregistry/v1"
)

type operationWaiter struct {
	op      *gcpartifact.Operation
	service *gcpartifact.ProjectsLocationsService
	project string
	logger  logger
}

func NewOperationWaiter(op *gcpartifact.Operation, service *gcpartifact.Service, project string, logger logger) operationWaiter {
	return operationWaiter{
		op:      op,
		service: service.Projects.Locations,
		project: project,
		logger:  logger,
	}
}

func (w *operationWaiter) Wait() error {
	state := common.NewState(w.logger, w.refreshFunc())

	raw, err := state.Wait()
	if err != nil {
		return fmt.Errorf("Waiting for operation to complete: %s", err)
	}

	result, ok := raw.(*gcpartifact.Operation)
	if ok && result.Error != nil {
		return fmt.Errorf("operation error: %s", result.Error.Message)
	}

	return nil
}

func (c *operationWaiter) refreshFunc() common.StateRefreshFunc {
	return func() (interface{}, string, error) {
		op, err := c.service.Operations.Get(c.op.Name).Do()

		if err != nil {
			return nil, "", fmt.Errorf("Refreshing operation request: %s", err)
		}

		// This service has no `Status` field, so we need to fake as
		// it's not possible to distinguish pending from running, so
		// I'm selecting a reasonable default.
		status := "RUNNING"
		if op.Done {
			status = "DONE"
		}

		return op, status, nil
	}
}
