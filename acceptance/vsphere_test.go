package acceptance

import (
	"bytes"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/vsphere"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("vSphere", func() {
	var acc *VSphereAcceptance

	BeforeEach(func() {
		acc = NewVSphereAcceptance()

		if !acc.ReadyToTest() {
			Skip("Skipping acceptance tests.")
		}
	})

	Describe("Dry run", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter vsphere.Leftovers
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "khaleesi"
			name := "leftovers-dry-run"
			acc.CreateFolder(filter, name)

			var err error
			deleter, err = vsphere.NewLeftovers(logger, acc.VCenterIP, acc.VCenterUser, acc.VCenterPassword, acc.Datacenter)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("folder: leftovers-dry-run"))
			Expect(stdout.String()).NotTo(ContainSubstring("Are you sure you want to delete"))
			Expect(stdout.String()).NotTo(ContainSubstring("FAILED"))
			Expect(stdout.String()).NotTo(ContainSubstring("SUCCESS deleting leftovers-acceptance!"))
		})
	})

	Describe("Delete", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter vsphere.Leftovers
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "khaleesi"
			name := "leftovers-acceptance"
			acc.CreateFolder(filter, name)

			var err error
			deleter, err = vsphere.NewLeftovers(logger, acc.VCenterIP, acc.VCenterUser, acc.VCenterPassword, acc.Datacenter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).NotTo(ContainSubstring("FAILED"))
			Expect(stdout.String()).To(ContainSubstring("SUCCESS deleting leftovers-acceptance!"))
		})
	})
})
