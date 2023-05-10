package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	openstackvolumes "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volumes", func() {
	var (
		volumes openstack.Volumes
		logger  *fakes.Logger
		filter  string

		client *fakes.VolumesClient
	)

	BeforeEach(func() {
		logger = &fakes.Logger{}
		client = &fakes.VolumesClient{}
		filter = ""

		volumes = openstack.NewVolumes(client, logger)
	})

	Describe("List", func() {
		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Bool = true
			client.ListCall.Returns.VolumeSlice = []openstackvolumes.Volume{
				{ID: "some-ID", Name: "some-name"},
				{ID: "other-ID", Name: "other-name"},
			}
		})

		It("returns all the deletables", func() {
			list, err := volumes.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(list).To(HaveLen(2))
			Expect(list[0].Name()).To(Equal("some-name some-ID"))
			Expect(list[1].Name()).To(Equal("other-name other-ID"))
		})

		Context("when the resource does not contain the filter", func() {
			BeforeEach(func() {
				client.ListCall.Returns.VolumeSlice = []openstackvolumes.Volume{
					{ID: "id", Name: "banana"},
				}
			})

			It("returns all the deletables", func() {
				list, err := volumes.List("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})

		Context("when listing fails", func() {
			BeforeEach(func() {
				client.ListCall.Returns.Error = errors.New("error-description")
			})

			It("returns an error", func() {
				_, err := volumes.List(filter)
				Expect(err).To(MatchError("List Volumes: error-description"))
			})
		})
	})

	Describe("Type", func() {
		It("returns the type of the resource", func() {
			Expect(volumes.Type()).To(Equal("Volume"))
		})
	})
})
