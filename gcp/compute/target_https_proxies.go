package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type targetHttpsProxiesClient interface {
	ListTargetHttpsProxies() (*gcpcompute.TargetHttpsProxyList, error)
	DeleteTargetHttpsProxy(targetHttpsProxy string) error
}

type TargetHttpsProxies struct {
	client targetHttpsProxiesClient
	logger logger
}

func NewTargetHttpsProxies(client targetHttpsProxiesClient, logger logger) TargetHttpsProxies {
	return TargetHttpsProxies{
		client: client,
		logger: logger,
	}
}

func (t TargetHttpsProxies) List(filter string) (map[string]string, error) {
	targetHttpsProxies, err := t.client.ListTargetHttpsProxies()
	if err != nil {
		return nil, fmt.Errorf("Listing target https proxies: %s", err)
	}

	delete := map[string]string{}
	for _, targetHttpsProxy := range targetHttpsProxies.Items {
		if !strings.Contains(targetHttpsProxy.Name, filter) {
			continue
		}

		proceed := t.logger.Prompt(fmt.Sprintf("Are you sure you want to delete target https proxy %s?", targetHttpsProxy.Name))
		if !proceed {
			continue
		}

		delete[targetHttpsProxy.Name] = ""
	}

	return delete, nil
}

func (t TargetHttpsProxies) Delete(targetHttpsProxies map[string]string) {
	var resources []TargetHttpsProxy
	for name, _ := range targetHttpsProxies {
		resources = append(resources, NewTargetHttpsProxy(t.client, name))
	}

	t.delete(resources)
}

func (t TargetHttpsProxies) delete(resources []TargetHttpsProxy) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			t.logger.Printf("%s\n", err)
		} else {
			t.logger.Printf("SUCCESS deleting target https proxy %s\n", resource.name)
		}
	}
}
