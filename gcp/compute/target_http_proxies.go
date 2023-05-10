package compute

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface targetHttpProxiesClient --output fakes/target_http_proxies_client.go
type targetHttpProxiesClient interface {
	ListTargetHttpProxies() (*gcpcompute.TargetHttpProxyList, error)
	DeleteTargetHttpProxy(targetHttpProxy string) error
}

type TargetHttpProxies struct {
	client targetHttpProxiesClient
	logger logger
}

func NewTargetHttpProxies(client targetHttpProxiesClient, logger logger) TargetHttpProxies {
	return TargetHttpProxies{
		client: client,
		logger: logger,
	}
}

func (t TargetHttpProxies) List(filter string, regex bool) ([]common.Deletable, error) {
	t.logger.Debugln("Listing Target Http Proxies...")
	targetHttpProxies, err := t.client.ListTargetHttpProxies()
	if err != nil {
		return nil, fmt.Errorf("List Target Http Proxies: %s", err)
	}

	var resources []common.Deletable
	for _, targetHttpProxy := range targetHttpProxies.Items {
		resource := NewTargetHttpProxy(t.client, targetHttpProxy.Name)

		if !common.ResourceMatches(resource.Name(), filter, regex) {
			continue
		}

		proceed := t.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (t TargetHttpProxies) Type() string {
	return "target-http-proxy"
}
