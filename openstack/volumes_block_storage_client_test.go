package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VolumesBlockStorageClient", func() {
	Context("when listing", func() {
		var (
			volumesAPI                *fakes.VolumesAPI
			volumesBlockStorageClient openstack.VolumesBlockStorageClient
		)
		BeforeEach(func() {
			volumesAPI = &fakes.VolumesAPI{}
			volumesBlockStorageClient = openstack.NewVolumesBlockStorageClient(nil, volumesAPI)
		})
		Context("when converting a pager to a page returns an error", func() {
			It("should propogate the error", func() {
				volumesAPI.PagerToPageCall.Returns.Error = errors.New("error description")

				result, err := volumesBlockStorageClient.List()
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error description"))
			})
		})

		Context("when converting a page to volumes returns an error", func() {
			It("should propogate the error", func() {
				volumesAPI.PageToVolumesCall.Returns.Error = errors.New("error description")

				result, err := volumesBlockStorageClient.List()

				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error description"))
			})
		})

		Context("when a volume exists and no errors occur", func() {
			It("should return the volume", func() {
				volumesAPI.GetVolumesPagerCall.Returns.Pager = pagination.Pager{Headers: map[string]string{"header": "test"}}
				volumesAPI.PagerToPageCall.Returns.Page = fakes.Page{Name: "page name"}
				volume := volumes.Volume{Name: "volume name"}
				volumesAPI.PageToVolumesCall.Returns.Volumes = []volumes.Volume{volume}

				result, err := volumesBlockStorageClient.List()

				Expect(volumesAPI.PagerToPageCall.Receives.Pager.Headers["header"]).To(Equal("test"))
				Expect((volumesAPI.PageToVolumesCall.Receives.Page.(fakes.Page)).Name).To(Equal("page name"))
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveLen(1))
				Expect(result[0].Name).To(Equal("volume name"))
			})
		})
	})

	Context("when deleting", func() {
		var (
			volumesAPI                *fakes.VolumesAPI
			volumesBlockStorageClient openstack.VolumesBlockStorageClient
		)
		BeforeEach(func() {
			volumesAPI = &fakes.VolumesAPI{}
			volumesBlockStorageClient = openstack.NewVolumesBlockStorageClient(nil, volumesAPI)
		})
		Context("when the client returns an error", func() {
			It("should propogate the error", func() {
				volumesAPI.DeleteVolumeCall.Returns.Error = errors.New("some error")

				err := volumesBlockStorageClient.Delete("some id")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("some error"))
			})
		})
		It("delete the correct volume", func() {
			err := volumesBlockStorageClient.Delete("some id")
			Expect(err).NotTo(HaveOccurred())
			err = volumesBlockStorageClient.Delete("some other id")
			Expect(err).NotTo(HaveOccurred())
			Expect(volumesAPI.DeleteVolumeCall.CallCount).To(Equal(2))

			id := volumesAPI.DeleteVolumeCall.ReceivesForCall[0].VolumeID
			Expect(id).To(Equal("some id"))
			id = volumesAPI.DeleteVolumeCall.ReceivesForCall[1].VolumeID
			Expect(id).To(Equal("some other id"))
		})
	})
})
