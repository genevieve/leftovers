package artifacts_test

import (
	"errors"
	"github.com/genevieve/leftovers/gcp/artifacts"
	"github.com/genevieve/leftovers/gcp/artifacts/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cluster", func() {
	var (
		client *fakes.RepositoriesClient
		name   string

		repository artifacts.Repository
	)

	BeforeEach(func() {
		client = &fakes.RepositoriesClient{}
		name = "banana"

		repository = artifacts.NewRepository(client, name)
	})

	Describe("Delete", func() {
		It("deletes the resource", func() {
			err := repository.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteRepositoryCall.Receives.Cluster).To(Equal("banana"))
		})

		Context("when the client returns an error", func() {
			BeforeEach(func() {
				client.DeleteRepositoryCall.Returns.Error = errors.New("kiwi")
			})

			It("returns a helpful error message", func() {
				err := repository.Delete()
				Expect(err).To(MatchError("Delete: kiwi"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(repository.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(repository.Type()).To(Equal("Artifact Repository"))
		})
	})
})
