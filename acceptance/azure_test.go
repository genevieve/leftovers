package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/azure"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Azure", func() {
	var (
		acc     AzureAcceptance
		stdout  *bytes.Buffer
		filter  string
		deleter azure.Leftovers
	)

	BeforeEach(func() {
		iaas := os.Getenv(LEFTOVERS_ACCEPTANCE)
		if strings.ToLower(iaas) != "azure" {
			Skip("Skipping Azure acceptance tests.")
		}

		acc = NewAzureAcceptance()

		noConfirm := true
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm)

		var err error
		deleter, err = azure.NewLeftovers(logger, acc.ClientId, acc.ClientSecret, acc.SubscriptionId, acc.TenantId)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Dry run", func() {
		BeforeEach(func() {
			filter = "leftovers-dry-run"
			acc.CreateResourceGroup(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("[Resource Group: leftovers-dry-run]"))
			Expect(stdout.String()).NotTo(ContainSubstring("[Resource Group: leftovers-acceptance] Deleting..."))
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			filter = "leftovers-acceptance"
			acc.CreateResourceGroup(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[Resource Group: leftovers-acceptance] Deleting..."))
			Expect(stdout.String()).To(ContainSubstring("[Resource Group: leftovers-acceptance] Deleted!"))
		})
	})
})
