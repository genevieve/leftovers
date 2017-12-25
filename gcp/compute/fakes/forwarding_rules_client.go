package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type ForwardingRulesClient struct {
	ListForwardingRulesCall struct {
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			Output *gcpcompute.ForwardingRuleList
			Error  error
		}
	}

	DeleteForwardingRuleCall struct {
		CallCount int
		Receives  struct {
			Region         string
			ForwardingRule string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *ForwardingRulesClient) ListForwardingRules(region string) (*gcpcompute.ForwardingRuleList, error) {
	n.ListForwardingRulesCall.CallCount++
	n.ListForwardingRulesCall.Receives.Region = region

	return n.ListForwardingRulesCall.Returns.Output, n.ListForwardingRulesCall.Returns.Error
}

func (n *ForwardingRulesClient) DeleteForwardingRule(region, forwardingRule string) error {
	n.DeleteForwardingRuleCall.CallCount++
	n.DeleteForwardingRuleCall.Receives.ForwardingRule = forwardingRule
	n.DeleteForwardingRuleCall.Receives.Region = region

	return n.DeleteForwardingRuleCall.Returns.Error
}
