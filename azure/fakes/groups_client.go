package fakes

import (
	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/go-autorest/autorest"
)

type GroupsClient struct {
	ListCall struct {
		CallCount int
		Receives  struct {
			Filter string
			Top    *int32
		}
		Returns struct {
			Output resources.GroupListResult
			Error  error
		}
	}

	DeleteCall struct {
		CallCount int
		Receives  struct {
			Name    string
			Channel <-chan struct{}
		}
		Returns struct {
			Output <-chan autorest.Response
			Error  <-chan error
		}
	}
}

func (i *GroupsClient) List(filter string, top *int32) (resources.GroupListResult, error) {
	i.ListCall.CallCount++
	i.ListCall.Receives.Filter = filter
	i.ListCall.Receives.Top = top

	return i.ListCall.Returns.Output, i.ListCall.Returns.Error
}

func (i *GroupsClient) Delete(name string, channel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	i.DeleteCall.CallCount++
	i.DeleteCall.Receives.Name = name
	i.DeleteCall.Receives.Channel = channel

	return i.DeleteCall.Returns.Output, i.DeleteCall.Returns.Error
}
