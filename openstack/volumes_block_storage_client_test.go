package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("VolumesBlockStorageClient", func() {
	var (
		page   *fakes.Page
		api    *fakes.VolumesAPI
		client openstack.VolumesBlockStorageClient
	)

	BeforeEach(func() {
		page = &fakes.Page{}
		api = &fakes.VolumesAPI{}
		client = openstack.NewVolumesBlockStorageClient(api)
	})

	Describe("List", func() {
		BeforeEach(func() {
			page = &fakes.Page{}
			api.GetVolumesPagerCall.Returns.Pager = pagination.Pager{Headers: map[string]string{"header": "test"}}
			api.PagerToPageCall.Returns.Page = page
			api.PageToVolumesCall.Returns.VolumeSlice = []volumes.Volume{{Name: "volume name"}}
		})

		It("returns all the volumes", func() {
			list, err := client.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(api.PagerToPageCall.Receives.Pager.Headers["header"]).To(Equal("test"))
			Expect(api.PageToVolumesCall.Receives.Page).To(Equal(page))

			Expect(list).To(HaveLen(1))
			Expect(list[0].Name).To(Equal("volume name"))
		})

		Context("when converting a pager to page fails", func() {
			BeforeEach(func() {
				api.PagerToPageCall.Returns.Error = errors.New("error description")
			})

			It("returns a helpful error message", func() {
				_, err := client.List()
				Expect(err).To(MatchError("pager to page: error description"))
			})
		})

		Context("when converting a page to volumes fails", func() {
			BeforeEach(func() {
				api.PageToVolumesCall.Returns.Error = errors.New("error description")
			})

			It("returns a helpful error message", func() {
				_, err := client.List()
				Expect(err).To(MatchError("page to volumes: error description"))
			})
		})
	})

	Describe("Delete", func() {
		It("delete the correct volume", func() {
			err := client.Delete("some id")
			Expect(err).NotTo(HaveOccurred())

			Expect(api.DeleteVolumeCall.CallCount).To(Equal(1))
			Expect(api.DeleteVolumeCall.Receives.Id).To(Equal("some id"))
		})

		Context("when the api fails", func() {
			BeforeEach(func() {
				api.DeleteVolumeCall.Returns.Error = errors.New("some error")
			})

			It("returns an error", func() {
				err := client.Delete("some id")
				Expect(err).To(MatchError("some error"))
			})
		})
	})
})
