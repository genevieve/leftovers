package fakes

import (
	"context"
	"net/http"
	"sync"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type GroupingObjectsAPI struct {
	DeleteIPSetCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			String             string
			MapStringInterface map[string]interface {
			}
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
		Stub func(context.Context, string, map[string]interface {
		}) (*http.Response, error)
	}
	DeleteNSGroupCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			String             string
			MapStringInterface map[string]interface {
			}
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
		Stub func(context.Context, string, map[string]interface {
		}) (*http.Response, error)
	}
	DeleteNSServiceCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			String             string
			MapStringInterface map[string]interface {
			}
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
		Stub func(context.Context, string, map[string]interface {
		}) (*http.Response, error)
	}
	ListIPSetsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			MapStringInterface map[string]interface {
			}
		}
		Returns struct {
			IpSetListResult manager.IpSetListResult
			Response        *http.Response
			Error           error
		}
		Stub func(context.Context, map[string]interface {
		}) (manager.IpSetListResult, *http.Response, error)
	}
	ListNSGroupsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			MapStringInterface map[string]interface {
			}
		}
		Returns struct {
			NsGroupListResult manager.NsGroupListResult
			Response          *http.Response
			Error             error
		}
		Stub func(context.Context, map[string]interface {
		}) (manager.NsGroupListResult, *http.Response, error)
	}
	ListNSServicesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Context            context.Context
			MapStringInterface map[string]interface {
			}
		}
		Returns struct {
			NsServiceListResult manager.NsServiceListResult
			Response            *http.Response
			Error               error
		}
		Stub func(context.Context, map[string]interface {
		}) (manager.NsServiceListResult, *http.Response, error)
	}
}

func (f *GroupingObjectsAPI) DeleteIPSet(param1 context.Context, param2 string, param3 map[string]interface {
}) (*http.Response, error) {
	f.DeleteIPSetCall.Lock()
	defer f.DeleteIPSetCall.Unlock()
	f.DeleteIPSetCall.CallCount++
	f.DeleteIPSetCall.Receives.Context = param1
	f.DeleteIPSetCall.Receives.String = param2
	f.DeleteIPSetCall.Receives.MapStringInterface = param3
	if f.DeleteIPSetCall.Stub != nil {
		return f.DeleteIPSetCall.Stub(param1, param2, param3)
	}
	return f.DeleteIPSetCall.Returns.Response, f.DeleteIPSetCall.Returns.Error
}
func (f *GroupingObjectsAPI) DeleteNSGroup(param1 context.Context, param2 string, param3 map[string]interface {
}) (*http.Response, error) {
	f.DeleteNSGroupCall.Lock()
	defer f.DeleteNSGroupCall.Unlock()
	f.DeleteNSGroupCall.CallCount++
	f.DeleteNSGroupCall.Receives.Context = param1
	f.DeleteNSGroupCall.Receives.String = param2
	f.DeleteNSGroupCall.Receives.MapStringInterface = param3
	if f.DeleteNSGroupCall.Stub != nil {
		return f.DeleteNSGroupCall.Stub(param1, param2, param3)
	}
	return f.DeleteNSGroupCall.Returns.Response, f.DeleteNSGroupCall.Returns.Error
}
func (f *GroupingObjectsAPI) DeleteNSService(param1 context.Context, param2 string, param3 map[string]interface {
}) (*http.Response, error) {
	f.DeleteNSServiceCall.Lock()
	defer f.DeleteNSServiceCall.Unlock()
	f.DeleteNSServiceCall.CallCount++
	f.DeleteNSServiceCall.Receives.Context = param1
	f.DeleteNSServiceCall.Receives.String = param2
	f.DeleteNSServiceCall.Receives.MapStringInterface = param3
	if f.DeleteNSServiceCall.Stub != nil {
		return f.DeleteNSServiceCall.Stub(param1, param2, param3)
	}
	return f.DeleteNSServiceCall.Returns.Response, f.DeleteNSServiceCall.Returns.Error
}
func (f *GroupingObjectsAPI) ListIPSets(param1 context.Context, param2 map[string]interface {
}) (manager.IpSetListResult, *http.Response, error) {
	f.ListIPSetsCall.Lock()
	defer f.ListIPSetsCall.Unlock()
	f.ListIPSetsCall.CallCount++
	f.ListIPSetsCall.Receives.Context = param1
	f.ListIPSetsCall.Receives.MapStringInterface = param2
	if f.ListIPSetsCall.Stub != nil {
		return f.ListIPSetsCall.Stub(param1, param2)
	}
	return f.ListIPSetsCall.Returns.IpSetListResult, f.ListIPSetsCall.Returns.Response, f.ListIPSetsCall.Returns.Error
}
func (f *GroupingObjectsAPI) ListNSGroups(param1 context.Context, param2 map[string]interface {
}) (manager.NsGroupListResult, *http.Response, error) {
	f.ListNSGroupsCall.Lock()
	defer f.ListNSGroupsCall.Unlock()
	f.ListNSGroupsCall.CallCount++
	f.ListNSGroupsCall.Receives.Context = param1
	f.ListNSGroupsCall.Receives.MapStringInterface = param2
	if f.ListNSGroupsCall.Stub != nil {
		return f.ListNSGroupsCall.Stub(param1, param2)
	}
	return f.ListNSGroupsCall.Returns.NsGroupListResult, f.ListNSGroupsCall.Returns.Response, f.ListNSGroupsCall.Returns.Error
}
func (f *GroupingObjectsAPI) ListNSServices(param1 context.Context, param2 map[string]interface {
}) (manager.NsServiceListResult, *http.Response, error) {
	f.ListNSServicesCall.Lock()
	defer f.ListNSServicesCall.Unlock()
	f.ListNSServicesCall.CallCount++
	f.ListNSServicesCall.Receives.Context = param1
	f.ListNSServicesCall.Receives.MapStringInterface = param2
	if f.ListNSServicesCall.Stub != nil {
		return f.ListNSServicesCall.Stub(param1, param2)
	}
	return f.ListNSServicesCall.Returns.NsServiceListResult, f.ListNSServicesCall.Returns.Response, f.ListNSServicesCall.Returns.Error
}
