package azure_test

import (
	"errors"

	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/azure/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Group", func() {
	var (
		client *fakes.GroupsClient
		name   string

		group azure.Group
	)

	BeforeEach(func() {
		client = &fakes.GroupsClient{}
		name = "banana-group"

		group = azure.NewGroup(client, name)
	})

	Describe("Delete", func() {
		It("deletes resource groups", func() {
			err := group.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteGroupCall.CallCount).To(Equal(1))
			Expect(client.DeleteGroupCall.Receives.Name).To(Equal(name))
		})

		Context("when client fails to delete the resource group", func() {
			BeforeEach(func() {
				client.DeleteGroupCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := group.Delete()
				Expect(err).To(MatchError("Delete: some error"))
			})
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(group.Type()).To(Equal("Resource Group"))
		})
	})
})
