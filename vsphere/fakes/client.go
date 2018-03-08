package fakes

import "github.com/vmware/govmomi/object"

type Client struct {
	GetRootFolderCall struct {
		CallCount int
		Receives  struct {
			Filter string
		}
		Returns struct {
			RootFolder *object.Folder
			Error      error
		}
	}
}

func (c Client) GetRootFolder(filter string) (*object.Folder, error) {
	c.GetRootFolderCall.CallCount++
	c.GetRootFolderCall.Receives.Filter = filter

	return c.GetRootFolderCall.Returns.RootFolder, c.GetRootFolderCall.Returns.Error
}
