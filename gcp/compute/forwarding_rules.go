package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type forwardingRulesClient interface {
	ListForwardingRules(region string) (*gcpcompute.ForwardingRuleList, error)
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

func (f ForwardingRules) List(filter string) (map[string]string, error) {
	rules := []*gcpcompute.ForwardingRule{}
	for _, region := range f.regions {
		l, err := f.client.ListForwardingRules(region)
		if err != nil {
			return nil, fmt.Errorf("Listing forwarding rules for region %s: %s", region, err)
		}

		rules = append(rules, l.Items...)
	}

	delete := map[string]string{}
	for _, rule := range rules {
		if !strings.Contains(rule.Name, filter) {
			continue
		}

		proceed := f.logger.Prompt(fmt.Sprintf("Are you sure you want to delete forwarding rule %s?", rule.Name))
		if !proceed {
			continue
		}

		delete[rule.Name] = f.regions[rule.Region]
	}

	return delete, nil
}

func (f ForwardingRules) Delete(forwardingRules map[string]string) {
	var resources []ForwardingRule
	for name, region := range forwardingRules {
		resources = append(resources, NewForwardingRule(f.client, name, region))
	}

	f.delete(resources)
}

func (f ForwardingRules) delete(resources []ForwardingRule) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			f.logger.Printf("%s\n", err)
		} else {
			f.logger.Printf("SUCCESS deleting forwarding rule %s\n", resource.name)
		}
	}
}
