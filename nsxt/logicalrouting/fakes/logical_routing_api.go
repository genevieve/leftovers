package fakes

import (
	"context"
	"net/http"
	"sync"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type LogicalRoutingAPI struct {
	DeleteLogicalRouterCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Ctx               context.Context
			Id                string
			LocalVarOptionals map[string]interface {
			}
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
		Stub func(context.Context, string, map[string]interface {
		}) (*http.Response, error)
	}
	ListLogicalRoutersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Ctx               context.Context
			LocalVarOptionals map[string]interface {
			}
		}
		Returns struct {
			LogicalRouterListResult manager.LogicalRouterListResult
			Response                *http.Response
			Error                   error
		}
		Stub func(context.Context, map[string]interface {
		}) (manager.LogicalRouterListResult, *http.Response, error)
	}
}

func (f *LogicalRoutingAPI) DeleteLogicalRouter(param1 context.Context, param2 string, param3 map[string]interface {
}) (*http.Response, error) {
	f.DeleteLogicalRouterCall.Lock()
	defer f.DeleteLogicalRouterCall.Unlock()
	f.DeleteLogicalRouterCall.CallCount++
	f.DeleteLogicalRouterCall.Receives.Ctx = param1
	f.DeleteLogicalRouterCall.Receives.Id = param2
	f.DeleteLogicalRouterCall.Receives.LocalVarOptionals = param3
	if f.DeleteLogicalRouterCall.Stub != nil {
		return f.DeleteLogicalRouterCall.Stub(param1, param2, param3)
	}
	return f.DeleteLogicalRouterCall.Returns.Response, f.DeleteLogicalRouterCall.Returns.Error
}
func (f *LogicalRoutingAPI) ListLogicalRouters(param1 context.Context, param2 map[string]interface {
}) (manager.LogicalRouterListResult, *http.Response, error) {
	f.ListLogicalRoutersCall.Lock()
	defer f.ListLogicalRoutersCall.Unlock()
	f.ListLogicalRoutersCall.CallCount++
	f.ListLogicalRoutersCall.Receives.Ctx = param1
	f.ListLogicalRoutersCall.Receives.LocalVarOptionals = param2
	if f.ListLogicalRoutersCall.Stub != nil {
		return f.ListLogicalRoutersCall.Stub(param1, param2)
	}
	return f.ListLogicalRoutersCall.Returns.LogicalRouterListResult, f.ListLogicalRoutersCall.Returns.Response, f.ListLogicalRoutersCall.Returns.Error
}
