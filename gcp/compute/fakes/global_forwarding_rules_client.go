package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type GlobalForwardingRulesClient struct {
	ListGlobalForwardingRulesCall struct {
		CallCount int
		Returns   struct {
			Output []*gcpcompute.ForwardingRule
			Error  error
		}
	}

	DeleteGlobalForwardingRuleCall struct {
		CallCount int
		Receives  struct {
			GlobalForwardingRule string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *GlobalForwardingRulesClient) ListGlobalForwardingRules() ([]*gcpcompute.ForwardingRule, error) {
	n.ListGlobalForwardingRulesCall.CallCount++

	return n.ListGlobalForwardingRulesCall.Returns.Output, n.ListGlobalForwardingRulesCall.Returns.Error
}

func (n *GlobalForwardingRulesClient) DeleteGlobalForwardingRule(globalForwardingRule string) error {
	n.DeleteGlobalForwardingRuleCall.CallCount++
	n.DeleteGlobalForwardingRuleCall.Receives.GlobalForwardingRule = globalForwardingRule

	return n.DeleteGlobalForwardingRuleCall.Returns.Error
}
