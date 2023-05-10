package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/nsxt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NSX-T", func() {
	var (
		acc NSXTAcceptance

		stdout  *bytes.Buffer
		deleter nsxt.Leftovers
	)

	BeforeEach(func() {
		iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
		if strings.ToLower(iaas) != "nsxt" {
			Skip("Skipping NSX-T acceptance tests.")
		}

		acc = NewNSXTAcceptance()

		noConfirm := true
		debug := true
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm, debug)

		var err error
		deleter, err = nsxt.NewLeftovers(logger, acc.ManagerHost, acc.User, acc.Password)
		Expect(err).NotTo(HaveOccurred())

		color.NoColor = true
	})

	Describe("leftovers", func() {
		BeforeEach(func() {
			acc.CreateT1Router("leftover-tier1-router")
		})

		It("can list and delete resources with the filter", func() {
			By("listing resources first", func() {
				deleter.List("leftover")
				Expect(stdout.String()).NotTo(ContainSubstring("403"))

				Expect(stdout.String()).To(ContainSubstring("Listing Tier 1 Routers..."))
				Expect(stdout.String()).To(ContainSubstring("[Tier 1 Router: leftover-tier1-router]"))
				Expect(stdout.String()).NotTo(ContainSubstring("[Tier 1 Router: leftover-tier1-router] Deleting"))
			})

			By("successfully deleting resources", func() {
				err := deleter.Delete("leftover")
				Expect(err).NotTo(HaveOccurred())

				Expect(stdout.String()).To(ContainSubstring("[Tier 1 Router: leftover-tier1-router] Deleting..."))
				Expect(stdout.String()).To(ContainSubstring("[Tier 1 Router: leftover-tier1-router] Deleted!"))

				Expect(stdout.String()).NotTo(ContainSubstring("[Tier 1 Router: toolsmiths-T1] Delet"))
			})
		})
	})
})
