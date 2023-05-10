package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compute Instance", func() {
	Describe("NewComputeInstance", func() {
		It("has a name and type", func() {
			computeInstance := openstack.NewComputeInstance("some-name", "some-id", nil)
			Expect(computeInstance.Name()).To(Equal("some-name some-id"))
			Expect(computeInstance.Type()).To(Equal("Compute Instance"))
		})
	})

	Describe("Delete", func() {
		var (
			fakeComputeClient *fakes.ComputeClient
			computeInstance   openstack.ComputeInstance
		)

		BeforeEach(func() {
			fakeComputeClient = &fakes.ComputeClient{}
			computeInstance = openstack.NewComputeInstance("some-name", "some-id", fakeComputeClient)
		})

		It("deletes the compute instance", func() {
			err := computeInstance.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeComputeClient.DeleteCall.Receives.InstanceID).To(Equal("some-id"))
		})

		Context("when an error occurs", func() {
			It("returns an error", func() {
				fakeComputeClient.DeleteCall.Returns.Error = errors.New("error description")

				err := computeInstance.Delete()
				Expect(err).To(HaveOccurred())

				Expect(err).To(MatchError("error description"))
			})
		})
	})
})
