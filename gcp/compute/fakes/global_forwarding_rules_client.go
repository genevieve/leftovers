package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type GlobalForwardingRulesClient struct {
	DeleteGlobalForwardingRuleCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Rule string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListGlobalForwardingRulesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			ForwardingRuleSlice []*gcpcompute.ForwardingRule
			Error               error
		}
		Stub func() ([]*gcpcompute.ForwardingRule, error)
	}
}

func (f *GlobalForwardingRulesClient) DeleteGlobalForwardingRule(param1 string) error {
	f.DeleteGlobalForwardingRuleCall.Lock()
	defer f.DeleteGlobalForwardingRuleCall.Unlock()
	f.DeleteGlobalForwardingRuleCall.CallCount++
	f.DeleteGlobalForwardingRuleCall.Receives.Rule = param1
	if f.DeleteGlobalForwardingRuleCall.Stub != nil {
		return f.DeleteGlobalForwardingRuleCall.Stub(param1)
	}
	return f.DeleteGlobalForwardingRuleCall.Returns.Error
}
func (f *GlobalForwardingRulesClient) ListGlobalForwardingRules() ([]*gcpcompute.ForwardingRule, error) {
	f.ListGlobalForwardingRulesCall.Lock()
	defer f.ListGlobalForwardingRulesCall.Unlock()
	f.ListGlobalForwardingRulesCall.CallCount++
	if f.ListGlobalForwardingRulesCall.Stub != nil {
		return f.ListGlobalForwardingRulesCall.Stub()
	}
	return f.ListGlobalForwardingRulesCall.Returns.ForwardingRuleSlice, f.ListGlobalForwardingRulesCall.Returns.Error
}
