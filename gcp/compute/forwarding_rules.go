package compute

import (
	"fmt"

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

func (o ForwardingRules) Delete() error {
	var rules []*gcpcompute.ForwardingRule
	for _, region := range o.regions {
		l, err := o.client.ListForwardingRules(region)
		if err != nil {
			return fmt.Errorf("Listing forwarding rules for region %s: %s", region, err)
		}
		rules = append(rules, l.Items...)
	}

	for _, r := range rules {
		n := r.Name

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete forwarding rule %s?", n))
		if !proceed {
			continue
		}

		regionName := o.regions[r.Region]
		if err := o.client.DeleteForwardingRule(regionName, n); err != nil {
			o.logger.Printf("ERROR deleting forwarding rule %s: %s\n", n, err)
		} else {
			o.logger.Printf("SUCCESS deleting forwarding rule %s\n", n)
		}
	}

	return nil
}
