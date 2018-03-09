package acceptance

import (
	"bytes"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/azure"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Azure", func() {
	var acc *AzureAcceptance

	BeforeEach(func() {
		acc = NewAzureAcceptance()

		if !acc.ReadyToTest() {
			Skip("Skipping acceptance tests.")
		}
	})

	Describe("Dry run", func() {
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

			filter = "leftovers-acceptance-dry-run"
			acc.CreateResourceGroup(filter)

			var err error
			deleter, err = azure.NewLeftovers(logger, acc.ClientId, acc.ClientSecret, acc.SubscriptionId, acc.TenantId)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("resource group: leftovers-acceptance-dry-run"))
			Expect(stdout.String()).NotTo(ContainSubstring("SUCCESS deleting leftovers-acceptance-dry-run"))
			Expect(stdout.String()).NotTo(ContainSubstring("FAILED deleting resource group"))
		})
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

			Expect(stdout.String()).To(ContainSubstring("SUCCESS deleting leftovers-acceptance"))
			Expect(stdout.String()).NotTo(ContainSubstring("FAILED deleting resource group"))
		})
	})
})
