package acceptance_test

import (
	"bytes"
	"os"

	"github.com/genevievelesperance/leftovers"
	"github.com/genevievelesperance/leftovers/acceptance"
	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/gcp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GCP", func() {
	var acc *acceptance.GCPAcceptance

	BeforeEach(func() {
		acc = acceptance.NewGCPAcceptance()

		if !acc.ReadyToTest() {
			Skip("Skipping acceptance tests.")
		}
	})

	Describe("Delete", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter leftovers.Deleter
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "leftovers-acceptance"
			acc.InsertDisk(filter)

			var err error
			deleter, err = gcp.NewLeftovers(logger, acc.KeyPath)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout).To(ContainSubstring("SUCCESS deleting disk"))
		})
	})
})
