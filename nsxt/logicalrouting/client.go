package logicalrouting

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type logicalRoutingAPI interface {
	ListLogicalRouters(context.Context, map[string]interface{}) (manager.LogicalRouterListResult, *http.Response, error)
	DeleteLogicalRouter(context.Context, string, map[string]interface{}) (*http.Response, error)

	ListLogicalRouterPorts(context.Context, map[string]interface{}) (manager.LogicalRouterPortListResult, *http.Response, error)
	DeleteLogicalRouterPort(context.Context, string, map[string]interface{}) (*http.Response, error)
}
