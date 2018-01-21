package acceptance_test

import (
	"os"

	"github.com/genevievelesperance/leftovers"
	"github.com/genevievelesperance/leftovers/acceptance"
	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/gcp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GCP", func() {
	BeforeEach(func() {
		if !acceptance.ReadyToTest() {
			Skip("Skipping acceptance tests.")
		}
	})

	Describe("Delete", func() {
		var (
			deleter           leftovers.Deleter
			logger            *app.Logger
			serviceAccountKey string
		)

		BeforeEach(func() {
			logger = app.NewLogger(os.Stdout, os.Stdin, false)
			serviceAccountKey = os.Getenv("BBL_GCP_SERVICE_ACCOUNT_KEY")

			var err error
			deleter, err = gcp.NewDeleter(logger, serviceAccountKey)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete("fake")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
