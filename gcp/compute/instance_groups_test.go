package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("InstanceGroups", func() {
	var (
		client *fakes.InstanceGroupsClient
		logger *fakes.Logger
		zones  map[string]string
		filter string

		instanceGroups compute.InstanceGroups
	)

	BeforeEach(func() {
		client = &fakes.InstanceGroupsClient{}
		logger = &fakes.Logger{}
		zones = map[string]string{
			"https://zone-1": "zone-1",
		}
		filter = "grape"

		instanceGroups = compute.NewInstanceGroups(client, logger, zones)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListInstanceGroupsCall.Returns.Output = &gcpcompute.InstanceGroupList{
				Items: []*gcpcompute.InstanceGroup{{
					Name: "banana",
					Zone: "https://zone-1",
				}},
			}
		})

		It("deletes instance groups", func() {
			err := instanceGroups.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListInstanceGroupsCall.CallCount).To(Equal(1))
			Expect(client.ListInstanceGroupsCall.Receives.Zone).To(Equal("zone-1"))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete instance group banana?"))

			Expect(client.DeleteInstanceGroupCall.CallCount).To(Equal(1))
			Expect(client.DeleteInstanceGroupCall.Receives.Zone).To(Equal("zone-1"))
			Expect(client.DeleteInstanceGroupCall.Receives.InstanceGroup).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting instance group banana\n"}))
		})

		Context("when the client fails to list instance groups", func() {
			BeforeEach(func() {
				client.ListInstanceGroupsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := instanceGroups.Delete(filter)
				Expect(err).To(MatchError("Listing instance groups for zone zone-1: some error"))
			})
		})

		Context("when the client fails to delete the instance", func() {
			BeforeEach(func() {
				client.DeleteInstanceGroupCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := instanceGroups.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting instance group banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the instance", func() {
				err := instanceGroups.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteInstanceGroupCall.CallCount).To(Equal(0))
			})
		})
	})
})
