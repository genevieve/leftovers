package fakes

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type LogicalRoutingAndServicesAPI struct {
	ListLogicalRoutersCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			ListResult manager.LogicalRouterListResult
			Response   *http.Response
			Error      error
		}
	}

	DeleteLogicalRouterCall struct {
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

	ListLogicalRouterPortsCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			ListResult manager.LogicalRouterPortListResult
			Response   *http.Response
			Error      error
		}
	}

	DeleteLogicalRouterPortCall struct {
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

func (l *LogicalRoutingAndServicesAPI) ListLogicalRouters(ctx context.Context, localVarOptionals map[string]interface{}) (manager.LogicalRouterListResult, *http.Response, error) {
	l.ListLogicalRoutersCall.CallCount++
	l.ListLogicalRoutersCall.Receives.Context = ctx
	l.ListLogicalRoutersCall.Receives.LocalVarOptionals = localVarOptionals

	return l.ListLogicalRoutersCall.Returns.ListResult, l.ListLogicalRoutersCall.Returns.Response, l.ListLogicalRoutersCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) DeleteLogicalRouter(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	l.DeleteLogicalRouterCall.CallCount++
	l.DeleteLogicalRouterCall.Receives.Context = ctx
	l.DeleteLogicalRouterCall.Receives.ID = id
	l.DeleteLogicalRouterCall.Receives.LocalVarOptionals = localVarOptionals

	return l.DeleteLogicalRouterCall.Returns.Response, l.DeleteLogicalRouterCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) ListLogicalRouterPorts(ctx context.Context, localVarOptionals map[string]interface{}) (manager.LogicalRouterPortListResult, *http.Response, error) {
	l.ListLogicalRouterPortsCall.CallCount++
	l.ListLogicalRouterPortsCall.Receives.Context = ctx
	l.ListLogicalRouterPortsCall.Receives.LocalVarOptionals = localVarOptionals

	return l.ListLogicalRouterPortsCall.Returns.ListResult, l.ListLogicalRouterPortsCall.Returns.Response, l.ListLogicalRouterPortsCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) DeleteLogicalRouterPort(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	l.DeleteLogicalRouterPortCall.CallCount++
	l.DeleteLogicalRouterPortCall.Receives.Context = ctx
	l.DeleteLogicalRouterPortCall.Receives.ID = id
	l.DeleteLogicalRouterPortCall.Receives.LocalVarOptionals = localVarOptionals

	return l.DeleteLogicalRouterPortCall.Returns.Response, l.DeleteLogicalRouterPortCall.Returns.Error
}
