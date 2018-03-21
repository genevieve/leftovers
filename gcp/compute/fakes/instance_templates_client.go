package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type InstanceTemplatesClient struct {
	ListInstanceTemplatesCall struct {
		CallCount int
		Returns   struct {
			Output *gcpcompute.InstanceTemplateList
			Error  error
		}
	}

	DeleteInstanceTemplateCall struct {
		CallCount int
		Receives  struct {
			InstanceTemplate string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *InstanceTemplatesClient) ListInstanceTemplates() (*gcpcompute.InstanceTemplateList, error) {
	n.ListInstanceTemplatesCall.CallCount++

	return n.ListInstanceTemplatesCall.Returns.Output, n.ListInstanceTemplatesCall.Returns.Error
}

func (n *InstanceTemplatesClient) DeleteInstanceTemplate(instanceTemplate string) error {
	n.DeleteInstanceTemplateCall.CallCount++
	n.DeleteInstanceTemplateCall.Receives.InstanceTemplate = instanceTemplate

	return n.DeleteInstanceTemplateCall.Returns.Error
}
