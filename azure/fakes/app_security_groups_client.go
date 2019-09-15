package fakes

type AppSecurityGroupsClient struct {
	ListAppSecurityGroupsCall struct {
		CallCount int
		Receives  struct {
			ResourceGroupName string
		}
		Returns struct {
			List  []string
			Error error
		}
	}

	DeleteAppSecurityGroupCall struct {
		CallCount int
		Receives  struct {
			ResourceGroupName string
			Name              string
		}
		Returns struct {
			Error error
		}
	}
}

func (c *AppSecurityGroupsClient) ListAppSecurityGroups(rgName string) ([]string, error) {
	c.ListAppSecurityGroupsCall.CallCount++
	c.ListAppSecurityGroupsCall.Receives.ResourceGroupName = rgName
	return c.ListAppSecurityGroupsCall.Returns.List, c.ListAppSecurityGroupsCall.Returns.Error
}

func (c *AppSecurityGroupsClient) DeleteAppSecurityGroup(rgName, name string) error {
	c.DeleteAppSecurityGroupCall.CallCount++
	c.DeleteAppSecurityGroupCall.Receives.ResourceGroupName = rgName
	c.DeleteAppSecurityGroupCall.Receives.Name = name

	return c.DeleteAppSecurityGroupCall.Returns.Error
}
