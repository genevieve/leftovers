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

	ListNSServicesCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			NSServiceListResult manager.NsServiceListResult
			Response            *http.Response
			Error               error
		}
	}
	DeleteNSServiceCall struct {
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

	ListNSGroupsCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			NSGroupListResult manager.NsGroupListResult
			Response          *http.Response
			Error             error
		}
	}
	DeleteNSGroupCall struct {
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

func (g *GroupingObjectsAPI) ListNSServices(ctx context.Context, localVarOptionals map[string]interface{}) (manager.NsServiceListResult, *http.Response, error) {
	g.ListNSServicesCall.CallCount++

	g.ListNSServicesCall.Receives.Context = ctx
	g.ListNSServicesCall.Receives.LocalVarOptionals = localVarOptionals

	return g.ListNSServicesCall.Returns.NSServiceListResult, g.ListNSServicesCall.Returns.Response, g.ListNSServicesCall.Returns.Error
}

func (g *GroupingObjectsAPI) DeleteNSService(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	g.DeleteNSServiceCall.CallCount++

	g.DeleteNSServiceCall.Receives.Context = ctx
	g.DeleteNSServiceCall.Receives.ID = id
	g.DeleteNSServiceCall.Receives.LocalVarOptionals = localVarOptionals

	return g.DeleteNSServiceCall.Returns.Response, g.DeleteNSServiceCall.Returns.Error
}

func (g *GroupingObjectsAPI) ListNSGroups(ctx context.Context, localVarOptionals map[string]interface{}) (manager.NsGroupListResult, *http.Response, error) {
	g.ListNSGroupsCall.CallCount++

	g.ListNSGroupsCall.Receives.Context = ctx
	g.ListNSGroupsCall.Receives.LocalVarOptionals = localVarOptionals

	return g.ListNSGroupsCall.Returns.NSGroupListResult, g.ListNSGroupsCall.Returns.Response, g.ListNSGroupsCall.Returns.Error
}

func (g *GroupingObjectsAPI) DeleteNSGroup(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	g.DeleteNSGroupCall.CallCount++

	g.DeleteNSGroupCall.Receives.Context = ctx
	g.DeleteNSGroupCall.Receives.ID = id
	g.DeleteNSGroupCall.Receives.LocalVarOptionals = localVarOptionals

	return g.DeleteNSGroupCall.Returns.Response, g.DeleteNSGroupCall.Returns.Error
}
