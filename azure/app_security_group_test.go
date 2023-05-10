package azure_test

import (
	"errors"

	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/azure/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppSecurityGroup", func() {
	var (
		client *fakes.AppSecurityGroupsClient
		name   string
		rgName string

		group azure.AppSecurityGroup
	)

	BeforeEach(func() {
		client = &fakes.AppSecurityGroupsClient{}
		name = "banana-group"
		rgName = "major-banana-group"

		group = azure.NewAppSecurityGroup(client, rgName, name)
	})

	Describe("Delete", func() {
		It("deletes resource groups", func() {
			err := group.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteAppSecurityGroupCall.CallCount).To(Equal(1))
			Expect(client.DeleteAppSecurityGroupCall.Receives.Name).To(Equal(name))
			Expect(client.DeleteAppSecurityGroupCall.Receives.RgName).To(Equal(rgName))
		})

		Context("when client fails to delete the app security group", func() {
			BeforeEach(func() {
				client.DeleteAppSecurityGroupCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := group.Delete()
				Expect(err).To(MatchError("Delete: some error"))
			})
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(group.Type()).To(Equal("Application Security Group"))
		})
	})
})
