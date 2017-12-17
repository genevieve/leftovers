package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("Disks", func() {
	var (
		client *fakes.DisksClient
		logger *fakes.Logger
		zones  map[string]string

		disks compute.Disks
	)

	BeforeEach(func() {
		client = &fakes.DisksClient{}
		logger = &fakes.Logger{}
		zones = map[string]string{
			"https://zone-1": "zone-1",
		}

		disks = compute.NewDisks(client, logger, zones)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListDisksCall.Returns.Output = &gcpcompute.DiskList{
				Items: []*gcpcompute.Disk{{
					Name: "banana",
					Zone: "https://zone-1",
				}},
			}
		})

		It("deletes disks", func() {
			err := disks.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListDisksCall.CallCount).To(Equal(1))
			Expect(client.ListDisksCall.Receives.Zone).To(Equal("zone-1"))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete disk banana?"))

			Expect(client.DeleteDiskCall.CallCount).To(Equal(1))
			Expect(client.DeleteDiskCall.Receives.Zone).To(Equal("zone-1"))
			Expect(client.DeleteDiskCall.Receives.Disk).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting disk banana\n"}))
		})

		Context("when the client fails to list disks", func() {
			BeforeEach(func() {
				client.ListDisksCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := disks.Delete()
				Expect(err).To(MatchError("Listing disks for zone zone-1: some error"))
			})
		})

		Context("when the disk is in use by an instance", func() {
			BeforeEach(func() {
				client.ListDisksCall.Returns.Output = &gcpcompute.DiskList{
					Items: []*gcpcompute.Disk{{
						Name:  "banana",
						Zone:  "zone-1",
						Users: []string{"instance-using-banana"},
					}},
				}
			})

			It("does not try deleting it", func() {
				err := disks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ListDisksCall.CallCount).To(Equal(1))
				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteDiskCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the disk", func() {
			BeforeEach(func() {
				client.DeleteDiskCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := disks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting disk banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the disk", func() {
				err := disks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteDiskCall.CallCount).To(Equal(0))
			})
		})
	})
})
