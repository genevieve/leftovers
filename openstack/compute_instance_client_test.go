package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ComputeInstanceClient", func() {
	var (
		api    *fakes.ComputeInstanceAPI
		client openstack.ComputeInstanceClient
		page   *fakes.Page
	)

	BeforeEach(func() {
		api = &fakes.ComputeInstanceAPI{}
		client = openstack.NewComputeInstanceClient(api)
		page = &fakes.Page{}
	})

	Describe("List", func() {
		BeforeEach(func() {
			pager := pagination.Pager{Headers: map[string]string{"header": "test"}}
			api.GetComputeInstancePagerCall.Returns.Pager = pager
			api.PageToServersCall.Returns.ServerSlice = []servers.Server{servers.Server{ID: "server-id"}}
			api.PagerToPageCall.Returns.Page = page
		})

		It("returns a list of servers", func() {
			result, err := client.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(api.GetComputeInstancePagerCall.CallCount).To(Equal(1))
			Expect(api.PagerToPageCall.CallCount).To(Equal(1))

			pagerToPageArgCapture := api.PagerToPageCall.Receives.Pager
			Expect(pagerToPageArgCapture.Headers["header"]).To(Equal("test"))

			Expect(api.PageToServersCall.CallCount).To(Equal(1))
			Expect(api.PageToServersCall.Receives.Page).To(Equal(page))

			Expect(len(result)).To(Equal(1))
			Expect(result[0].ID).To(Equal("server-id"))
		})

		Context("when converting a pager to a page returns an error", func() {
			BeforeEach(func() {
				api.PagerToPageCall.Returns.Error = errors.New("banana")
			})
			It("should propogate the error", func() {
				_, err := client.List()
				Expect(err).To(MatchError("pager to page: banana"))
			})
		})

		Context("when converting a page to servers returns an error", func() {
			BeforeEach(func() {
				api.PageToServersCall.Returns.Error = errors.New("banana")

			})
			It("should propogate the error", func() {
				_, err := client.List()
				Expect(err).To(MatchError("page to servers: banana"))
			})
		})
	})

	Describe("Delete", func() {
		It("deletes the compute instance", func() {
			err := client.Delete("some instance id")
			Expect(err).NotTo(HaveOccurred())

			Expect(api.DeleteCall.Receives.InstanceID).To(Equal("some instance id"))
		})

		Context("when the api returns an error", func() {
			BeforeEach(func() {
				api.DeleteCall.Returns.Error = errors.New("error deleting instance")
			})

			It("should return a helpful error message", func() {
				err := client.Delete("some instance id")
				Expect(err).To(MatchError("error deleting instance"))
			})
		})
	})
})
