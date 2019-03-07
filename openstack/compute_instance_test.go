package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compute Instance", func() {
	Context("when a compute instance is created", func() {

		It("should return the compute instance", func() {
			computeInstance := openstack.NewComputeInstance("some-name", "some-id", nil)
			Expect(computeInstance.Name()).To(Equal("some-name some-id"))
		})

		It("should create a compute instance with the correct type name", func() {
			computeInstance := openstack.NewComputeInstance("some-name", "some-id", nil)
			Expect(computeInstance.Type()).To(Equal("Compute Instance"))
		})

		Context("when Delete is called", func() {
			var fakeComputeInstanceDeleter *fakes.ComputeInstanceDeleter
			var computeInstance openstack.ComputeInstance
			BeforeEach(func() {
				fakeComputeInstanceDeleter = &fakes.ComputeInstanceDeleter{}
				computeInstance = openstack.NewComputeInstance("some-name", "some-id", fakeComputeInstanceDeleter)
			})
			Context("when there is an error", func() {
				It("should return the error", func() {
					fakeComputeInstanceDeleter.DeleteCall.Returns.Error = errors.New("error description")
					err := computeInstance.Delete()

					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("error description"))
				})
			})
			It("should delete the compute instance", func() {
				err := computeInstance.Delete()
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeComputeInstanceDeleter.DeleteCall.Receives.InstanceID).To(Equal("some-id"))
			})
		})
	})
})
