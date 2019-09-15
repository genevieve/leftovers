package azure

import "fmt"

type AppSecurityGroup struct {
	client appSecurityGroupsClient
	name   string
	rgName string
}

// AppSecurityGroup represents an Azure application security group.
func NewAppSecurityGroup(client appSecurityGroupsClient, rgName, name string) AppSecurityGroup {
	return AppSecurityGroup{
		client: client,
		rgName: rgName,
		name:   name,
	}
}

// Delete deletes an Azure application security group.
func (g AppSecurityGroup) Delete() error {
	err := g.client.DeleteAppSecurityGroup(g.rgName, g.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (g AppSecurityGroup) Name() string {
	return g.name
}

func (g AppSecurityGroup) Type() string {
	return "Application Security Group"
}
