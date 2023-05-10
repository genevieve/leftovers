package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/gcp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GCP", func() {
	var (
		acc     GCPAcceptance
		stdout  *bytes.Buffer
		filter  string
		deleter gcp.Leftovers
	)

	BeforeEach(func() {
		iaas := os.Getenv(LEFTOVERS_ACCEPTANCE)
		if strings.ToLower(iaas) != "gcp" {
			Skip("Skipping GCP acceptance tests.")
		}

		acc = NewGCPAcceptance()

		noConfirm := true
		debug := true
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm, debug)

		var err error
		deleter, err = gcp.NewLeftovers(logger, acc.KeyPath)
		Expect(err).NotTo(HaveOccurred())

		color.NoColor = true
	})

	Describe("List", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-list-all"
			acc.InsertDisk(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists only the deletable resources that contain the specified filter", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("[Disk: %s]", filter))
			Expect(stdout.String()).To(ContainSubstring("Listing Disks for Zone"))
			Expect(stdout.String()).NotTo(ContainSubstring("[Disk: %s] Deleting...", filter))
		})
	})

	Describe("ListByType", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-list-type"
			acc.InsertDisk(filter)
			acc.InsertCloudRouter(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists only the deletable resources of the specified type", func() {
			deleter.ListByType(filter, "disk")

			Expect(stdout.String()).To(ContainSubstring("[Disk: %s]", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[Disk: %s] Deleting...", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[Router: %s] Deleting...", filter))
		})
	})

	Describe("Types", func() {
		It("lists the resource types that leftovers can delete", func() {
			deleter.Types()

			Expect(stdout.String()).To(ContainSubstring("address"))
			Expect(stdout.String()).To(ContainSubstring("service-account"))
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-delete-all"
			acc.InsertDisk(filter)
			acc.InsertCloudRouter(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[Disk: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[Disk: %s] Deleted!", filter))

			Expect(stdout.String()).To(ContainSubstring("[Router: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[Router: %s] Deleted!", filter))
		})
	})

	Describe("DeleteByType", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-delete-type"
			acc.InsertDisk(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.DeleteByType(filter, "disk")
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[Disk: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[Disk: %s] Deleted!", filter))
		})
	})
})
