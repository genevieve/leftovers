package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/openstackfakes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakePage struct {
	name string
}

func (fakePage fakePage) NextPageURL() (string, error) {
	return "", nil
}

func (fakePage fakePage) IsEmpty() (bool, error) {
	return false, nil
}

func (fakePage fakePage) GetBody() interface{} {
	return nil
}

var _ = Describe("VolumesBlockStorageClient", func() {
	Context("when listing", func() {
		var volumesAPI *openstackfakes.FakeVolumesAPI
		var volumesBlockStorageClient openstack.VolumesBlockStorageClient
		BeforeEach(func() {
			volumesAPI = &openstackfakes.FakeVolumesAPI{}
			volumesBlockStorageClient = openstack.NewVolumesBlockStorageClient(nil, volumesAPI)
		})
		Context("when converting a pager to a page returns an error", func() {
			It("should propogate the error", func() {
				volumesAPI.PagerToPageReturns(nil, errors.New("error description"))

				result, err := volumesBlockStorageClient.List()
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error description"))
			})
		})

		Context("when converting a page to volumes returns an error", func() {
			It("should propogate the error", func() {
				volumesAPI.PageToVolumesReturns(nil, errors.New("error description"))

				result, err := volumesBlockStorageClient.List()

				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error description"))
			})
		})

		Context("when a volume exists and no errors occur", func() {
			It("should return the volume", func() {
				volumesAPI.GetVolumesPagerReturns(pagination.Pager{Headers: map[string]string{"header": "test"}})
				volumesAPI.PagerToPageReturns(fakePage{name: "page name"}, nil)
				volume := volumes.Volume{Name: "volume name"}
				volumesAPI.PageToVolumesReturns([]volumes.Volume{volume}, nil)

				result, err := volumesBlockStorageClient.List()

				Expect(volumesAPI.PagerToPageArgsForCall(0).Headers["header"]).To(Equal("test"))
				Expect((volumesAPI.PageToVolumesArgsForCall(0).(fakePage)).name).To(Equal("page name"))
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveLen(1))
				Expect(result[0].Name).To(Equal("volume name"))
			})
		})
	})
	Context("when deleting", func() {
		var volumesAPI *openstackfakes.FakeVolumesAPI
		var volumesBlockStorageClient openstack.VolumesBlockStorageClient
		BeforeEach(func() {
			volumesAPI = &openstackfakes.FakeVolumesAPI{}
			volumesBlockStorageClient = openstack.NewVolumesBlockStorageClient(nil, volumesAPI)
		})
		Context("when the client returns an error", func() {
			It("should propogate the error", func() {
				volumesAPI.DeleteVolumeReturns(errors.New("some error"))

				err := volumesBlockStorageClient.Delete("some id")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("some error"))
			})
		})
		It("delete the correct volume", func() {
			volumesAPI.DeleteVolumeReturns(nil)

			err := volumesBlockStorageClient.Delete("some id")
			Expect(err).NotTo(HaveOccurred())
			err = volumesBlockStorageClient.Delete("some other id")
			Expect(err).NotTo(HaveOccurred())
			Expect(volumesAPI.DeleteVolumeCallCount()).To(Equal(2))

			_, id, _ := volumesAPI.DeleteVolumeArgsForCall(0)
			Expect(id).To(Equal("some id"))
			_, id, _ = volumesAPI.DeleteVolumeArgsForCall(1)
			Expect(id).To(Equal("some other id"))
		})
	})
})
