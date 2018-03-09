package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AWS", func() {
	var acc AWSAcceptance

	BeforeEach(func() {
		iaas := os.Getenv(LEFTOVERS_ACCEPTANCE)
		if strings.ToLower(iaas) != "aws" {
			Skip("Skipping AWS acceptance tests.")
		}

		acc = NewAWSAcceptance()
	})

	Describe("Dry run", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter aws.Leftovers
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "leftovers-dry-run"
			acc.CreateKeyPair(filter)

			var err error
			deleter, err = aws.NewLeftovers(logger, acc.AccessKeyId, acc.SecretAccessKey, acc.Region)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("key pair: leftovers-dry-run"))
			Expect(stdout.String()).NotTo(ContainSubstring("Deleting leftovers-dry-run."))
			Expect(stdout.String()).NotTo(ContainSubstring("SUCCESS deleting leftovers-dry-run!"))
			Expect(stdout.String()).NotTo(ContainSubstring("FAILED deleting key pair leftovers-dry-run"))
		})
	})

	Describe("Delete", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter aws.Leftovers
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "leftovers-acceptance"
			acc.CreateKeyPair(filter)

			var err error
			deleter, err = aws.NewLeftovers(logger, acc.AccessKeyId, acc.SecretAccessKey, acc.Region)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("Deleting leftovers-acceptance."))
			Expect(stdout.String()).To(ContainSubstring("SUCCESS deleting leftovers-acceptance!"))
			Expect(stdout.String()).NotTo(ContainSubstring("FAILED deleting key pair leftovers-acceptance"))
		})
	})
})
