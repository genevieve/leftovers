package acceptance

import (
	"bytes"
	"os"
	"strings"

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
	})

	Describe("Dry run", func() {
		BeforeEach(func() {
			acc.CreateFolder(filter, "leftovers-dry-run")
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("[Folder: leftovers-dry-run]"))
			Expect(stdout.String()).NotTo(ContainSubstring("[Folder: leftovers-acceptance] Deleting..."))
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			acc.CreateFolder(filter, "leftovers-acceptance")
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[Folder: leftovers-acceptance] Deleting..."))
			Expect(stdout.String()).To(MatchRegexp(".Folder. leftovers.acceptance. .*Deleted!.*"))
		})
	})
})
