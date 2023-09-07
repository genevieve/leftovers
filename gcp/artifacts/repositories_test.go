package artifacts_test

import (
	"errors"
	"github.com/genevieve/leftovers/gcp/artifacts"
	"github.com/genevieve/leftovers/gcp/artifacts/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gcpartifact "google.golang.org/api/artifactregistry/v1"
)

var _ = Describe("Clusters", func() {
	var (
		client *fakes.RepositoriesClient
		logger *fakes.Logger
		filter string
		zones  map[string]string

		repositories artifacts.Repositories
	)

	BeforeEach(func() {
		client = &fakes.RepositoriesClient{}
		logger = &fakes.Logger{}
		filter = "banana"
		zones = map[string]string{"https://zone-1": "zone-1"}

		logger.PromptWithDetailsCall.Returns.Proceed = true

		repositories = artifacts.NewRepositories(client, logger, zones)
	})

	Describe("List", func() {
		BeforeEach(func() {
			client.ListRepositoriesCall.Returns.ListRepositoriesResponse = []*gcpartifact.Repository{{
				Name: "banana-repository",
			}}
		})

		It("returns a list of clusters to delete", func() {
			list, err := repositories.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Artifact Repository"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("banana-repository"))

			Expect(list).To(HaveLen(1))
		})

		Context("when the user does not want to delete that resource", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not return the resource in the list", func() {
				list, err := repositories.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the resource name does not contain the filter", func() {
			BeforeEach(func() {
				client.ListRepositoriesCall.Returns.ListRepositoriesResponse = []*gcpartifact.Repository{{
					Name: "kiwi-repository",
				}}
			})

			It("does not return the resource in the list", func() {
				list, err := repositories.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the client returns an error", func() {
			BeforeEach(func() {
				client.ListRepositoriesCall.Returns.Error = errors.New("panic time")
			})

			It("wraps it in a helpful error message", func() {
				_, err := repositories.List(filter)
				Expect(err).To(MatchError(client.ListRepositoriesCall.Returns.Error))
			})
		})
	})
})
