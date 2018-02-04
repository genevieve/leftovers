package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type globalForwardingRulesClient interface {
	ListGlobalForwardingRules() (*gcpcompute.ForwardingRuleList, error)
	DeleteGlobalForwardingRule(rule string) error
}

type GlobalForwardingRules struct {
	client globalForwardingRulesClient
	logger logger
}

func NewGlobalForwardingRules(client globalForwardingRulesClient, logger logger) GlobalForwardingRules {
	return GlobalForwardingRules{
		client: client,
		logger: logger,
	}
}

func (g GlobalForwardingRules) List(filter string) (map[string]string, error) {
	rules, err := g.client.ListGlobalForwardingRules()
	if err != nil {
		return nil, fmt.Errorf("Listing global forwarding rules: %s", err)
	}

	delete := map[string]string{}
	for _, rule := range rules.Items {
		if !strings.Contains(rule.Name, filter) {
			continue
		}

		proceed := g.logger.Prompt(fmt.Sprintf("Are you sure you want to delete global forwarding rule %s?", rule.Name))
		if !proceed {
			continue
		}

		delete[rule.Name] = ""
	}

	return delete, nil
}

func (g GlobalForwardingRules) Delete(globalForwardingRules map[string]string) {
	var resources []GlobalForwardingRule
	for name, _ := range globalForwardingRules {
		resources = append(resources, NewGlobalForwardingRule(g.client, name))
	}

	g.delete(resources)
}

func (g GlobalForwardingRules) delete(resources []GlobalForwardingRule) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			g.logger.Printf("%s\n", err)
		} else {
			g.logger.Printf("SUCCESS deleting global forwarding rule %s\n", resource.name)
		}
	}
}
