package fakes

type GroupsClient struct {
	ListGroupsCall struct {
		CallCount int
		Returns   struct {
			List  []string
			Error error
		}
	}

	DeleteGroupCall struct {
		CallCount int
		Receives  struct {
			Name string
		}
		Returns struct {
			Error error
		}
	}
}

func (i *GroupsClient) ListGroups() ([]string, error) {
	i.ListGroupsCall.CallCount++

	return i.ListGroupsCall.Returns.List, i.ListGroupsCall.Returns.Error
}

func (i *GroupsClient) DeleteGroup(name string) error {
	i.DeleteGroupCall.CallCount++
	i.DeleteGroupCall.Receives.Name = name

	return i.DeleteGroupCall.Returns.Error
}
