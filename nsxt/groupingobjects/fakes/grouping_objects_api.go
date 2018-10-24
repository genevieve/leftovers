package fakes

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type GroupingObjectsAPI struct {
	ListIPSetsCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			IPSetListResult manager.IpSetListResult
			Response        *http.Response
			Error           error
		}
	}
	DeleteIPSetCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			ID                string
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
	}
}

func (g *GroupingObjectsAPI) ListIPSets(ctx context.Context, localVarOptionals map[string]interface{}) (manager.IpSetListResult, *http.Response, error) {
	g.ListIPSetsCall.CallCount++

	g.ListIPSetsCall.Receives.Context = ctx
	g.ListIPSetsCall.Receives.LocalVarOptionals = localVarOptionals

	return g.ListIPSetsCall.Returns.IPSetListResult, g.ListIPSetsCall.Returns.Response, g.ListIPSetsCall.Returns.Error
}

func (g *GroupingObjectsAPI) DeleteIPSet(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	g.DeleteIPSetCall.CallCount++

	g.DeleteIPSetCall.Receives.Context = ctx
	g.DeleteIPSetCall.Receives.ID = id
	g.DeleteIPSetCall.Receives.LocalVarOptionals = localVarOptionals

	return g.DeleteIPSetCall.Returns.Response, g.DeleteIPSetCall.Returns.Error
}
