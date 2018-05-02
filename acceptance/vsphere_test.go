package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/vsphere"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("vSphere", func() {
	var (
		acc VSphereAcceptance

		stdout  *bytes.Buffer
		filter  string
		deleter vsphere.Leftovers
	)

	BeforeEach(func() {
		iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
		if strings.ToLower(iaas) != "vsphere" {
			Skip("Skipping vSphere acceptance tests.")
		}

		acc = NewVSphereAcceptance()

		filter = "khaleesi"
		noConfirm := true
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm)

		var err error
		deleter, err = vsphere.NewLeftovers(logger, acc.VCenterIP, acc.VCenterUser, acc.VCenterPassword, acc.Datacenter)
		Expect(err).NotTo(HaveOccurred())

		color.NoColor = true
	})

	Describe("leftovers", func() {
		BeforeEach(func() {
			acc.CreateFolder(filter, "leftovers-acceptance")
		})

		It("can list and delete resources with the filter", func() {
			By("listing resources first", func() {
				deleter.List(filter)

				Expect(stdout.String()).To(ContainSubstring("[Folder: leftovers-acceptance]"))
				Expect(stdout.String()).NotTo(ContainSubstring("[Folder: leftovers-acceptance] Deleting..."))
			})

			By("successfully deleting resources", func() {
				err := deleter.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(stdout.String()).To(ContainSubstring("[Folder: leftovers-acceptance] Deleting..."))
				Expect(stdout.String()).To(MatchRegexp(".Folder. leftovers.acceptance. .*Deleted!.*"))
			})
		})
	})
})
