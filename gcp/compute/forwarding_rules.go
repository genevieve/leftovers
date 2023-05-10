package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface forwardingRulesClient --output fakes/forwarding_rules_client.go
type forwardingRulesClient interface {
	ListForwardingRules(region string) ([]*gcpcompute.ForwardingRule, error)
	DeleteForwardingRule(region, rule string) error
}

type ForwardingRules struct {
	client  forwardingRulesClient
	logger  logger
	regions map[string]string
}

func NewForwardingRules(client forwardingRulesClient, logger logger, regions map[string]string) ForwardingRules {
	return ForwardingRules{
		client:  client,
		logger:  logger,
		regions: regions,
	}
}

func (f ForwardingRules) List(filter string, regex bool) ([]common.Deletable, error) {
	rules := []*gcpcompute.ForwardingRule{}
	for _, region := range f.regions {
		f.logger.Debugf("Listing Forwarding Rules for Region %s...\n", region)
		l, err := f.client.ListForwardingRules(region)
		if err != nil {
			return nil, fmt.Errorf("List Forwarding Rules for region %s: %s", region, err)
		}

		rules = append(rules, l...)
	}

	var resources []common.Deletable
	for _, rule := range rules {
		resource := NewForwardingRule(f.client, rule.Name, f.regions[rule.Region])

		if !common.ResourceMatches(rule.Name, filter, regex) {
			continue
		}

		proceed := f.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (f ForwardingRules) Type() string {
	return "forwarding-rule"
}
