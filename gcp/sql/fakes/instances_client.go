package fakes

import gcpsql "google.golang.org/api/sqladmin/v1beta4"

type InstancesClient struct {
	ListInstancesCall struct {
		CallCount int
		Returns   struct {
			Output *gcpsql.InstancesListResponse
			Error  error
		}
	}

	DeleteInstanceCall struct {
		CallCount int
		Receives  struct {
			Instance string
		}
		Returns struct {
			Error error
		}
	}
}

func (u *InstancesClient) ListInstances() (*gcpsql.InstancesListResponse, error) {
	u.ListInstancesCall.CallCount++

	return u.ListInstancesCall.Returns.Output, u.ListInstancesCall.Returns.Error
}

func (u *InstancesClient) DeleteInstance(instance string) error {
	u.DeleteInstanceCall.CallCount++
	u.DeleteInstanceCall.Receives.Instance = instance

	return u.DeleteInstanceCall.Returns.Error
}
