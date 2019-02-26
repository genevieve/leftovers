package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volumes", func() {
	Context("when Type method is called", func() {
		It("should return volume", func() {
			volumes, err := openstack.NewVolumes(nil, nil)
			Expect(err).ToNot(HaveOccurred())

			result := volumes.Type()

			Expect(result).To(Equal("Volume"))
		})
	})

	Context("when List method is called", func() {
		var (
			fakeVolumesLister *fakes.VolumesLister
			subject           openstack.Volumes
			fakeLogger        *fakes.Logger
		)

		BeforeEach(func() {
			fakeLogger = &fakes.Logger{}
			fakeVolumesLister = &fakes.VolumesLister{}
			fakeVolumesDeleter := &fakes.VolumesDeleter{}
			fakeVolumesServiceProvider := &fakes.VolumesServiceProvider{}
			fakeVolumesServiceProvider.GetVolumesListerCall.Returns.VolumesLister = fakeVolumesLister
			fakeVolumesServiceProvider.GetVolumesDeleterCall.Returns.VolumesDeleter = fakeVolumesDeleter

			var err error
			subject, err = openstack.NewVolumes(fakeVolumesServiceProvider, fakeLogger)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("and there is a volumes service error", func() {
			It("should propogate the error", func() {
				fakeVolumesLister.ListCall.Returns.Error = errors.New("error-description")

				result, err := subject.List()

				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error-description"))
			})
		})

		Context("and there are no volumes", func() {
			It("should return an empty list", func() {
				result, err := subject.List()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeEmpty())
			})
		})

		Context("and there are many volumes", func() {
			It("should return the corresponding deletables", func() {
				fakeLogger.PromptWithDetailsCall.ReturnsForCall = append(
					fakeLogger.PromptWithDetailsCall.ReturnsForCall,
					fakes.LoggerPromptWithDetailsCallReturn{Bool: true},
					fakes.LoggerPromptWithDetailsCallReturn{Bool: true},
					fakes.LoggerPromptWithDetailsCallReturn{Bool: false},
				)

				volume := volumes.Volume{
					ID:   "some-ID",
					Name: "some-name",
				}
				otherVolume := volumes.Volume{
					ID:   "other-ID",
					Name: "other-name",
				}
				anotherVolume := volumes.Volume{
					ID:   "another-ID",
					Name: "another-name",
				}
				fakeVolumesLister.ListCall.Returns.Volumes = []volumes.Volume{
					volume,
					otherVolume,
					anotherVolume,
				}

				result, err := subject.List()

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveLen(2))
				Expect(result[0].Name()).To(Equal("some-name some-ID"))
				Expect(result[1].Name()).To(Equal("other-name other-ID"))
				Expect((result[0].(openstack.Volume)).VolumesDeleter).NotTo(BeNil())
				Expect((result[1].(openstack.Volume)).VolumesDeleter).NotTo(BeNil())
			})
		})

	})
})
