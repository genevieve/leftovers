package fakes

import "sync"

type AppSecurityGroupsClient struct {
	DeleteAppSecurityGroupCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			RgName string
			Name   string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListAppSecurityGroupsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			RgName string
		}
		Returns struct {
			StringSlice []string
			Error       error
		}
		Stub func(string) ([]string, error)
	}
}

func (f *AppSecurityGroupsClient) DeleteAppSecurityGroup(param1 string, param2 string) error {
	f.DeleteAppSecurityGroupCall.Lock()
	defer f.DeleteAppSecurityGroupCall.Unlock()
	f.DeleteAppSecurityGroupCall.CallCount++
	f.DeleteAppSecurityGroupCall.Receives.RgName = param1
	f.DeleteAppSecurityGroupCall.Receives.Name = param2
	if f.DeleteAppSecurityGroupCall.Stub != nil {
		return f.DeleteAppSecurityGroupCall.Stub(param1, param2)
	}
	return f.DeleteAppSecurityGroupCall.Returns.Error
}
func (f *AppSecurityGroupsClient) ListAppSecurityGroups(param1 string) ([]string, error) {
	f.ListAppSecurityGroupsCall.Lock()
	defer f.ListAppSecurityGroupsCall.Unlock()
	f.ListAppSecurityGroupsCall.CallCount++
	f.ListAppSecurityGroupsCall.Receives.RgName = param1
	if f.ListAppSecurityGroupsCall.Stub != nil {
		return f.ListAppSecurityGroupsCall.Stub(param1)
	}
	return f.ListAppSecurityGroupsCall.Returns.StringSlice, f.ListAppSecurityGroupsCall.Returns.Error
}
