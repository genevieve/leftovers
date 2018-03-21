package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/gcp"

	. "github.com/onsi/ginkgo"
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
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm)

		var err error
		deleter, err = gcp.NewLeftovers(logger, acc.KeyPath)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Dry run", func() {
		BeforeEach(func() {
			filter = "leftovers-dry-run"
			acc.InsertDisk(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("disk: leftovers-dry-run"))
			Expect(stdout.String()).NotTo(ContainSubstring("Are you sure you want to delete"))
			Expect(stdout.String()).NotTo(ContainSubstring("Deleting leftovers-dry-run."))
			Expect(stdout.String()).NotTo(ContainSubstring("SUCCESS deleting leftovers-dry-run!"))
			Expect(stdout.String()).NotTo(ContainSubstring("ERROR deleting disk"))
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			filter = "leftovers-acceptance"
			acc.InsertDisk(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("Deleting disk: leftovers-acceptance."))
			Expect(stdout.String()).To(ContainSubstring("SUCCESS deleting disk: leftovers-acceptance!"))
			Expect(stdout.String()).NotTo(ContainSubstring("ERROR deleting disk"))
		})
	})
})
