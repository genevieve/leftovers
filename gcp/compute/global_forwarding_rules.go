package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface globalForwardingRulesClient --output fakes/global_forwarding_rules_client.go
type globalForwardingRulesClient interface {
	ListGlobalForwardingRules() ([]*gcpcompute.ForwardingRule, error)
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

func (g GlobalForwardingRules) List(filter string, regex bool) ([]common.Deletable, error) {
	g.logger.Debugln("Listing Global Forwarding Rules...")
	rules, err := g.client.ListGlobalForwardingRules()
	if err != nil {
		return nil, fmt.Errorf("List Global Forwarding Rules: %s", err)
	}

	var resources []common.Deletable
	for _, rule := range rules {
		resource := NewGlobalForwardingRule(g.client, rule.Name)

		if !common.MatchRegex(rule.Name, filter, regex) {
			continue
		}

		proceed := g.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (g GlobalForwardingRules) Type() string {
	return "global-forwarding-rule"
}
