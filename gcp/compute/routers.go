package compute

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface routersClient --output fakes/routers_client.go
type routersClient interface {
	ListRouters(region string) ([]*gcpcompute.Router, error)
	DeleteRouter(region, router string) error
}

type Routers struct {
	routersClient routersClient
	logger        logger
	regions       map[string]string
}

func NewRouters(routersClient routersClient, logger logger, regions map[string]string) Routers {
	return Routers{
		routersClient: routersClient,
		logger:        logger,
		regions:       regions,
	}
}

func (r Routers) List(filter string) ([]common.Deletable, error) {
	routers := []*gcpcompute.Router{}
	for _, region := range r.regions {
		r.logger.Debugf("Listing Routers for Region %s...\n", region)
		l, err := r.routersClient.ListRouters(region)
		if err != nil {
			return []common.Deletable{}, fmt.Errorf("List Routers for region %s: %s", region, err)
		}

		routers = append(routers, l...)
	}

	var resources []common.Deletable
	for _, router := range routers {
		resource := NewRouter(r.routersClient, router.Name, r.regions[router.Region])

		if !strings.Contains(resource.Name(), filter) {
			continue
		}

		proceed := r.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (r Routers) Type() string {
	return "cloud-router"
}
