package logicalrouting

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type logicalRoutingAPI interface {
	DeleteLogicalRouter(context.Context, string, map[string]interface{}) (*http.Response, error)
	ListLogicalRouters(context.Context, map[string]interface{}) (manager.LogicalRouterListResult, *http.Response, error)
}
