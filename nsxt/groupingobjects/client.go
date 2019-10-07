package groupingobjects

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

//go:generate faux --interface groupingObjectsAPI --output fakes/grouping_objects_api.go
type groupingObjectsAPI interface {
	DeleteIPSet(context.Context, string, map[string]interface{}) (*http.Response, error)
	ListIPSets(context.Context, map[string]interface{}) (manager.IpSetListResult, *http.Response, error)

	DeleteNSService(context.Context, string, map[string]interface{}) (*http.Response, error)
	ListNSServices(context.Context, map[string]interface{}) (manager.NsServiceListResult, *http.Response, error)

	DeleteNSGroup(context.Context, string, map[string]interface{}) (*http.Response, error)
	ListNSGroups(context.Context, map[string]interface{}) (manager.NsGroupListResult, *http.Response, error)
}
