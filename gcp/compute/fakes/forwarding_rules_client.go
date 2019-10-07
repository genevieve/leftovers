package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type ForwardingRulesClient struct {
	DeleteForwardingRuleCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
			Rule   string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListForwardingRulesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			ForwardingRuleSlice []*gcpcompute.ForwardingRule
			Error               error
		}
		Stub func(string) ([]*gcpcompute.ForwardingRule, error)
	}
}

func (f *ForwardingRulesClient) DeleteForwardingRule(param1 string, param2 string) error {
	f.DeleteForwardingRuleCall.Lock()
	defer f.DeleteForwardingRuleCall.Unlock()
	f.DeleteForwardingRuleCall.CallCount++
	f.DeleteForwardingRuleCall.Receives.Region = param1
	f.DeleteForwardingRuleCall.Receives.Rule = param2
	if f.DeleteForwardingRuleCall.Stub != nil {
		return f.DeleteForwardingRuleCall.Stub(param1, param2)
	}
	return f.DeleteForwardingRuleCall.Returns.Error
}
func (f *ForwardingRulesClient) ListForwardingRules(param1 string) ([]*gcpcompute.ForwardingRule, error) {
	f.ListForwardingRulesCall.Lock()
	defer f.ListForwardingRulesCall.Unlock()
	f.ListForwardingRulesCall.CallCount++
	f.ListForwardingRulesCall.Receives.Region = param1
	if f.ListForwardingRulesCall.Stub != nil {
		return f.ListForwardingRulesCall.Stub(param1)
	}
	return f.ListForwardingRulesCall.Returns.ForwardingRuleSlice, f.ListForwardingRulesCall.Returns.Error
}
