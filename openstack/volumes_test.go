package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/openstackfakes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volumes", func() {
	Context("when Type method is called", func() {
		It("should return volume", func() {
			volumes, err := openstack.NewVolumes(nil)
			Expect(err).ToNot(HaveOccurred())

			result := volumes.Type()

			Expect(result).To(Equal("Volume"))
		})
	})

	Context("when List method is called", func() {
		var fakeVolumesLister *openstackfakes.FakeVolumesLister
		var subject openstack.Volumes

		BeforeEach(func() {
			fakeVolumesLister = &openstackfakes.FakeVolumesLister{}
			fakeVolumesDeleter := &openstackfakes.FakeVolumesDeleter{}
			fakeVolumesServiceProvider := &openstackfakes.FakeVolumesServiceProvider{}
			fakeVolumesServiceProvider.GetVolumesListerReturns(fakeVolumesLister)
			fakeVolumesServiceProvider.GetVolumesDeleterReturns(fakeVolumesDeleter)

			var err error
			subject, err = openstack.NewVolumes(fakeVolumesServiceProvider)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("and there is a volumes service error", func() {
			It("should propogate the error", func() {
				fakeVolumesLister.ListReturns(nil, errors.New("error-description"))

				result, err := subject.List()

				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error-description"))
			})
		})

		Context("and there are no volumes", func() {
			It("should return an empty list", func() {
				fakeVolumesLister.ListReturns(nil, nil)

				result, err := subject.List()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeEmpty())
			})
		})

		Context("and there are many volumes", func() {
			It("should return the corresponding deletables", func() {
				volume := volumes.Volume{
					ID:   "some-ID",
					Name: "some-name",
				}
				otherVolume := volumes.Volume{
					ID:   "other-ID",
					Name: "some-name",
				}
				fakeVolumesLister.ListReturns([]volumes.Volume{
					volume,
					otherVolume,
				}, nil)

				result, err := subject.List()

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveLen(2))
				Expect(result[0].Name()).To(Equal("some-name some-ID"))
				Expect(result[1].Name()).To(Equal("some-name other-ID"))
				Expect((result[0].(openstack.Volume)).VolumesDeleter).NotTo(BeNil())
				Expect((result[1].(openstack.Volume)).VolumesDeleter).NotTo(BeNil())
			})
		})

	})
})
