package compute

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

type routersClient interface {
	ListRouters() ([]*gcpcompute.Router, error)
	DeleteRouter(router string) error
}

type Routers struct {
	routersClient routersClient
	logger        logger
}

func NewRouters(routersClient routersClient, logger logger) Routers {
	return Routers{
		routersClient: routersClient,
		logger:        logger,
	}
}

func (r Routers) List(filter string) ([]common.Deletable, error) {
	list, err := r.routersClient.ListRouters()
	if err != nil {
		return []common.Deletable{}, fmt.Errorf("List Routers: %s", err)
	}

	var resources []common.Deletable
	for _, router := range list {
		if !strings.Contains(router.Name, filter) {
			continue
		}
		r.logger.PromptWithDetails("cloudRouter", router.Name)
		resources = append(resources, NewRouter(r.routersClient))
	}

	return resources, nil
}
