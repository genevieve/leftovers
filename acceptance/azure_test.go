package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/azure"

	. "github.com/onsi/ginkgo/v2"
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
		debug := true
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm, debug)

		var err error
		deleter, err = azure.NewLeftovers(logger, acc.ClientId, acc.ClientSecret, acc.SubscriptionId, acc.TenantId)
		Expect(err).NotTo(HaveOccurred())

		color.NoColor = true
	})

	Describe("List", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-list"
			acc.CreateResourceGroup(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("[Resource Group: %s]", filter))
			Expect(stdout.String()).To(ContainSubstring("Listing Resource Groups..."))
			Expect(stdout.String()).NotTo(ContainSubstring("[Resource Group: %s] Deleting...", filter))
		})
	})

	Describe("Types", func() {
		It("lists the resource types that can be deleted", func() {
			deleter.Types()

			Expect(stdout.String()).To(ContainSubstring("resource-group"))
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-delete"
			acc.CreateResourceGroup(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[Resource Group: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[Resource Group: %s] Deleted!", filter))
		})

		Context("when the user wants to delete subresources of the resource group", func() {
			BeforeEach(func() {
				filter = "leftovers-acc-sub-delete"
				acc.CreateResourceGroup(filter)
				// acc.CreateAppSecurityGroup(filter)
			})

			PIt("prompts them for subresources after they say no to the resource group", func() {
				err := deleter.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(stdout.String()).NotTo(ContainSubstring("[Resource Group: %s] Deleting...", filter))
				Expect(stdout.String()).NotTo(ContainSubstring("[Resource Group: %s] Deleted!", filter))

				Expect(stdout.String()).To(ContainSubstring("[Application Security Group: %s] Deleting...", filter))
				Expect(stdout.String()).To(ContainSubstring("[Application Security Group: %s] Deleted!", filter))
			})
		})
	})
})
