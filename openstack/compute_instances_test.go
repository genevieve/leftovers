package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compute Instance", func() {
	Context("when Type method is called", func() {
		It("should return Compute Instance", func() {
			computeInstances := openstack.NewComputeInstances(nil, nil)

			result := computeInstances.Type()

			Expect(result).To(Equal("Compute Instance"))
		})
	})

	Context("when List method is called", func() {
		var (
			fakeClient       *fakes.ComputeInstanceClient
			fakeLogger       *fakes.Logger
			computeInstances openstack.ComputeInstances
		)

		BeforeEach(func() {
			fakeClient = &fakes.ComputeInstanceClient{}
			fakeLogger = &fakes.Logger{}
			computeInstances = openstack.NewComputeInstances(fakeClient, fakeLogger)
		})

		It("should return many compute instances", func() {
			fakeLogger.PromptWithDetailsCall.Returns.Bool = true
			fakeClient.ListCall.Returns.ComputeInstances = []servers.Server{
				servers.Server{
					ID:   "some id",
					Name: "some name",
				},
				servers.Server{
					ID:   "other id",
					Name: "other name",
				},
			}

			result, err := computeInstances.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(HaveLen(2))
			Expect(result[0].Name()).To(Equal("some name some id"))
			Expect(result[1].Name()).To(Equal("other name other id"))
		})

		Context("when prompt with details is false", func() {
			It("should not return a compute instance", func() {
				fakeClient.ListCall.Returns.ComputeInstances = []servers.Server{
					servers.Server{},
				}

				fakeLogger.PromptWithDetailsCall.Returns.Bool = false

				result, err := computeInstances.List()
				Expect(err).NotTo(HaveOccurred())

				Expect(result).To(HaveLen(0))
			})
		})

		Context("and there is an error", func() {
			It("should return the error", func() {
				fakeClient.ListCall.Returns.Error = errors.New("error getting list")

				result, err := computeInstances.List()
				Expect(err).To(HaveOccurred())

				Expect(result).To(BeNil())
				Expect(err).To(MatchError("error getting list"))
			})
		})
	})
})
