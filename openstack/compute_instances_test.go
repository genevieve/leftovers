package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compute Instance", func() {
	var (
		fakeClient *fakes.ComputeClient
		fakeLogger *fakes.Logger
		filter     string

		computeInstances openstack.ComputeInstances
	)

	BeforeEach(func() {
		fakeClient = &fakes.ComputeClient{}
		fakeLogger = &fakes.Logger{}
		filter = ""

		computeInstances = openstack.NewComputeInstances(fakeClient, fakeLogger)
	})

	Describe("List", func() {
		BeforeEach(func() {
			fakeLogger.PromptWithDetailsCall.Returns.Bool = true
			fakeClient.ListCall.Returns.ServerSlice = []servers.Server{
				{ID: "some id", Name: "some name"},
				{ID: "other id", Name: "other name"},
			}
		})

		It("should return all compute instances", func() {
			result, err := computeInstances.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(HaveLen(2))
			Expect(result[0].Name()).To(Equal("some name some id"))
			Expect(result[1].Name()).To(Equal("other name other id"))
		})

		Context("when the resource does not contain the filter", func() {
			BeforeEach(func() {
				fakeClient.ListCall.Returns.ServerSlice = []servers.Server{
					{ID: "id", Name: "banana"},
				}
			})
			It("does not get returned", func() {
				result, err := computeInstances.List("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(result).To(HaveLen(0))
			})
		})

		Context("when prompt with details is false", func() {
			BeforeEach(func() {
				fakeLogger.PromptWithDetailsCall.Returns.Bool = false
			})

			It("should not return a compute instance", func() {
				result, err := computeInstances.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(result).To(HaveLen(0))
			})
		})

		Context("when there is an error", func() {
			BeforeEach(func() {
				fakeClient.ListCall.Returns.Error = errors.New("error getting list")
			})

			It("should return a helpful error message", func() {
				_, err := computeInstances.List(filter)
				Expect(err).To(MatchError("List Compute Instances: error getting list"))
			})
		})
	})

	Describe("Type", func() {
		It("should return Compute Instance", func() {
			result := computeInstances.Type()

			Expect(result).To(Equal("Compute Instance"))
		})
	})
})
