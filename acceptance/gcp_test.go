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
	var acc GCPAcceptance

	BeforeEach(func() {
		iaas := os.Getenv(LEFTOVERS_ACCEPTANCE)
		if strings.ToLower(iaas) != "gcp" {
			Skip("Skipping GCP acceptance tests.")
		}

		acc = NewGCPAcceptance()
	})

	Describe("Dry run", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter gcp.Leftovers
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "leftovers-dry-run"
			acc.InsertDisk(filter)

			var err error
			deleter, err = gcp.NewLeftovers(logger, acc.KeyPath)
			Expect(err).NotTo(HaveOccurred())
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
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter gcp.Leftovers
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "leftovers-acceptance"
			acc.InsertDisk(filter)

			var err error
			deleter, err = gcp.NewLeftovers(logger, acc.KeyPath)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("Deleting leftovers-acceptance."))
			Expect(stdout.String()).To(ContainSubstring("SUCCESS deleting leftovers-acceptance!"))
			Expect(stdout.String()).NotTo(ContainSubstring("ERROR deleting disk"))
		})
	})
})
