package acceptance_test

import (
	"bytes"
	"os"

	"github.com/genevievelesperance/leftovers/acceptance"
	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/azure"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Azure", func() {
	var acc *acceptance.AzureAcceptance

	BeforeEach(func() {
		acc = acceptance.NewAzureAcceptance()

		if !acc.ReadyToTest() {
			Skip("Skipping acceptance tests.")
		}
	})

	Describe("Delete", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter azure.Leftovers
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "leftovers-acceptance"
			acc.CreateResourceGroup(filter)

			var err error
			deleter, err = azure.NewLeftovers(logger, acc.ClientId, acc.ClientSecret, acc.SubscriptionId, acc.TenantId)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout).To(ContainSubstring("SUCCESS deleting resource group"))
		})
	})
})
