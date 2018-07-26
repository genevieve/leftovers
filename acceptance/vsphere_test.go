package acceptance

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/vsphere"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("vSphere", func() {
	var (
		acc VSphereAcceptance

		stdout  *gbytes.Buffer
		filter  string
		deleter vsphere.Leftovers
	)

	BeforeEach(func() {
		iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
		if strings.ToLower(iaas) != "vsphere" {
			Skip("Skipping vSphere acceptance tests.")
		}
		filter = os.Getenv("LEFTOVERS_VSPHERE_FILTER")
		if filter == "" {
			filter = "khaleesi"
		}

		acc = NewVSphereAcceptance()

		noConfirm := true
		stdout = gbytes.NewBuffer()
		logger := app.NewLogger(stdout, os.Stdin, noConfirm)

		var err error
		deleter, err = vsphere.NewLeftovers(logger, acc.VCenterIP, acc.VCenterUser, acc.VCenterPassword, acc.Datacenter)
		Expect(err).NotTo(HaveOccurred())

		color.NoColor = true
	})

	Describe("leftovers", func() {
		BeforeEach(func() {
			rootFolder := acc.FindFolder(filter)
			acc.CreateVM(rootFolder, "leftover-vm")
			nestedFolder := acc.CreateFolder(filter, "leftovers-acceptance")
			acc.CreateVM(nestedFolder, "leftover-nested-vm")
			twiceNestedFolder := acc.CreateFolder(fmt.Sprintf("%s/leftovers-acceptance", filter), "leftovers-nested-acceptance")
			acc.CreateVM(twiceNestedFolder, "leftover-twice-nested-vm")
		})

		It("can list and delete resources with the filter", func() {
			By("listing resources first", func() {
				deleter.List(filter)

				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-vm\]`))
				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-nested-vm\]`))
				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-twice-nested-vm\]`))
				Expect(stdout).NotTo(gbytes.Say(`\[Virtual Machine: leftover-vm\] Deleting\.\.\.`))
				Expect(stdout).NotTo(gbytes.Say(`\[Virtual Machine: leftover-nested-vm\] Deleting\.\.\.`))
				Expect(stdout).NotTo(gbytes.Say(`\[Virtual Machine: leftover-twice-nested-vm\] Deleting\.\.\.`))
			})

			By("successfully deleting VMs", func() {
				err := deleter.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-vm\] Deleting\.\.\.`))
				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-vm\] Deleted!`))
				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-nested-vm\] Deleting\.\.\.`))
				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-nested-vm\] Deleted!`))
				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-twice-nested-vm\] Deleting\.\.\.`))
				Expect(stdout).To(gbytes.Say(`\[Virtual Machine: leftover-twice-nested-vm\] Deleted!`))
				Expect(stdout).To(gbytes.Say(`\[Folder: leftovers-nested-acceptance\] Deleting\.\.\.`))
				Expect(stdout).To(gbytes.Say(`\[Folder: leftovers-nested-acceptance\] Deleted!`))
				Expect(stdout).To(gbytes.Say(`\[Folder: leftovers-acceptance\] Deleting\.\.\.`))
				Expect(stdout).To(gbytes.Say(`\[Folder: leftovers-acceptance\] Deleted!`))
			})
		})
	})
})
