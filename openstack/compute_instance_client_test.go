package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ComputeInstanceClient", func() {
	var (
		computeInstanceAPI    *fakes.ComputeInstanceAPI
		computeInstanceClient openstack.ComputeInstanceClient
	)
	BeforeEach(func() {
		computeInstanceAPI = &fakes.ComputeInstanceAPI{}
		computeInstanceClient = openstack.NewComputeInstanceClient(computeInstanceAPI)
	})
	Context("when listing", func() {
		Context("when converting a pager to a page returns an error", func() {
			It("should propogate the error", func() {
				computeInstanceAPI.PagerToPageCall.Returns.Error = errors.New("error description")

				result, err := computeInstanceClient.List()
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error description"))
			})
		})

		Context("when converting a page to servers returns an error", func() {
			It("should propogate the error", func() {
				computeInstanceAPI.PageToServersCall.Returns.Error = errors.New("error description")

				result, err := computeInstanceClient.List()
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error description"))
			})
		})

		Context("when servers exist and there are no errors", func() {
			It("should return a list of servers", func() {
				pager := pagination.Pager{Headers: map[string]string{"header": "test"}}
				computeInstanceAPI.GetComputeInstancePagerCall.Returns.ComputeInstancePager = pager
				computeInstanceAPI.PageToServersCall.Returns.Servers = []servers.Server{servers.Server{ID: "server-id"}}
				computeInstanceAPI.PagerToPageCall.Returns.Page = fakes.Page{Name: "server page"}

				result, err := computeInstanceClient.List()

				Expect(computeInstanceAPI.GetComputeInstancePagerCall.CallCount).To(Equal(1))
				Expect(computeInstanceAPI.PagerToPageCall.CallCount).To(Equal(1))
				pagerToPageArgCapture := computeInstanceAPI.PagerToPageCall.Receives.Pager
				Expect(pagerToPageArgCapture.Headers["header"]).To(Equal("test"))
				Expect(computeInstanceAPI.PageToServersCall.CallCount).To(Equal(1))
				Expect(((computeInstanceAPI.PageToServersCall.Receives.Page).(fakes.Page)).Name).To(Equal("server page"))
				Expect(len(result)).To(Equal(1))
				Expect(result[0].ID).To(Equal("server-id"))
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Context("when deleting", func() {
		Context("when there is an error", func() {
			It("should return the error", func() {
				computeInstanceAPI.DeleteCall.Returns.Error = errors.New("error deleting instance")
				err := computeInstanceClient.Delete("some instance id")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error deleting instance"))
			})
		})
		It("should delete the compute instance", func() {
			err := computeInstanceClient.Delete("some instance id")

			Expect(computeInstanceAPI.DeleteCall.Receives.InstanceID).To(Equal("some instance id"))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
