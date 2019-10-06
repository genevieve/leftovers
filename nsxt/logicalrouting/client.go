package logicalrouting

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

//go:generate faux --interface logicalRoutingAPI --output fakes/logical_routing_api.go
type logicalRoutingAPI interface {
	DeleteLogicalRouter(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error)
	ListLogicalRouters(ctx context.Context, localVarOptionals map[string]interface{}) (manager.LogicalRouterListResult, *http.Response, error)
}
