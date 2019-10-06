package fakes

import "sync"

type GroupsClient struct {
	DeleteGroupCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Name string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListGroupsCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Groups []string
			Err    error
		}
		Stub func() ([]string, error)
	}
}

func (f *GroupsClient) DeleteGroup(param1 string) error {
	f.DeleteGroupCall.Lock()
	defer f.DeleteGroupCall.Unlock()
	f.DeleteGroupCall.CallCount++
	f.DeleteGroupCall.Receives.Name = param1
	if f.DeleteGroupCall.Stub != nil {
		return f.DeleteGroupCall.Stub(param1)
	}
	return f.DeleteGroupCall.Returns.Error
}
func (f *GroupsClient) ListGroups() ([]string, error) {
	f.ListGroupsCall.Lock()
	defer f.ListGroupsCall.Unlock()
	f.ListGroupsCall.CallCount++
	if f.ListGroupsCall.Stub != nil {
		return f.ListGroupsCall.Stub()
	}
	return f.ListGroupsCall.Returns.Groups, f.ListGroupsCall.Returns.Err
}
