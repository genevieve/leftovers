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
	var (
		acc *VSphereAcceptance

		stdout  *bytes.Buffer
		filter  string
		deleter vsphere.Leftovers
	)

	BeforeEach(func() {
		acc = NewVSphereAcceptance()

		if !acc.ReadyToTest() {
			Skip("Skipping acceptance tests.")
		}

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
			name := "leftovers-dry-run"
			acc.CreateFolder(filter, name)
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
		BeforeEach(func() {
			name := "leftovers-acceptance"
			acc.CreateFolder(filter, name)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).NotTo(ContainSubstring("FAILED"))
			Expect(stdout.String()).To(ContainSubstring("SUCCESS deleting leftovers-acceptance!"))
		})
	})
})
